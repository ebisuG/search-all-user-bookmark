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
