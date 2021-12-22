package server

import (
	"context"
	"fmt"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/protocol"
)

var noRepoOwner = fmt.Errorf("no repo owner")

type repoOwnersServer struct {
	c *cache.Cache
	protocol.UnimplementedRepoOwnersServer
}

func (r *repoOwnersServer) FindApproverOwnersForFile(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Path, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Path{
		Path: o.FindApproverOwnersForFile(fp.GetFile()),
	}, nil
}

func (r *repoOwnersServer) FindReviewersOwnersForFile(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Path, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Path{
		Path: o.FindReviewersOwnersForFile(fp.GetFile()),
	}, nil
}

func (r *repoOwnersServer) LeafApprovers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, nil
	}

	return &protocol.Owners{
		Owners: o.LeafApprovers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) Approvers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Owners{
		Owners: o.Approvers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) LeafReviewers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Owners{
		Owners: o.LeafReviewers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) Reviewers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o := r.loadRepoOwners(fp.GetBranch())
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Owners{
		Owners: o.Reviewers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) AllReviewers(ctx context.Context, b *protocol.Branch) (*protocol.Owners, error) {
	o := r.loadRepoOwners(b)
	if o == nil {
		return nil, fmt.Errorf("no data")
	}

	return &protocol.Owners{
		Owners: o.AllReviewers().UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) TopLevelApprovers(ctx context.Context, b *protocol.Branch) (*protocol.Owners, error) {
	o := r.loadRepoOwners(b)
	if o == nil {
		return nil, noRepoOwner
	}

	return &protocol.Owners{
		Owners: o.TopLevelApprovers().UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) loadRepoOwners(b *protocol.Branch) cache.RepoOwner {
	return r.c.LoadRepoOwners(cache.RepoBranch{
		Platform: b.GetPlatform(),
		Org:      b.GetOrg(),
		Repo:     b.GetRepo(),
		Branch:   b.GetBranch(),
	})
}
