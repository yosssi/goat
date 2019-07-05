package context

import "testing"

func TestExcludeFile(t *testing.T) {
	testCases := []struct {
		name            string
		excludes        []Exclude
		filename        string
		expectedExclude bool
	}{
		{
			name: "test exact match algorithm",
			excludes: []Exclude{
				{Pattern: "test1.go", Algorithm: "exact"},
			},
			filename:        "test1.go",
			expectedExclude: true,
		},
		{
			name: "test empty match algorithm",
			excludes: []Exclude{
				{Pattern: "test2.go", Algorithm: ""},
			},
			filename:        "test2.go",
			expectedExclude: true,
		},
		{
			name: "test chaotic chars match algorithm",
			excludes: []Exclude{
				{Pattern: "test2.go", Algorithm: "sdfsfwefw23l22r211>"},
			},
			filename:        "test2.go",
			expectedExclude: true,
		},
		{
			name: "test suffix chars match algorithm",
			excludes: []Exclude{
				{Pattern: "Gen.go", Algorithm: "suffix"},
			},
			filename:        "app_stateGen.go",
			expectedExclude: true,
		},
		{
			name: "test suffix chars match algorithm",
			excludes: []Exclude{
				{Pattern: "Gen.go", Algorithm: "suffix"},
			},
			filename:        "app_stategen.go",
			expectedExclude: false,
		},
		{
			name: "test prefix chars match algorithm",
			excludes: []Exclude{
				{Pattern: "test", Algorithm: "prefix"},
			},
			filename:        "test_app.go",
			expectedExclude: true,
		},
		{
			name: "test prefix chars match algorithm",
			excludes: []Exclude{
				{Pattern: "test", Algorithm: "prefix"},
			},
			filename:        "app_test.go",
			expectedExclude: false,
		},
		{
			name: "test regexp chars match algorithm",
			excludes: []Exclude{
				{Pattern: "^app.go$", Algorithm: "regexp"},
			},
			filename:        "app.go",
			expectedExclude: true,
		},
		{
			name: "test regexp chars match algorithm",
			excludes: []Exclude{
				{Pattern: "^app.go$", Algorithm: "regexp"},
			},
			filename:        "app.go.",
			expectedExclude: false,
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := Watcher{
				Excludes: tc.excludes,
			}
			exclude := w.exclude(tc.filename)
			if exclude != tc.expectedExclude {
				t.Fatalf("[%d] exp: %t got: %t", i, tc.expectedExclude, exclude)
			}
		})
	}
}

func TestExcludeDir(t *testing.T) {
	testData := []struct {
		excludes []Exclude
		dirname  string
		exclude  bool
	}{
		{
			excludes: []Exclude{
				{Pattern: "a/", Algorithm: "exact"},
				{Pattern: "b/*", Algorithm: "exact"},
			},
			dirname: "/n/p/a",
			exclude: true,
		},
		{
			excludes: []Exclude{
				{Pattern: "a/", Algorithm: "exact"},
				{Pattern: "b/*", Algorithm: "exact"},
			},
			dirname: "/n/p/b",
			exclude: true,
		},
		{
			excludes: []Exclude{
				{Pattern: "a/", Algorithm: "exact"},
				{Pattern: "b/*", Algorithm: "exact"},
			},
			dirname: "/n/p/c",
			exclude: false,
		},
		{
			excludes: []Exclude{
				{Pattern: "a/", Algorithm: "exact"},
				{Pattern: "../../b/*", Algorithm: "exact"},
			},
			dirname: "/n/p/../../b",
			exclude: true,
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
