package helpers

type CacheService struct {
	cache map[string]map[string]interface{}
}

func NewCacheService() CacheService {
	return CacheService{}
}

func (self *CacheService) GetData(group string, key string) interface{} {
	grp := self.getGroup(group)
	if grp == nil {
		return nil
	}
	data, err := grp[key]
	if err {
		return nil
	}
	return data
}

func (self *CacheService) getGroup(group string) map[string]interface{} {
	grp, err := self.cache[group]
	if err {
		return nil
	}
	return grp
}

func (self *CacheService) SetData(group string, key string, value interface{}) bool {
	grp, err := self.cache[group]
	if err {
		self.cache[group] = map[string]interface{}{
			key: value,
		}
		return true
	}
	_, err = grp[key]
	if err {
		self.cache[group][key] = value
	}
	return true
}

type Error struct {
	error string
}

func (self *Error) Error() string {
	return self.error
}

func (self *CacheService) FlushKey(group string, key string) error {
	grp := self.getGroup(group)
	if grp == nil {
		return &Error{"group doesn't exist"}
	}
	delete(grp, key)
	return nil
}
