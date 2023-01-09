package data

import "time"

type Token struct {
	ID        int       `db:"id" json:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	FirstName string    `json:"first_name" db:"first_name"`
	Email     string    `json:"email" db:"email"`
	PlainText string    `json:"token" db:"-"`
	Hash      []byte    `json:"-" db:"token_hash"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Expires   time.Time `json:"expiry" db:"expiry"`
}

func (u *Token) Table() string {
	return "tokens"
}
