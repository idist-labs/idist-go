package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"idist-go/app/models"
	"idist-go/app/providers/configProvider"
	"idist-go/app/providers/postgresProvider"
	"idist-go/database/seeds"
)

func DB() *gorm.DB {
	return postgresProvider.GetInstance()
}

func Runner() {
	config := configProvider.GetConfig()
	err := DB().AutoMigrate(
		models.Config{},
		models.IdentifierType{},
		models.Country{},
		models.Province{},
		models.District{},
		models.Ward{},
		models.EmailType{},
		models.Ethnic{},
		models.Gender{},
		models.OrganizationGroup{},
		models.Page{},
		models.PracticingStatus{},
		models.PracticingCertificate{},
		models.User{},
		models.Role{},
		models.Permission{},
		models.UserRole{},
		models.RoleHasPermission{},
		models.UserMail{},
		models.UserAddress{},
		models.UserPhone{},
		models.UserIdentifier{},
		models.UserOrganization{},
		models.PracticeInformation{},
	)

	if err == nil && config.GetBool("postgres.seed") {
		fmt.Println("Seeding data")
		seeds.SeedUsers()
		seeds.SeedPages()
	}
}
