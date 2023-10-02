package scriptor

import (
	"github.com/go-redis/redis/v8"
	"github.com/yshengliao/goscriptor"
)

// 以roundid作儲存
// 開局檢查是否存在
// 不存在則建立

// 當玩家需要選擇時，Position 不為 0

type Option goscriptor.Option

const (
	_scriptDefinition       = "games_scriptor|1.2.202302221649"
	_PROBSCRIPTINFO         = "PROBSCRIPTINFO"
	_PROBSCRIPTINFOTEMPLATE = `
	-- probscriptor
	-- script use database: 1
		`
)

type Scriptor struct {
	Cache *goscriptor.Scriptor
}

// NewScriptor create a new ProbScriptor
func NewScriptor(opts *Option) (*Scriptor, error) {
	scripts := map[string]string{
		_hello:                  _HelloworldTemplate,
		_round_next:             _round_next_Template,
		_round_update:           _round_update_Template,
		_jackpot_sweep:          _jackpot_sweep_Template,
		_jackpot_push:           _jackpot_push_Template,
		_jackpot_peek:           _jackpot_peek_Template,
		_jackpot_peeks:          _jackpot_peeks_Template,
		_jackpot_currency_peeks: _jackpot_currency_peeks_Template,
	}

	redis_client := redis.NewClient(&redis.Options{
		Addr:     opts.Host,
		Password: opts.Password,
		DB:       opts.DB,
		PoolSize: 10,
	})
	scriptor, err := goscriptor.New(redis_client, 1, _scriptDefinition, &scripts)
	if err != nil {
		return nil, err
	}

	return &Scriptor{
		Cache: scriptor,
	}, nil
}

var (
	_hello              = "hello"
	_HelloworldTemplate = `
	return 'Hello, World!!!'
    `
)

// hello function
func (s *Scriptor) Hello() (string, error) {
	res, err := s.Cache.ExecSha(_hello, []string{})
	if err != nil {
		return "", err
	}

	return res.(string), nil
}
