package table

import "time"

type FollowerRow struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	FollowerID int       `db:"follower_id"`
	CreatedAt  time.Time `db:"created_at"`
}
