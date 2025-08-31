package gdb

import (
	"github.com/ProjectsTask/Base/stores/gdb/orderbookmodel/multi"
)

func GetMultiProjectCollectionTableName(project string, chain string) string {
	if project == OrderBookDexProject {
		return multi.CollectionTableName(chain)
	} else {
		return ""
	}
}

func GetMultiProjectCollectionFloorPriceTableName(project string, chain string) string {
	if project == OrderBookDexProject {
		return multi.CollectionFloorPriceTableName(chain)
	} else {
		return ""
	}
}

func GetMultiProjectItemTableName(project string, chain string) string {
	if project == OrderBookDexProject {
		return multi.ItemTableName(chain)
	} else {
		return ""
	}
}

func GetMultiProjectOrderTableName(project string, chain string) string {
	if project == OrderBookDexProject {
		return multi.OrderTableName(chain)
	} else {
		return ""
	}

}
