package services

type TenantContext struct {
	Id     uint
	Domain string
}

type BaseService struct {
	tenantContext *TenantContext
}
