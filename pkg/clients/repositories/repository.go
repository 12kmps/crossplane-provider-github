/*
Copyright 2021 The Crossplane Authors.
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

package repositories

import (
	"context"

	"github.com/crossplane-contrib/provider-github/apis/repositories/v1alpha1"
	ghclient "github.com/crossplane-contrib/provider-github/pkg/clients"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/go-github/v33/github"
	"github.com/mitchellh/copystructure"
	"github.com/pkg/errors"
)

const (
	errCheckUpToDate = "unable to determine if external resource is up to date"
)

// Service defines the Repositories operations
type Service interface {
	Create(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error)
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
	Edit(ctx context.Context, owner, repo string, repository *github.Repository) (*github.Repository, *github.Response, error)
	Delete(ctx context.Context, owner, repo string) (*github.Response, error)
}

// NewService creates a new Service based on the *github.Client
// returned by the NewClient SDK method.
func NewService(token string) *Service {
	c := ghclient.NewClient(token)
	r := Service(c.Repositories)
	return &r
}

// IsUpToDate checks whether Repository is configured with given RepositoryParameters.
func IsUpToDate(rp *v1alpha1.RepositoryParameters, observed *github.Repository) (bool, error) {
	generated, err := copystructure.Copy(observed)
	if err != nil {
		return true, errors.Wrap(err, errCheckUpToDate)
	}
	clone, ok := generated.(*github.Repository)
	if !ok {
		return true, errors.New(errCheckUpToDate)
	}

	desired := OverrideParameters(*rp, *clone)

	return cmp.Equal(
		desired,
		*observed,
		cmpopts.IgnoreFields(github.Repository{}, "AutoInit"),
	), nil
}

// OverrideParameters override the parameters in github.Repository
// that are defined in RepositoryParameters
func OverrideParameters(rp v1alpha1.RepositoryParameters, r github.Repository) github.Repository { // nolint:gocyclo
	if len(rp.Name) != 0 {
		r.Name = ghclient.StringPtr(rp.Name)
	}
	if rp.Description != nil {
		r.Description = rp.Description
	}
	if rp.Homepage != nil {
		r.Homepage = rp.Homepage
	}
	if rp.Private != nil {
		r.Private = rp.Private
	}
	if rp.Visibility != nil {
		r.Visibility = rp.Visibility
	}
	if rp.HasIssues != nil {
		r.HasIssues = rp.HasIssues
	}
	if rp.HasProjects != nil {
		r.HasProjects = rp.HasProjects
	}
	if rp.HasWiki != nil {
		r.HasWiki = rp.HasWiki
	}
	if rp.AutoInit != nil {
		r.AutoInit = rp.AutoInit
	}
	if rp.IsTemplate != nil {
		r.IsTemplate = rp.IsTemplate
	}
	if rp.TeamID != nil {
		r.TeamID = rp.TeamID
	}
	if rp.GitignoreTemplate != nil {
		r.GitignoreTemplate = rp.GitignoreTemplate
	}
	if rp.LicenseTemplate != nil {
		r.LicenseTemplate = rp.LicenseTemplate
	}
	if rp.AllowSquashMerge != nil {
		r.AllowSquashMerge = rp.AllowSquashMerge
	}
	if rp.AllowMergeCommit != nil {
		r.AllowMergeCommit = rp.AllowMergeCommit
	}
	if rp.AllowRebaseMerge != nil {
		r.AllowRebaseMerge = rp.AllowRebaseMerge
	}
	if rp.DeleteBranchOnMerge != nil {
		r.DeleteBranchOnMerge = rp.DeleteBranchOnMerge
	}
	if rp.HasPages != nil {
		r.HasPages = rp.HasPages
	}
	if rp.HasDownloads != nil {
		r.HasDownloads = rp.HasDownloads
	}
	if rp.DefaultBranch != nil {
		r.DefaultBranch = rp.DefaultBranch
	}
	if rp.Archived != nil {
		r.Archived = rp.Archived
	}
	return r
}

// GenerateObservation produces RepositoryObservation object from github.Repository object.
func GenerateObservation(r github.Repository) v1alpha1.RepositoryObservation {
	o := v1alpha1.RepositoryObservation{
		ID:               ghclient.Int64Value(r.ID),
		NodeID:           ghclient.StringValue(r.NodeID),
		FullName:         ghclient.StringValue(r.FullName),
		Name:             ghclient.StringValue(r.Name),
		URL:              ghclient.StringValue(r.URL),
		ArchiveURL:       ghclient.StringValue(r.ArchiveURL),
		AssigneesURL:     ghclient.StringValue(r.AssigneesURL),
		BlobsURL:         ghclient.StringValue(r.BlobsURL),
		CollaboratorsURL: ghclient.StringValue(r.CollaboratorsURL),
		CommentsURL:      ghclient.StringValue(r.CommentsURL),
		CommitsURL:       ghclient.StringValue(r.CommitsURL),
		CompareURL:       ghclient.StringValue(r.CompareURL),
		ContentsURL:      ghclient.StringValue(r.ContentsURL),
		ContributorsURL:  ghclient.StringValue(r.ContributorsURL),
		DeploymentsURL:   ghclient.StringValue(r.DeploymentsURL),
		DownloadsURL:     ghclient.StringValue(r.DownloadsURL),
		EventsURL:        ghclient.StringValue(r.EventsURL),
		ForksURL:         ghclient.StringValue(r.ForksURL),
		GitCommitsURL:    ghclient.StringValue(r.GitCommitsURL),
		GitRefsURL:       ghclient.StringValue(r.GitRefsURL),
		GitTagsURL:       ghclient.StringValue(r.GitTagsURL),
		HooksURL:         ghclient.StringValue(r.HooksURL),
		IssueCommentURL:  ghclient.StringValue(r.IssueCommentURL),
		IssueEventsURL:   ghclient.StringValue(r.IssueEventsURL),
		IssuesURL:        ghclient.StringValue(r.IssuesURL),
		KeysURL:          ghclient.StringValue(r.KeysURL),
		LabelsURL:        ghclient.StringValue(r.LabelsURL),
		LanguagesURL:     ghclient.StringValue(r.LanguagesURL),
		MergesURL:        ghclient.StringValue(r.MergesURL),
		MilestonesURL:    ghclient.StringValue(r.MilestonesURL),
		NotificationsURL: ghclient.StringValue(r.NotificationsURL),
		PullsURL:         ghclient.StringValue(r.PullsURL),
		ReleasesURL:      ghclient.StringValue(r.ReleasesURL),
		StargazersURL:    ghclient.StringValue(r.StargazersURL),
		StatusesURL:      ghclient.StringValue(r.StatusesURL),
		SubscribersURL:   ghclient.StringValue(r.SubscribersURL),
		SubscriptionURL:  ghclient.StringValue(r.SubscriptionURL),
		TagsURL:          ghclient.StringValue(r.TagsURL),
		TreesURL:         ghclient.StringValue(r.TreesURL),
		TeamsURL:         ghclient.StringValue(r.TeamsURL),
		HTMLURL:          ghclient.StringValue(r.HTMLURL),
		CloneURL:         ghclient.StringValue(r.CloneURL),
		GitURL:           ghclient.StringValue(r.GitURL),
		MirrorURL:        ghclient.StringValue(r.MirrorURL),
		SSHURL:           ghclient.StringValue(r.SSHURL),
		SVNURL:           ghclient.StringValue(r.SVNURL),
		ForksCount:       ghclient.IntValue(r.ForksCount),
		NetworkCount:     ghclient.IntValue(r.NetworkCount),
		OpenIssuesCount:  ghclient.IntValue(r.OpenIssuesCount),
		StargazersCount:  ghclient.IntValue(r.StargazersCount),
		SubscribersCount: ghclient.IntValue(r.SubscribersCount),
		WatchersCount:    ghclient.IntValue(r.WatchersCount),
		CreatedAt:        ghclient.ConvertTimestamp(r.CreatedAt),
		PushedAt:         ghclient.ConvertTimestamp(r.PushedAt),
		UpdatedAt:        ghclient.ConvertTimestamp(r.UpdatedAt),
		Language:         ghclient.StringValue(r.Language),
		Fork:             ghclient.BoolValue(r.Fork),
		Size:             ghclient.IntValue(r.Size),
		Disabled:         ghclient.BoolValue(r.Disabled),
		Topics:           r.Topics,
	}
	if r.Permissions != nil {
		o.Permissions = map[string]bool{}
		for k, v := range *r.Permissions {
			o.Permissions[k] = v
		}
	}

	return o
}

// LateInitialize fills the empty fields of RepositoryParameters if the corresponding
// fields are given in Repository.
func LateInitialize(rp *v1alpha1.RepositoryParameters, r github.Repository) { // nolint:gocyclo
	if rp.Organization == nil && ghclient.StringValue(r.Owner.Type) == "Organization" {
		if r.Organization.Login != nil {
			rp.Organization = r.Organization.Login
		}
	}
	if rp.Description == nil && r.Description != nil {
		rp.Description = r.Description
	}
	if rp.Homepage == nil && r.Homepage != nil {
		rp.Homepage = r.Homepage
	}
	if rp.Private == nil && r.Private != nil {
		rp.Private = r.Private
	}
	if rp.Visibility == nil && r.Visibility != nil {
		rp.Visibility = r.Visibility
	}
	if rp.HasIssues == nil && r.HasIssues != nil {
		rp.HasIssues = r.HasIssues
	}
	if rp.HasProjects == nil && r.HasProjects != nil {
		rp.HasProjects = r.HasProjects
	}
	if rp.HasWiki == nil && r.HasWiki != nil {
		rp.HasWiki = r.HasWiki
	}
	if rp.IsTemplate == nil && r.IsTemplate != nil {
		rp.IsTemplate = r.IsTemplate
	}
	if rp.TeamID == nil && r.TeamID != nil {
		rp.TeamID = r.TeamID
	}
	if rp.AutoInit == nil && r.AutoInit != nil {
		rp.AutoInit = r.AutoInit
	}
	if rp.GitignoreTemplate == nil && r.GitignoreTemplate != nil {
		rp.GitignoreTemplate = r.GitignoreTemplate
	}
	if rp.LicenseTemplate == nil && r.LicenseTemplate != nil {
		rp.LicenseTemplate = r.LicenseTemplate
	}
	if rp.AllowSquashMerge == nil && r.AllowSquashMerge != nil {
		rp.AllowSquashMerge = r.AllowSquashMerge
	}
	if rp.AllowMergeCommit == nil && r.AllowMergeCommit != nil {
		rp.AllowMergeCommit = r.AllowMergeCommit
	}
	if rp.AllowRebaseMerge == nil && r.AllowRebaseMerge != nil {
		rp.AllowRebaseMerge = r.AllowRebaseMerge
	}
	if rp.DeleteBranchOnMerge == nil && r.DeleteBranchOnMerge != nil {
		rp.DeleteBranchOnMerge = r.DeleteBranchOnMerge
	}
	if rp.HasPages == nil && r.HasPages != nil {
		rp.HasPages = r.HasPages
	}
	if rp.HasDownloads == nil && r.HasDownloads != nil {
		rp.HasDownloads = r.HasDownloads
	}
	if rp.DefaultBranch == nil && r.DefaultBranch != nil {
		rp.DefaultBranch = r.DefaultBranch
	}
	if rp.Archived == nil && r.Archived != nil {
		rp.Archived = r.Archived
	}
}
