package analytics

// On this package we have 2 main keys on redis:
// analytics:hour and analytics:counter

import (
	"github.com/allegro/bigcache/v3"
	"github.com/bsm/redislock"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Dependency struct {
	Memory *bigcache.BigCache
	Redis  *redis.Client
	Locker *redislock.Client
	Bot    *tb.Bot
	Logger *sentry.Client
	DB     *sqlx.DB
}