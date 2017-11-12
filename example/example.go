package example

import (
	"math/rand"
	"time"

	"github.com/morikuni/dummy"
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
	dummy.Gen
}

// User generate example random user from a key.
func (g *Gen) User(key ...string) *User {
	return g.WithCache(key, (*User)(nil), func() interface{} { return GenUser(g.Rand(key)) }).(*User)
}

// GenUser user generate a example random use from rand.
func GenUser(r *rand.Rand) *User {
	createdAt := genTime(r).Gen()
	return &User{
		ID:          dummy.Int(r).Gen64(),
		Name:        dummy.String(r).Len(8).Gen(),
		ContactInfo: genContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   genTime(r).After(createdAt).Gen(),
	}
}

func genContactInfo(r *rand.Rand) ContactInfo {
	return ContactInfo{
		Email:         genEmail(r),
		EmailVerified: dummy.Bool(r).Gen(),
	}
}

func genEmail(r *rand.Rand) string {
	return dummy.String(r).Gen() + "@" + dummy.String(r).Gen() + ".com"
}

func genTime(r *rand.Rand) *dummy.TimeGen {
	return dummy.Time(r).
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
