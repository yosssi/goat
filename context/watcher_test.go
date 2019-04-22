package context

import "testing"

func TestExcludeDir(t *testing.T) {
	testData := []struct {
		excludes []string
		dirname  string
		exclude  bool
	}{
		{
			excludes: []string{"a/", "b/*"},
			dirname:  "/n/p/a",
			exclude:  true,
		},
		{
			excludes: []string{"a/", "b/*"},
			dirname:  "/n/p/b",
			exclude:  true,
		},
		{
			excludes: []string{"a/", "b/*"},
			dirname:  "/n/p/c",
			exclude:  false,
		},
		{
			excludes: []string{"a/", "../../b/*"},
			dirname:  "/n/p/../../b",
			exclude:  true,
		},
	}
	for i, td := range testData {
		w := Watcher{
			Excludes: td.excludes,
		}
		exclude := w.excludeDir(td.dirname)
		if td.exclude != exclude {
			t.Fatalf("[%d] exp: %t got: %t", i, td.exclude, exclude)
		}
	}
}
