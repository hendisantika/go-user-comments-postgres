package models

type Comment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
	User    *User  `pg:"rel:has-one" json:"user"`
}
