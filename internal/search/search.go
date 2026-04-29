package search

type Searcher interface {
	Search(bookmarks Bookmarks, keyword string) (Bookmarks, error)
}

type Parser interface {
	Parse(path string) ([]Bookmark, error)
}

type Bookmarks []Bookmark

type Bookmark struct {
	BookmarkTitle BookmarkTitle
	BookmarkUrl   BookmarkUrl
}

type BookmarkTitle struct {
	Record Record
}

type BookmarkUrl struct {
	Record Record
}

type Record struct {
	Raw  string
	Norm string
}
