package models

// URL represents entity in urls table
type URL struct {
	ID          uint64    `db:"id" json:"id"`
	UserID      NullInt64 `db:"user_id" json:"userId"`
	RedirectURL string    `db:"redirect_url" json:"redirectUrl"`
	EncodedURL  string    `json:"encodedUrl"`
}
