package example

import (
	"math/rand"
	"time"

	"github.com/morikuni/gen"
)

type User struct {
	ID          int64
	Name        string
	ContactInfo ContactInfo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ContactInfo struct {
	Email         string
	EmailVerified bool
}

type Gen struct {
	gen.Gen
}

func (g *Gen) User(key ...string) *User {
	return g.WithCache(key, (*User)(nil), func() interface{} { return GenUser(g.Rand(key)) }).(*User)
}

func GenUser(r *rand.Rand) *User {
	createdAt := *genTime(r).Gen()
	return &User{
		ID:          gen.Int(r).Gen64(),
		Name:        gen.String(r).Len(8).Gen(),
		ContactInfo: genContactInfo(r),
		CreatedAt:   createdAt,
		UpdatedAt:   *genTime(r).After(createdAt).Gen(),
	}
}

func genContactInfo(r *rand.Rand) ContactInfo {
	return ContactInfo{
		Email:         genEmail(r),
		EmailVerified: gen.Bool(r).Gen(),
	}
}

func genEmail(r *rand.Rand) string {
	return gen.String(r).Gen() + "@" + gen.String(r).Gen() + ".com"
}

func genTime(r *rand.Rand) *gen.TimeGen {
	return gen.Time(r).
		After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)).
		Before(time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local))
}

type KVS struct {
	store map[int64]*User
}

func (kvs *KVS) Get(id int64) *User {
	return kvs.store[id]
}
