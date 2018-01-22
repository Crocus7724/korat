package api

import (
	"testing"
	"github.com/shurcooL/githubql"
)

func TestGetViewerPullRequests(t *testing.T) {
	p, err := GetViewerPullRequests("SengokuConquest", []githubql.PullRequestState{
		githubql.PullRequestStateOpen,
		githubql.PullRequestStateClosed,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(p)
}
