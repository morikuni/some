package example

import (
	"math/rand"
	"time"

	"github.com/morikuni/some"
)

// User is a example struct.
type User struct {
	ID          int64
	Name        string
	ContactInfo ContactInfo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ContactInfo is a internal struct for user.
type ContactInfo struct {
	Email         string
	EmailVerified bool
}

// Some is a generator for example.
type Some struct {
	some.Some
}

// User generate example random user from a key.
func (s *Some) User(key string, spec UserSpec) *User {
	u := s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(User)
	return &u
}

type UserSpec struct {
	WithID int64
}

// SomeUser user generate a example random use from rand.
func (s UserSpec) Generate(r *rand.Rand) User {
	t := some.Time.
		After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)).
		Before(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local))
	createdAt := t.Generate(r)
	return User{
		ID:          s.someID(r),
		Name:        some.String.Len(8).Generate(r),
		ContactInfo: someContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   t.After(createdAt).Generate(r),
	}
}

func (s UserSpec) someID(r *rand.Rand) int64 {
	if s.WithID != 0 {
		return s.WithID
	}
	return some.Int64.Generate(r)
}

func someContactInfo(r *rand.Rand) ContactInfo {
	return ContactInfo{
		Email:         someEmail(r),
		EmailVerified: some.Bool.Generate(r),
	}
}

func someEmail(r *rand.Rand) string {
	return some.String.Generate(r) + "@" + some.String.Generate(r) + ".com"
}

// KVS is a example store for user.
type KVS struct {
	store map[int64]*User
}

// Get finds user by id.
func (kvs *KVS) Get(id int64) *User {
	return kvs.store[id]
}
