package models

type Role struct {
	BaseModel
	Title  string  `json:"title"`
	Domain string  `json:"domain"`
	Slug   string  `gorm:"unique" json:"slug"`
	User   *[]User `json:"user"`
}
