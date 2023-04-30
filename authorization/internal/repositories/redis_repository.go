package repositories

import (
	"authorization/internal/database"
	"authorization/internal/helpers"
	"authorization/internal/models"
	"context"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	connection *redis.Client
	ctx        context.Context
}

func NewCacheRepository() CacheRepository {
	return CacheRepository{
		connection: database.GetRedisConnection(),
		ctx:        context.Background(),
	}
}

func (self *CacheRepository) SetTenantInfo(tenant string, organization *models.Organization) error {
	data := helpers.ToJson(organization)
	_, err := self.connection.Set(self.ctx, "tenant."+tenant, string(data), 0).Result()
	return err
}

func (self *CacheRepository) GetTenantInfo(tenant string) (*models.Organization, error) {
	val, err := self.connection.Get(self.ctx, "tenant."+tenant).Result()
	if err != nil {
		return nil, err
	}
	var organization models.Organization
	helpers.FromJson(&organization, []byte(val))
	return &organization, err
}
