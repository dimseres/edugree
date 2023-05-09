package repositories

import (
	"gorm.io/gorm"
)

type BaseRepositoryHelpers struct {
	db          *gorm.DB
	requestUuid string
}

type PaginationConfig struct {
	Page    int
	PerPage int
}

func (self *BaseRepositoryHelpers) LoadRelation(model interface{}, relation ...string) (interface{}, error) {
	tx := self.db
	for _, relate := range relation {
		tx = tx.Preload(relate)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	res := tx.First(model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model, nil
}

func (self *BaseRepositoryHelpers) StartTransaction() {
	self.db = self.db.Begin()
}

func (self *BaseRepositoryHelpers) EndTransaction() {
	self.db = self.db.Commit()
}

func (self *BaseRepositoryHelpers) RollbackTransaction() {
	self.db = self.db.Rollback()
}

func (self *BaseRepositoryHelpers) GetDb() *gorm.DB {
	return self.db
}

func (self BaseRepositoryHelpers) Paginate(config *PaginationConfig) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := config.Page
		if page <= 1 {
			page = 1
		}
		perPage := config.PerPage
		switch {
		case perPage > 100:
			perPage = 100
		case perPage < 10:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
