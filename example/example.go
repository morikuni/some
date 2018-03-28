package example

import (
	"math/rand"
	"strconv"
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

// Generator is a generator for example.
type Generator struct {
	some.Generator
}

// User generate example random user from a key.
func (s *Generator) User(key string, spec UserSpec) *User {
	u := s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(User)
	return &u
}

type Some struct {
	some.Some
}

func (s Some) User() UserSpec {
	return UserSpec{}
}

type UserSpec struct {
	ID int64
}

func (s UserSpec) WithID(id int64) UserSpec {
	s.ID = id
	return s
}

func (s UserSpec) CacheKey() string {
	return strconv.FormatInt(s.ID, 10)
}

// SomeUser user generate a example random use from rand.
func (s UserSpec) Generate(r *rand.Rand) User {
	ts := some.Time.
		WithAfter(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)).
		WithBefore(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local))
	createdAt := ts.Generate(r)
	return User{
		ID:          s.someID(r),
		Name:        some.String.WithLen(8).Generate(r),
		ContactInfo: someContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   ts.WithAfter(createdAt).Generate(r),
	}
}

func (s UserSpec) someID(r *rand.Rand) int64 {
	if s.ID != 0 {
		return s.ID
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
