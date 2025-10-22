package preload

import (
	"gorm.io/gorm"
)

func PreloadDepartmentRelations(db *gorm.DB) *gorm.DB {
	return db.Preload("Division", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	})
}
