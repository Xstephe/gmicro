package admin

import (
	"mxshop/app/user/srv/config"
	"mxshop/gmicro/server/restserver"
)

func NewUserHTTPServer(cfg *config.Config) (*restserver.Server, error) {

	urpcRestServer := restserver.NewServer(
		restserver.WithPort(cfg.Server.HttpPort),
		restserver.WithMiddlewares(cfg.Server.Middlewares),
		restserver.WithEnableProfiling(true),
		restserver.WithMetrics(true),
	)
	//配置好路由
	initRouter(urpcRestServer)
	return urpcRestServer, nil
}