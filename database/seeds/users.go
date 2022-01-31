package seeds

import (
	"errors"
	"gorm.io/gorm"
	"idist-go/app/models"
)

func SeedUsers() {
	if DB().Migrator().HasTable(&models.User{}) {
		if err := DB().First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
		}
	}
}
