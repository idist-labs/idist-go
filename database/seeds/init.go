package seeds

import (
	"gorm.io/gorm"
	"idist-go/app/providers/postgresProvider"
)

func DB() *gorm.DB {
	return postgresProvider.GetInstance()
}
