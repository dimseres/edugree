package repositories

import (
	"gorm.io/gorm"
)

type BaseRepositoryHelpers struct {
	db *gorm.DB
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
