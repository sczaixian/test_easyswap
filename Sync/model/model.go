package model

import (
	"context"

	"github.com/ProjectsTask/Base/stores/gdb"
	"gorm.io/gorm"
)

func NewDB(ndb *gdb.Config) *gorm.DB {
	db := gdb.MustNewDB(ndb)
	ctx := context.Background()
	err := InitModel(ctx, db)
	if err != nil {
		panic(err)
	}

	return db
}

func InitModel(ctx context.Context, db *gorm.DB) error {
	err := db.Set(
		"gorm:table_options",
		"ENGINE=InnoDB AUTO_INCREMENT=1 CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci",
	).Error
	if err != nil {
		return err
	}

	return nil
}
