package preload

import (
	"gorm.io/gorm"
)

func PreloadUserRelations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Division", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Department", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, department_name, division_id")
		}).
		Preload("Department.Division", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		})
}
