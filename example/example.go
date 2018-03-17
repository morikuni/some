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

// Gen is a generator for example.
type Gen struct {
	some.Generator
}

// User generate example random user from a key.
func (g *Gen) User(key ...string) *User {
	u := g.WithCache(key, User{}, func() interface{} { return GenUser(g.Rand(key)) }).(User)
	return &u
}

// GenUser user generate a example random use from rand.
func GenUser(r *rand.Rand) User {
	createdAt := genTime(r).Gen()
	return User{
		ID:          some.Int(r).Gen64(),
		Name:        some.String(r).Len(8).Gen(),
		ContactInfo: genContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   genTime(r).After(createdAt).Gen(),
	}
}

func genContactInfo(r *rand.Rand) ContactInfo {
	return ContactInfo{
		Email:         genEmail(r),
		EmailVerified: some.Bool(r).Gen(),
	}
}

func genEmail(r *rand.Rand) string {
	return some.String(r).Gen() + "@" + some.String(r).Gen() + ".com"
}

func genTime(r *rand.Rand) *some.SomeTime {
	return some.Time(r).
		After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)).
		Before(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local))
}

// KVS is a example store for user.
type KVS struct {
	store map[int64]*User
}

// Get finds user by id.
func (kvs *KVS) Get(id int64) *User {
	return kvs.store[id]
}
