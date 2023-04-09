package models

import "time"

// URL represents entity in urls table
type URL struct {
	ID          uint64    `db:"id"`
	RedirectURL string    `db:"redirect_url"`
	CreatedAt   time.Time `db:"created_at"`
}
