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
type Some struct {
	some.Generator
}

// User generate example random user from a key.
func (s *Some) User(key string, spec UserSpec) *User {
	u := s.Generate(key, spec, func(r *rand.Rand) interface{} { return spec.Generate(r) }).(User)
	return &u
}

type UserSpec struct {
	WithID int64
}

func (s UserSpec) CacheKey() string {
	return strconv.FormatInt(s.WithID, 10)
}

// SomeUser user generate a example random use from rand.
func (s UserSpec) Generate(r *rand.Rand) User {
	ts := some.TimeSpec{
		After:  time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local),
		Before: time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local),
	}
	createdAt := ts.Generate(r)
	ts.After = createdAt
	return User{
		ID:          s.someID(r),
		Name:        some.StringSpec{Len: 8}.Generate(r),
		ContactInfo: someContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   ts.Generate(r),
	}
}

func (s UserSpec) someID(r *rand.Rand) int64 {
	if s.WithID != 0 {
		return s.WithID
	}
	return some.AnyInt64.Generate(r)
}

func someContactInfo(r *rand.Rand) ContactInfo {
	return ContactInfo{
		Email:         someEmail(r),
		EmailVerified: some.AnyBool.Generate(r),
	}
}

func someEmail(r *rand.Rand) string {
	return some.AnyString.Generate(r) + "@" + some.AnyString.Generate(r) + ".com"
}

// KVS is a example store for user.
type KVS struct {
	store map[int64]*User
}

// Get finds user by id.
func (kvs *KVS) Get(id int64) *User {
	return kvs.store[id]
}
