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
