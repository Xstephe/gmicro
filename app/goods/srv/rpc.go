package srv

import (
	"fmt"

	gpbv1 "mxshop/api/goods/v1"
	"mxshop/app/goods/srv/config"
	v12 "mxshop/app/goods/srv/internal/controller/v1"
	"mxshop/app/goods/srv/internal/data/v1/db"
	"mxshop/app/goods/srv/internal/data_search/v1/es"
	"mxshop/app/goods/srv/internal/service/v1"
	"mxshop/gmicro/core/trace"
	"mxshop/gmicro/server/rpcserver"
	"mxshop/pkg/log"
)

func NewGoodsRPCServer(cfg *config.Config) (*rpcserver.Server, error) {
	//初始化open-telemetry的exporter
	trace.InitAgent(trace.Options{
		cfg.Telemetry.Name,
		cfg.Telemetry.Endpoint,
		cfg.Telemetry.Sampler,
		cfg.Telemetry.Batcher,
	})
	dataFactory, err := db.GetDBFactoryOr(cfg.MySQLOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	//构建，繁琐 - 工厂模式
	//searchClient, err := es.GetSearchFactoryOr(cfg.EsOptions)
	searchFactory, err := es.GetSearchFactoryOr(cfg.EsOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	//goodsData := db.NewGoods(gormDB)
	//categoryData := db.NewCategorys(gormDB)
	//brandData := db.NewBrands(gormDB)
	//SearchData := es.NewGoodsSearch(searchClient)
	//srv := v1.NewGoodsService(dataFactory, searchFactory)

	srvFactory := v1.NewService(dataFactory, searchFactory)
	goodsServer := v12.NewGoodsServer(srvFactory)

	rpcAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	grpcServer := rpcserver.NewServer(rpcserver.WithAddress(rpcAddr))
	gpbv1.RegisterGoodsServer(grpcServer.Server, goodsServer)

	return grpcServer, nil

}
