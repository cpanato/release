/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/release-sdk/git"

	"k8s.io/release/pkg/fastforward"
)

func (s *sut) getFfOptions() *fastforward.Options {
	return &fastforward.Options{
		MainRef:        git.Remotify(git.DefaultBranch),
		NonInteractive: true,
		RepoPath:       s.repo.Dir(),
	}
}

func TestFfFailedWithoutReleaseBranch(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	// When
	err := fastforward.New(&fastforward.Options{}).Run()
	// Then
	require.Error(t, err)
}

func TestFfFailedNoReleaseBranch(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	ffo := s.getFfOptions()
	ffo.Branch = "not-a-release-branch"

	// When
	err := fastforward.New(ffo).Run()

	// Then
	require.Error(t, err)
}

func TestFfFailedReleaseBranchDoesNotExist(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	ffo := s.getFfOptions()
	ffo.Branch = "release-1.999"

	// When
	err := fastforward.New(ffo).Run()

	// Then
	require.Error(t, err)
}

func TestFfFailedOldReleaseBranch(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	ffo := s.getFfOptions()
	ffo.Branch = "release-1.17"

	// When
	err := fastforward.New(ffo).Run()

	// Then
	require.Error(t, err)
}

func TestFfSuccessDryRun(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	ffo := s.getFfOptions()
	ffo.Branch = pseudoReleaseBranch

	// When
	err := fastforward.New(ffo).Run()

	// Then
	require.NoError(t, err)

	// Local should contain the commit
	lastLocalCommit := s.lastCommit(t, pseudoReleaseBranch)
	require.Contains(t, lastLocalCommit, testCommitMessage)

	// Remote should not be modified
	lastRemoteCommit := s.lastCommit(t, git.Remotify(pseudoReleaseBranch))
	require.NotContains(t, lastRemoteCommit, testCommitMessage)
}

func TestFfSuccess(t *testing.T) {
	// Given
	s := newSUT(t)
	defer s.cleanup(t)

	ffo := s.getFfOptions()
	ffo.Branch = pseudoReleaseBranch

	ffo.NoMock = true

	// When
	err := fastforward.New(ffo).Run()

	// Then
	require.NoError(t, err)

	// Local should contain the commit
	lastLocalCommit := s.lastCommit(t, pseudoReleaseBranch)
	require.Contains(t, lastLocalCommit, testCommitMessage)

	// Remote should be modified
	lastRemoteCommit := s.lastCommit(t, git.Remotify(pseudoReleaseBranch))
	require.Contains(t, lastRemoteCommit, testCommitMessage)
}
