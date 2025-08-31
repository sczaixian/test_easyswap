package collectionfilter

import (
	"context"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ProjectsTask/Base/stores/gdb"
	"github.com/ProjectsTask/Sync/service/comm"
)

type Filter struct {
	ctx     context.Context
	db      *gorm.DB
	chain   string
	set     map[string]bool
	lock    *sync.RWMutex
	project string
}

func New(ctx context.Context, db *gorm.DB, chain string, project string) *Filter {
	return &Filter{
		ctx:     ctx,
		db:      db,
		chain:   chain,
		set:     make(map[string]bool),
		lock:    &sync.RWMutex{},
		project: project,
	}
}

func (f *Filter) Add(element string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.set[strings.ToLower(element)] = true
}

func (f *Filter) Remove(element string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	delete(f.set, strings.ToLower(element))
}

func (f *Filter) Contains(element string) bool {
	f.lock.RLock()
	defer f.lock.RUnlock()
	_, exists := f.set[strings.ToLower(element)]
	return exists
}

func (f *Filter) PreloadCollections() error {
	var addresses []string
	var err error

	err = f.db.WithContext(f.ctx).Table(gdb.GetMultiProjectCollectionTableName(f.project, f.chain)).
		Select("address").
		Where("floor_price_status = ?", comm.CollectionFloorPriceImported).
		Scan(&addresses).Error

	if err != nil {
		return errors.Wrap(err, "failed on query collections from db")
	}

	for _, address := range addresses {
		f.Add(address)
	}
	return nil
}
