package casbin

import (
	"authorization/internal/models"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var enforcer *casbin.Enforcer

const TechincalDomain = "feelgoodinc"

func DefineInitialPolicies(domain string) error {
	for role, objects := range rolePolicies {
		for obj, actions := range objects {
			for _, action := range actions {
				if hasPolicy := enforcer.HasPolicy(role, obj, action, domain); !hasPolicy {
					_, err := enforcer.AddPolicy(role, obj, action, domain)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return enforcer.LoadPolicy()
}

func InitCasbin(db *gorm.DB) *casbin.Enforcer {
	//adapter, err := gormadapter.NewAdapterByDB(db)
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &models.Permissions{}, "permissions")
	if err != nil {
		panic(err)
	}
	_enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(err)
	}

	enforcer = _enforcer

	err = enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	//err = DefineInitialPolicies(TechincalDomain)

	if err != nil {
		panic(err)
	}
	fmt.Println("Casbin Loaded")
	return enforcer
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
