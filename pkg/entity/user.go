package entity

import "time"

type User struct {
	ID          string    `bun:",pk"`
	FirebaseUID string    `bun:"firebase_uid"`
	Name        string    `bun:"name"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
