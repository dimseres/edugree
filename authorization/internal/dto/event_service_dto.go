package dto

import "authorization/internal/models"

type UserUpdate struct {
	UserMembership *[]models.Membership
	User           *models.User
}

type UserMembershipUpdate struct {
	Membership *models.Membership
	User       *models.User
}
