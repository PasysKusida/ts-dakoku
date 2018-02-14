package app

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type Context struct {
	RedisConn              redis.Conn
	Request                *http.Request
	ClientSecret           string
	ClientID               string
	UserID                 string
	StateStoreKey          string
	TokenStoreKey          string
	TeamSpiritHost         string
	SlackVerificationToken string
}

func (app *App) CreateContext(r *http.Request) *Context {
	ctx := &Context{
		RedisConn:              app.RedisConn,
		ClientID:               app.ClientID,
		ClientSecret:           app.ClientSecret,
		StateStoreKey:          app.StateStoreKey,
		TokenStoreKey:          app.TokenStoreKey,
		TeamSpiritHost:         app.TeamSpiritHost,
		SlackVerificationToken: app.SlackVerificationToken,
		Request:                r,
	}
	return ctx
}

func (ctx *Context) getVariableInHash(hashKey string, key string) string {
	res, err := ctx.RedisConn.Do("HGET", hashKey, key)
	if err != nil {
		return ""
	}
	if data, ok := res.([]byte); ok {
		return string(data)
	}
	return ""
}
