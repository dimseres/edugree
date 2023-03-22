package repositories

import (
	casbin2 "authorization/internal/casbin"
	"github.com/casbin/casbin/v2"
)

type EnforcerRepository struct {
	enforcer *casbin.Enforcer
}

func NewEnforcerRepository() EnforcerRepository {
	return EnforcerRepository{
		enforcer: casbin2.GetEnforcer(),
	}
}

func (self *EnforcerRepository) EnforcePermission(sub string, obj string, act string, domain string) (bool, error) {
	ok, err := self.enforcer.Enforce(sub, obj, act, domain)
	return ok, err
}
