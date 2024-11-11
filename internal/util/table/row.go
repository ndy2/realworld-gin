package table

import "time"

type UserRow struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProfileRow struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Bio       string    `db:"bio"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type FollowerRow struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	FollowerID int       `db:"follower_id"`
	CreatedAt  time.Time `db:"created_at"`
}
