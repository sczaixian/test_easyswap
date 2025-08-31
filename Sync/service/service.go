package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/ProjectsTask/Base/chain"
	"github.com/ProjectsTask/Base/chain/chainclient"
	"github.com/ProjectsTask/Base/ordermanager"
	"github.com/ProjectsTask/Sync/service/orderbookindexer"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"gorm.io/gorm"

	"github.com/ProjectsTask/Base/stores/xkv"
	"github.com/ProjectsTask/Sync/model"
	"github.com/ProjectsTask/Sync/service/collectionfilter"
	"github.com/ProjectsTask/Sync/service/config"

	"github.com/pkg/errors"
)

type Service struct {
	ctx              context.Context
	config           *config.Config
	kvStore          *xkv.Store
	db               *gorm.DB
	wg               *sync.WaitGroup
	collectionFilter *collectionfilter.Filter
	orderbookIndexer *orderbookindexer.Service
	orderManager     *ordermanager.OrderManager
}

func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	var kvConf kv.KvConf
	for _, con := range cfg.Kv.Redis {
		kvConf = append(kvConf, cache.NodeConf{
			RedisConf: redis.RedisConf{
				Host: con.Host,
				Type: con.Type,
				Pass: con.Pass,
			}, Weight: 2,
		})
	}
	kvStore := xkv.NewStore(kvConf)
	var err error
	db := model.NewDB(cfg.DB)
	collectionFilter := collectionfilter.New(ctx, db, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	orderManager := ordermanager.New(ctx, db, kvStore, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	var orderbookSyncer *orderbookindexer.Service
	var chainClient chainclient.ChainClient
	fmt.Println("chainClient url:" + cfg.AnkrCfg.HttpsUrl + cfg.AnkrCfg.ApiKey)
	chainClient, err = chainclient.New(int(cfg.ChainCfg.ID), cfg.AnkrCfg.HttpsUrl+cfg.AnkrCfg.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed on create evm client")
	}

	switch cfg.ChainCfg.ID {
	case chain.EthChainID, chain.OptimismChainID, chain.SepoliaChainID:
		orderbookSyncer = orderbookindexer.New(ctx, cfg, db, kvStore, chainClient, cfg.ChainCfg.ID, cfg.ChainCfg.Name, orderManager)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed on create trade info server")
	}

	manager := Service{
		ctx:              ctx,
		config:           cfg,
		db:               db,
		kvStore:          kvStore,
		collectionFilter: collectionFilter,
		orderbookIndexer: orderbookSyncer,
		orderManager:     orderManager,
		wg:               &sync.WaitGroup{},
	}

	return &manager, nil
}

func (s *Service) Start() error {
	if err := s.collectionFilter.PreloadCollections(); err != nil {
		return errors.Wrap(err, "failed on preload collection to filter")
	}

	s.orderbookIndexer.Start()
	s.orderManager.Start()
	return nil
}
