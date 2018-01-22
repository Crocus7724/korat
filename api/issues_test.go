package api

import "testing"

func TestGetViewerIssues(t *testing.T) {
	i, err := GetViewerIssues("ToDo")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(i)
}
