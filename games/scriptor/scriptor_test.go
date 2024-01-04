package scriptor_test

import (
	"log"
	"testing"

	"github.com/death12358/digitalopn/games/scriptor"
)

var (
	testScrtiptor = new_scripttor()
	pipe          = testScrtiptor.Cache.Client.Pipeline()
	ctx           = testScrtiptor.Cache.Client.Context()
)

func new_scripttor() *scriptor.Scriptor {
	cacahe, err := scriptor.NewScriptor(&scriptor.Option{
		Host:     "203.66.13.192:6581",
		Port:     6581,
		PoolSize: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cacahe
}

// clear function
func reset() {
	pipe.FlushAll(ctx)
	pipe.Exec(ctx)
	testScrtiptor = new_scripttor()
}

func mockHset(db int, key, vkey, value string) {
	pipe.Do(ctx, "select", db)
	pipe.HSet(ctx, key, vkey, value)
	pipe.Exec(ctx)
}

func TestHello(t *testing.T) {
	res, err := testScrtiptor.Hello()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
