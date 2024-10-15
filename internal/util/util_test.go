package util

import (
	"os"
	"testing"
)

func TestGetAllBookmarkFilePath(t *testing.T) {
	username := os.Getenv("UserNameForTest")
	expect := []string{"C:\\Users\\" + username + "\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 1\\Bookmarks",
		"C:\\Users\\" + username + "\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 2\\Bookmarks",
		"C:\\Users\\" + username + "\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 3\\Bookmarks"}
	actual := GetAllBookmarkFilePath()
	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			t.Fatalf(`Path doesn't match. expect : %v, actual : %v`, expect[i], actual[i])
		}
	}
}

func TestGetPathName(t *testing.T) {
	username := os.Getenv("UserNameForTest")
	expect := "C:\\Users\\" + username + "\\AppData\\Local\\Google\\Chrome\\User Data"
	actual := GetPathName()
	if expect != actual {
		t.Fatalf(`Path doesn't match. expect : %v, actual : %v`, expect, actual)
	}
}

func TestFilterByString(t *testing.T) {
	actual := ReadBookmarkFile("./testData/Bookmarks")
	if FilterByString(actual, "https://go.dev/") == nil {
		t.Fatalf(`Failed to filter, no https://go.dev/`)
	}
	if len(FilterByString(actual, "")) != 11 {
		t.Fatalf(`Failed to filter, the case there is no char`)
	}
	if len(FilterByString(actual, "org")) != 4 {
		t.Fatalf(`Failed to filter, counting org in the wrong way`)
	}
	if len(FilterByString(actual, "シャトレーゼ")) != 1 {
		t.Fatalf(`Failed to filter, counting シャトレーゼ in the wrong way`)
	}
	if len(FilterByString(actual, "～本格・")) != 1 {
		t.Fatalf(`Failed to filter, counting ～本格・ in the wrong way`)
	}
}
