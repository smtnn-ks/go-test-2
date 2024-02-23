package models

//go:generate reform

// News represents a row in news table.
//
//reform:news
type News struct {
	ID    int64  `reform:"id,pk"`
	Title string `reform:"title"`
	Cnt   string `reform:"cnt"`
}
