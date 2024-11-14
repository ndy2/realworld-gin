package table

import "time"

type ProfileRow struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Bio       string    `db:"bio"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
