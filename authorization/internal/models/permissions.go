package models

// Permissions Casbin Table
// see https://github.com/casbin/gorm-adapter
type Permissions struct {
	BaseModel

	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`

	ModelTime

	//Role         *Role         `gorm:"foreignKey:v0;references:"`
	//Organization *Organization `gorm:"foreignKey:v3"`
}
