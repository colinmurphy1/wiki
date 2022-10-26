package page

import "testing"

type testGetTitle struct {
	usage       string
	pageContent string
	expected    string
}

func TestGetTitleFromHTML(t *testing.T) {
	// create test cases
	testCases := []testGetTitle{
		{
			usage:       "Page without a title",
			pageContent: "<p>This is a page without a title.</p>",
			expected:    "",
		},
		{
			usage:       "Page with a h1 heading",
			pageContent: "<h1>Hello, World!</h1>\n<p>Sometimes I dream of cheese.</p>",
			expected:    "Hello, World!",
		},
		{
			usage:       "Page with a h2 heading",
			pageContent: "<h2>Heading 2</h2>\n<p>Page Content</p>",
			expected:    "Heading 2",
		},
		{
			usage:       "Page with multiple headings",
			pageContent: "<h1>Heading One</h1>\n<h2>Heading Two</h2>\n<h3>Heading Three</h3>\n",
			expected:    "Heading One",
		},
	}

	// Test each testcase, comparing the result with expected result
	for _, test := range testCases {
		if getTitleFromHTML(test.pageContent) != test.expected {
			t.Errorf("Test \"%s\" FAILED, expected: %s", test.usage, test.expected)
			continue
		}
		t.Logf("Test \"%s\" passed", test.usage)
	}
}
