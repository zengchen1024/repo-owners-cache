package cache

import (
	"encoding/base64"
	"regexp"

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

type RepoOwnerInfo struct {
	dirOwners  map[string]dirOwnerInfo
	fileOwners map[string]fileOwnerInfo
}

func newRepoOwnerInfo() *RepoOwnerInfo {
	return &RepoOwnerInfo{
		dirOwners:  make(map[string]dirOwnerInfo),
		fileOwners: make(map[string]fileOwnerInfo),
	}
}

func (o *RepoOwnerInfo) isEmpty() bool {
	return o == nil || (len(o.dirOwners) == 0 && len(o.fileOwners) == 0)
}

func (o *RepoOwnerInfo) setDirOwners(path string, config *dirOwnerInfo) {
	o.dirOwners[path] = dirOwnerInfo{
		ownersConfig: *normalConfig(&config.ownersConfig),
		Options:      config.Options,
	}
}

func (o *RepoOwnerInfo) setFileOwners(path string, re *regexp.Regexp, config *ownersConfig) {
	if _, ok := o.fileOwners[path]; !ok {
		o.fileOwners[path] = make(fileOwnerInfo)
	}

	o.fileOwners[path].add(re, normalConfig(config))
}

func (o *RepoOwnerInfo) parseOwnerConfig(dir, content string, log *logrus.Entry) error {
	c := new(ownersFile)
	if err := parseYaml(content, c); err != nil {
		return err
	}

	for pattern, config := range c.Files {
		if pattern == "" || pattern == ".*" || config.empty() {
			continue
		}

		if re, err := regexp.Compile(pattern); err != nil {
			log.Errorf("Invalid regexp %q, err:%s", pattern, err.Error())
		} else {
			o.setFileOwners(dir, re, &config)
		}
	}

	if !c.dirOwnerInfo.isEmpty() {
		o.setDirOwners(dir, &c.dirOwnerInfo)
	}
	return nil
}

func parseYaml(content string, r *ownersFile) error {
	b, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, r)
}

func normalConfig(c *ownersConfig) *ownersConfig {
	// TODO: change to lowercase?
	return c
}
