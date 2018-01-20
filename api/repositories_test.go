package api

import "testing"

const token = "805417717f94ea666ca00fdf31687e6423db5113"

func TestMain(m *testing.M) {
	Init("https://api.github.com/graphql", token)
	m.Run()
}

func TestGetViewerRepositories(t *testing.T) {
	names, err := GetViewerRepositories()

	if err != nil {
		t.Fatal(err)
	}
	t.Log(names)
}
