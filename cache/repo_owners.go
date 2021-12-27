package cache

import (
	libpath "path"

	"k8s.io/apimachinery/pkg/util/sets"
)

const rootPath = "."

func getPath(p string) string {
	return libpath.Dir(p)
}

// RepoOwner is an interface to work with repoowners
type RepoOwner interface {
	FindApproverOwnersForFile(path string) string
	FindReviewersOwnersForFile(path string) string
	LeafApprovers(path string) sets.String
	Approvers(path string) sets.String
	LeafReviewers(path string) sets.String
	Reviewers(path string) sets.String
	IsNoParentOwners(path string) bool
	AllReviewers() sets.String
	TopLevelApprovers() sets.String
}

// findOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// The path variable should be a full path to a filename
func (o *RepoOwnerInfo) findOwnersForFile(path string, getValue getConfigItem) string {
	d := getPath(path)

	if fo, ok := o.fileOwners[d]; ok {
		if _, b := fo.getConfig(path, getValue); b {
			return d
		}
	}

	for ; d != rootPath; d = getPath(d) {
		if s, ok := o.dirOwners[d]; ok {
			// if the approver or reviewer is not set at this dir,
			// lookup until find it even if the no_parent_owners is set.
			if len(getValue(&s.ownersConfig)) > 0 {
				return d
			}
		}
	}

	return rootPath
}

// FindApproverOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// that contains an approvers section
func (o *RepoOwnerInfo) FindApproverOwnersForFile(path string) string {
	return o.findOwnersForFile(path, func(c *ownersConfig) []string {
		return c.Approvers
	})
}

// FindReviewersOwnersForFile returns the OWNERS file path furthest down the tree for a specified file
// that contains a reviewers section
func (o *RepoOwnerInfo) FindReviewersOwnersForFile(path string) string {
	return o.findOwnersForFile(path, func(c *ownersConfig) []string {
		return c.Reviewers
	})
}

// IsNoParentOwners checks if an OWNERS file path refers to an OWNERS file with NoParentOwners enabled.
func (o *RepoOwnerInfo) IsNoParentOwners(path string) bool {
	return o.dirOwners[path].Options.NoParentOwners
}

// entriesForFile returns a set of users who are assignees to the requested file.
// The path variable should be a full path to a filename.
// leafOnly indicates whether only the OWNERS deepest in the tree (closest to the file)
// should be returned or if all OWNERS in filepath should be returned.
func (o *RepoOwnerInfo) entriesForFile(path string, leafOnly bool, getValue getConfigItem) sets.String {
	d := getPath(path)

	if fo, ok := o.fileOwners[d]; ok {
		if c, b := fo.getConfig(path, getValue); b {
			return sets.NewString(getValue(&c)...)
		}
	}

	out := sets.NewString()

	for {
		if s, ok := o.dirOwners[d]; ok {
			out.Insert(getValue(&s.ownersConfig)...)

			if out.Len() > 0 && (s.Options.NoParentOwners || leafOnly) {
				break
			}
		}

		if d == rootPath {
			break
		}

		d = getPath(d)
	}

	return out
}

// LeafApprovers returns a set of users who are the closest approvers to the
// requested file. If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will only return user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) LeafApprovers(path string) sets.String {
	return o.entriesForFile(path, true, func(c *ownersConfig) []string {
		return c.Approvers
	})
}

// Approvers returns ALL of the users who are approvers for the
// requested file (including approvers in parent dirs' OWNERS).
// If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will return both user1 and user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) Approvers(path string) sets.String {
	return o.entriesForFile(path, false, func(c *ownersConfig) []string {
		return c.Approvers
	})
}

// LeafReviewers returns a set of users who are the closest reviewers to the
// requested file. If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will only return user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) LeafReviewers(path string) sets.String {
	return o.entriesForFile(path, true, func(c *ownersConfig) []string {
		return c.Reviewers
	})
}

// Reviewers returns ALL of the users who are reviewers for the
// requested file (including reviewers in parent dirs' OWNERS).
// If pkg/OWNERS has user1 and pkg/util/OWNERS has user2 this
// will return both user1 and user2 for the path pkg/util/sets/file.go
func (o *RepoOwnerInfo) Reviewers(path string) sets.String {
	return o.entriesForFile(path, false, func(c *ownersConfig) []string {
		return c.Reviewers
	})
}

// TopLevelApprovers gets the approvers at directory of '.'.
func (o *RepoOwnerInfo) TopLevelApprovers() sets.String {
	return o.entriesForFile(rootPath, false, func(c *ownersConfig) []string {
		return c.Approvers
	})
}

// AllReviewers gets all reviewers including approvers.
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
