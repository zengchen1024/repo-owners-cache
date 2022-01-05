package cache

import (
	"encoding/base64"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

type dirOptions struct {
	NoParentOwners bool `json:"no_parent_owners,omitempty"`
}

type ownersConfig struct {
	Approvers []string `json:"approvers,omitempty"`
	Reviewers []string `json:"reviewers,omitempty"`
}

func (c ownersConfig) empty() bool {
	return len(c.Approvers) == 0 && len(c.Reviewers) == 0
}

type dirOwnerInfo struct {
	ownersConfig `json:",inline"`

	Options dirOptions `json:"options,omitempty"`
}

func (s *dirOwnerInfo) isEmpty() bool {
	return s.ownersConfig.empty()
}

type ownersFile struct {
	dirOwnerInfo `json:",inline"`

	Files map[string]ownersConfig `json:"files,omitempty"`
}

type getConfigItem func(*ownersConfig) []string

type fileOwnerInfo map[*regexp.Regexp]ownersConfig

func (fo fileOwnerInfo) getConfig(path string, getValue getConfigItem) (ownersConfig, bool) {
	for re, s := range fo {
		if len(getValue(&s)) > 0 && re != nil && re.MatchString(path) {
			return s, true
		}
	}
	return ownersConfig{}, false
}

func (fo fileOwnerInfo) add(re *regexp.Regexp, config *ownersConfig) {
	fo[re] = *config
}

func (fo fileOwnerInfo) isEmpty() bool {
	return len(fo) == 0
}

type RepoOwnerInfo struct {
	dirOwners  map[string]dirOwnerInfo
	fileOwners map[string]fileOwnerInfo
	files      map[string]string
}

func newRepoOwnerInfo() *RepoOwnerInfo {
	return &RepoOwnerInfo{
		dirOwners:  make(map[string]dirOwnerInfo),
		fileOwners: make(map[string]fileOwnerInfo),
		files:      make(map[string]string),
	}
}

func (o *RepoOwnerInfo) isEmpty() bool {
	return o == nil || (len(o.dirOwners) == 0 && len(o.fileOwners) == 0)
}

func (o *RepoOwnerInfo) getFileSHA(dir string) string {
	return o.files[dir]
}

func (o *RepoOwnerInfo) copyOwnerFiles(no *RepoOwnerInfo) {
	for dir, sha := range o.files {
		if _, ok := no.files[dir]; !ok {
			no.files[dir] = sha
			no.dirOwners[dir] = o.dirOwners[dir]
			no.fileOwners[dir] = o.fileOwners[dir]
		}
	}
}

func (o *RepoOwnerInfo) update(path, sha string, dirOwner *dirOwnerInfo, fileOwner fileOwnerInfo) {
	done := false
	setSHA := func() {
		if !done {
			done = true
			o.files[path] = sha
		}
	}

	if !dirOwner.isEmpty() {
		o.dirOwners[path] = *dirOwner

		setSHA()
	}

	if !fileOwner.isEmpty() {
		o.fileOwners[path] = fileOwner

		setSHA()
	}
}

func (o *RepoOwnerInfo) parseOwnerConfig(dir, content, sha string, log *logrus.Entry) error {
	c := new(ownersFile)
	if err := parseYaml(content, c); err != nil {
		log.Errorf("parse file:%s/%s, err:%s", dir, sha, err.Error())
		return err
	}

	fileOwner := make(fileOwnerInfo)

	for pattern, config := range c.Files {
		if pattern == "" || pattern == ".*" || config.empty() {
			continue
		}

		if re, err := regexp.Compile(pattern); err != nil {
			log.Errorf("Invalid regexp %q, err:%s", pattern, err.Error())
		} else {
			v := normalConfig(&config)
			fileOwner.add(re, &v)
		}
	}

	dirOwner := &c.dirOwnerInfo

	if !dirOwner.isEmpty() {
		dirOwner.ownersConfig = normalConfig(&dirOwner.ownersConfig)
	}

	o.update(dir, sha, dirOwner, fileOwner)

	return nil
}

func parseYaml(content string, r *ownersFile) error {
	b, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, r)
}

func normalConfig(c *ownersConfig) ownersConfig {
	f := func(v []string) []string {
		n := len(v)
		if n == 0 {
			return nil
		}

		r := make([]string, n)
		for i := range v {
			r[i] = strings.ToLower(v[i])
		}

		return r
	}

	return ownersConfig{
		Approvers: f(c.Approvers),
		Reviewers: f(c.Reviewers),
	}
}
