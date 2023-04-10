package grpcerrors

import (
	"errors"
	"strings"
)

const noRepoOwner = "no repo owner"

func IsNoRepoOwner(err error) bool {
	return err != nil && strings.Contains(err.Error(), noRepoOwner)
}

func NewErrorNoRepoOwner() error {
	return errors.New(noRepoOwner)
}
