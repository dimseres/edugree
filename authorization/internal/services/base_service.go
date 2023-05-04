package services

import "authorization/internal/helpers"

//type TenantContext struct {
//	Id     uint
//	Domain string
//	UserId uint
//	Role   string
//}

type BaseService struct {
	tenantContext *helpers.TenantContext
}
