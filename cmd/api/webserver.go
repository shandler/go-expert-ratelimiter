package main

import (
	"net/http"

	"github.com/cshandler/go-expert-ratelimiter/config"
	"github.com/cshandler/go-expert-ratelimiter/internal/infra/database"
	internalHandler "github.com/cshandler/go-expert-ratelimiter/internal/infra/handler"
	"github.com/cshandler/go-expert-ratelimiter/internal/infra/webserver"
	"github.com/cshandler/go-expert-ratelimiter/internal/infra/webserver/middleware"
	"github.com/redis/go-redis/v9"
)

func CreateWebServer(cfg *config.Config, redisCli *redis.Client) *webserver.WebServer {
	newWebServer := webserver.NewWebServer(cfg.App.Port)
	newWebServer.InternalMiddleware = middleware.Middleware{
		RedisClient: redisCli,
		Config:      cfg,
	}
	apikeyHandler := internalHandler.NewAPIKeyHandler(database.NewAPIKeyRedis(redisCli))

	newWebServer.AddHandler(http.MethodPost, "/api-key", apikeyHandler.CreateAPIKey)
	newWebServer.AddHandler(http.MethodGet, "/hello-world", internalHandler.HelloWorld)
	newWebServer.AddHandler(http.MethodGet, "/hello-world-key", internalHandler.HelloWorldWithAPIKey)

	return newWebServer
}
