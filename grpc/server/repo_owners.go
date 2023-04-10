package server

import (
	"context"

	"github.com/opensourceways/repo-owners-cache/cache"
	"github.com/opensourceways/repo-owners-cache/grpc/grpcerrors"
	"github.com/opensourceways/repo-owners-cache/protocol"
)

type repoOwnersServer struct {
	c *cache.Cache
	protocol.UnimplementedRepoOwnersServer
}

func (r *repoOwnersServer) loadRepoOwners(b *protocol.Branch) (cache.RepoOwner, error) {
	o, err := r.c.LoadRepoOwners(cache.RepoBranch{
		Platform: b.GetPlatform(),
		Org:      b.GetOrg(),
		Repo:     b.GetRepo(),
		Branch:   b.GetBranch(),
	})
	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, grpcerrors.NewErrorNoRepoOwner()
	}

	return o, nil
}

func (r *repoOwnersServer) FindApproverOwnersForFile(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Path, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Path{
		Value: o.FindApproverOwnersForFile(fp.GetFile()),
	}, nil
}

func (r *repoOwnersServer) FindReviewersOwnersForFile(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Path, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Path{
		Value: o.FindReviewersOwnersForFile(fp.GetFile()),
	}, nil
}

func (r *repoOwnersServer) LeafApprovers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.LeafApprovers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) Approvers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.Approvers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) LeafReviewers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.LeafReviewers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) Reviewers(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.Reviewers(fp.GetFile()).UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) IsNoParentOwners(ctx context.Context, fp *protocol.RepoFilePath) (*protocol.NoParentOwners, error) {
	o, err := r.loadRepoOwners(fp.GetBranch())
	if err != nil {
		return nil, err
	}

	return &protocol.NoParentOwners{
		Value: o.IsNoParentOwners(fp.GetFile()),
	}, nil
}

func (r *repoOwnersServer) AllReviewers(ctx context.Context, b *protocol.Branch) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(b)
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.AllReviewers().UnsortedList(),
	}, nil
}

func (r *repoOwnersServer) TopLevelApprovers(ctx context.Context, b *protocol.Branch) (*protocol.Owners, error) {
	o, err := r.loadRepoOwners(b)
	if err != nil {
		return nil, err
	}

	return &protocol.Owners{
		Value: o.TopLevelApprovers().UnsortedList(),
	}, nil
}
