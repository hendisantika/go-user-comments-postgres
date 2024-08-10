package models

type Comment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
	User    *User  `pg:"rel:has-one" json:"user"`
}

func CreateComment(db *pg.DB, req *Comment) (*Comment, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	comment := &Comment{}

	err = db.Model(comment).
		Relation("User").
		Where("comment.id = ?", req.ID).
		Select()

	return comment, err
}
