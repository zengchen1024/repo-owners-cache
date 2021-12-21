package cache

import (
	libpath "path"
	"regexp"

	"k8s.io/apimachinery/pkg/util/sets"
)

const baseDirConvention = "."

// RepoOwner is an interface to work with repoowners
type RepoOwner interface {
	FindApproverOwnersForFile(path string) string
	FindReviewersOwnersForFile(path string) string
	LeafApprovers(path string) sets.String
	Approvers(path string) sets.String
	LeafReviewers(path string) sets.String
	Reviewers(path string) sets.String
	AllReviewers() sets.String
	TopLevelApprovers() sets.String
}

type getConfigItem func(*Config) []string

type fileOwnerInfo map[*regexp.Regexp]Config

func (fo fileOwnerInfo) getConfig(path string, getValue getConfigItem) (Config, bool) {
	for re, s := range fo {
		if len(getValue(&s)) > 0 && re != nil && re.MatchString(path) {
			return s, true
		}
	}
	return Config{}, false
}

func (fo fileOwnerInfo) add(re *regexp.Regexp, config *Config) {
	fo[re] = *config
}

type RepoOwnerInfo struct {
	dirOwners  map[string]SimpleConfig
	fileOwners map[string]fileOwnerInfo
}

func newRepoOwnerInfo() *RepoOwnerInfo {
	return &RepoOwnerInfo{
		dirOwners:  make(map[string]SimpleConfig),
		fileOwners: make(map[string]fileOwnerInfo),
	}
}

func (o *RepoOwnerInfo) IsEmpty() bool {
	return o == nil || (len(o.dirOwners) == 0 && len(o.fileOwners) == 0)
}

// findOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// The path variable should be a full path to a filename
func (o *RepoOwnerInfo) findOwnersForFile(path string, getValue getConfigItem) string {
	d := libpath.Dir(path)

	if fo, ok := o.fileOwners[d]; ok {
		if _, b := fo.getConfig(path, getValue); b {
			return d
		}
	}

	for ; d != baseDirConvention; d = libpath.Dir(d) {
		if s, ok := o.dirOwners[d]; ok {
			// if the approver or reviewer is not set at this dir,
			// lookup until find it even if the no_parent_owners is set.
			if len(getValue(&s.Config)) > 0 {
				return d
			}
		}
	}

	return baseDirConvention
}

// FindApproverOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// that contains an approvers section
func (o *RepoOwnerInfo) FindApproverOwnersForFile(path string) string {
	return o.findOwnersForFile(path, func(c *Config) []string {
		return c.Approvers
	})
}

// FindReviewersOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// that contains a reviewers section
func (o *RepoOwnerInfo) FindReviewersOwnersForFile(path string) string {
	return o.findOwnersForFile(path, func(c *Config) []string {
		return c.Reviewers
	})
}

// entriesForFile returns a set of users who are assignees to the requested file.
// The path variable should be a full path to a filename.
// leafOnly indicates whether only the OWNERS deepest in the tree (closest to the file)
// should be returned or if all OWNERS in filepath should be returned.
func (o *RepoOwnerInfo) entriesForFile(path string, leafOnly bool, getValue getConfigItem) sets.String {
	d := libpath.Dir(path)

	if fo, ok := o.fileOwners[d]; ok {
		if c, b := fo.getConfig(path, getValue); b {
			return sets.NewString(getValue(&c)...)
		}
	}

	out := sets.NewString()

	for {
		if s, ok := o.dirOwners[d]; ok {
			out.Insert(getValue(&s.Config)...)

			if out.Len() > 0 && (s.Options.NoParentOwners || leafOnly) {
				break
			}
		}

		if d == baseDirConvention {
			break
		}

		d = libpath.Dir(d)
	}

	return out
}

// LeafApprovers returns a set of users who are the closest approvers to the
// requested file. If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will only return user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) LeafApprovers(path string) sets.String {
	return o.entriesForFile(path, true, func(c *Config) []string {
		return c.Approvers
	})
}

// Approvers returns ALL of the users who are approvers for the
// requested file (including approvers in parent dirs' OWNERS).
// If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will return both user1 and user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) Approvers(path string) sets.String {
	return o.entriesForFile(path, false, func(c *Config) []string {
		return c.Approvers
	})
}

// LeafReviewers returns a set of users who are the closest reviewers to the
// requested file. If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will only return user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) LeafReviewers(path string) sets.String {
	return o.entriesForFile(path, true, func(c *Config) []string {
		return c.Reviewers
	})
}

// Reviewers returns ALL of the users who are reviewers for the
// requested file (including reviewers in parent dirs' OWNERS).
// If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will return both user1 and user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) Reviewers(path string) sets.String {
	return o.entriesForFile(path, false, func(c *Config) []string {
		return c.Reviewers
	})
}

func (o *RepoOwnerInfo) TopLevelApprovers() sets.String {
	return o.entriesForFile(".", false, func(c *Config) []string {
		return c.Approvers
	})
}

func (o *RepoOwnerInfo) AllReviewers() sets.String {
	r := sets.NewString()

	for _, s := range o.dirOwners {
		r.Insert(s.Approvers...)
		r.Insert(s.Reviewers...)
	}

	for _, v := range o.fileOwners {
		for _, s := range v {
			r.Insert(s.Approvers...)
			r.Insert(s.Reviewers...)
		}
	}

	return r
}
