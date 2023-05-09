package repositories

import (
	"authorization/internal/database"
	"authorization/internal/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	tenant *models.Organization
	db     *gorm.DB
}

func NewCourseRepository(organization *models.Organization) CourseRepository {
	return CourseRepository{
		tenant: organization,
		db:     database.GetConnection(),
	}
}

func (c *CourseRepository) CreateTenantCourse(userId uint) {

}
