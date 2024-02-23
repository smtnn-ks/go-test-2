package models

//go:generate reform

// Usr represents a row in usr table.
//
//reform:usr
type Usr struct {
	ID      int64  `reform:"id,pk"`
	UsrName string `reform:"usr_name"`
	Pass    string `reform:"pass"`
}
