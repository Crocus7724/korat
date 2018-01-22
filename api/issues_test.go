package api

import (
	"testing"
	"github.com/shurcooL/githubql"
)

func TestGetViewerIssues(t *testing.T) {
	i, err := GetViewerIssues("ToDo", []githubql.IssueState{
		githubql.IssueStateClosed,
		githubql.IssueStateOpen,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(i)
}
