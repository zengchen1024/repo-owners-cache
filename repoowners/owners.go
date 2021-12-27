package repoowners

import (
	"context"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/grpc/client"
	"github.com/opensourceways/repo-owners-cache/grpc/server"
	"github.com/opensourceways/repo-owners-cache/protocol"
)

type RepoOwner = cache.RepoOwner

type RepoBranch = cache.RepoBranch

type owners struct {
	c *client.Client
	b *protocol.Branch
}

func NewRepoOwners(branch RepoBranch, c *client.Client) (RepoOwner, error) {
	b := protocol.Branch{
		Platform: branch.Platform,
		Org:      branch.Org,
		Repo:     branch.Repo,
		Branch:   branch.Branch,
	}

	_, err := c.TopLevelApprovers(context.Background(), &b)
	if err != nil {
		if server.IsNoRepoOwner(err) {
			return nil, nil
		}

		return nil, err
	}

	return &owners{
		b: &b,
		c: c,
	}, nil
}

func (o *owners) repoFilePath(path string) protocol.RepoFilePath {
	return protocol.RepoFilePath{
		Branch: o.b,
		File:   path,
	}
}

func (o *owners) FindApproverOwnersForFile(path string) string {
	p := o.repoFilePath(path)
	v, _ := o.c.FindApproverOwnersForFile(context.Background(), &p)

	return v.GetValue()
}

func (o *owners) FindReviewersOwnersForFile(path string) string {
	p := o.repoFilePath(path)
	v, _ := o.c.FindReviewersOwnersForFile(context.Background(), &p)

	return v.GetValue()
}

func (o *owners) LeafApprovers(path string) sets.String {
	p := o.repoFilePath(path)
	v, _ := o.c.LeafApprovers(context.Background(), &p)

	return sets.NewString(v.GetValue()...)
}

func (o *owners) LeafReviewers(path string) sets.String {
	p := o.repoFilePath(path)
	v, _ := o.c.LeafReviewers(context.Background(), &p)

	return sets.NewString(v.GetValue()...)
}

func (o *owners) Approvers(path string) sets.String {
	p := o.repoFilePath(path)
	v, _ := o.c.Approvers(context.Background(), &p)

	return sets.NewString(v.GetValue()...)
}

func (o *owners) Reviewers(path string) sets.String {
	p := o.repoFilePath(path)
	v, _ := o.c.Reviewers(context.Background(), &p)

	return sets.NewString(v.GetValue()...)
}

func (o *owners) IsNoParentOwners(path string) bool {
	p := o.repoFilePath(path)
	v, _ := o.c.IsNoParentOwners(context.Background(), &p)
	return v.GetValue()
}

func (o *owners) AllReviewers() sets.String {
	v, _ := o.c.AllReviewers(context.Background(), o.b)
	return sets.NewString(v.GetValue()...)
}

func (o *owners) TopLevelApprovers() sets.String {
	v, _ := o.c.TopLevelApprovers(context.Background(), o.b)
	return sets.NewString(v.GetValue()...)
}

type repoMembers struct {
	members sets.String
}

func RepoMemberAsOwners(d []string) RepoOwner {
	return &repoMembers{
		members: sets.NewString(d...),
	}
}

func (o *repoMembers) FindApproverOwnersForFile(path string) string {
	return "."
}

func (o *repoMembers) FindReviewersOwnersForFile(path string) string {
	return "."
}

func (o *repoMembers) LeafApprovers(path string) sets.String {
	return o.members
}

func (o *repoMembers) LeafReviewers(path string) sets.String {
	return o.members
}

func (o *repoMembers) Approvers(path string) sets.String {
	return o.members
}

func (o *repoMembers) Reviewers(path string) sets.String {
	return o.members
}

func (o *repoMembers) IsNoParentOwners(path string) bool {
	if path == "." {
		return true
	}

	return false
}

func (o *repoMembers) AllReviewers() sets.String {
	return o.members
}

func (o *repoMembers) TopLevelApprovers() sets.String {
	return o.members
}
