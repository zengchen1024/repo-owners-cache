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

type Config struct {
	Approvers         []string `json:"approvers,omitempty"`
	Reviewers         []string `json:"reviewers,omitempty"`
	RequiredReviewers []string `json:"required_reviewers,omitempty"`
	Labels            []string `json:"labels,omitempty"`
}

func (c Config) empty() bool {
	return len(c.Approvers) == 0 && len(c.Reviewers) == 0 && len(c.RequiredReviewers) == 0 && len(c.Labels) == 0
}

type SimpleConfig struct {
	Options dirOptions `json:"options,omitempty"`
	Config  `json:",inline"`
}

// Empty checks if a SimpleConfig could be considered empty
func (s *SimpleConfig) Empty() bool {
	return s.Config.empty()
}

type ownersConfig struct {
	SimpleConfig

	Files map[string]Config `json:"files,omitempty"`
}

func normalConfig(c *Config) *Config {
	// TODO
	return c
}

func (o *RepoOwnerInfo) applyDirConfigToPath(path string, config *SimpleConfig) {
	o.dirOwners[path] = SimpleConfig{
		Config:  *normalConfig(&config.Config),
		Options: config.Options,
	}
}

func (o *RepoOwnerInfo) applyFileConfigToPath(path string, re *regexp.Regexp, config *Config) {
	if _, ok := o.fileOwners[path]; !ok {
		o.fileOwners[path] = make(fileOwnerInfo)
	}

	o.fileOwners[path].add(re, normalConfig(config))
}

func (o *RepoOwnerInfo) parseOwnerConfig(dir, content string, log *logrus.Entry) error {
	c := new(ownersConfig)
	if err := parseYaml(content, c); err != nil {
		return err
	}

	for pattern, config := range c.Files {
		if pattern == "" || pattern == ".*" || config.empty() {
			continue
		}

		if re, err := regexp.Compile(pattern); err != nil {
			log.WithError(err).Errorf("Invalid regexp %q.", pattern)
		} else {
			o.applyFileConfigToPath(dir, re, &config)
		}
	}

	if !c.SimpleConfig.Empty() {
		o.applyDirConfigToPath(dir, &c.SimpleConfig)
	}
	return nil
}

func parseYaml(content string, r *ownersConfig) error {
	b, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, r)
}
