package seeds

import (
	"errors"
	"gorm.io/gorm"
	"idist-go/app/models"
)

func SeedPages() {
	if DB().Migrator().HasTable(&models.Page{}) {
		if err := DB().First(&models.Page{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			page := models.Page{
				Title:       "Tin tức",
				Slug:        "tin-tuc",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book",
				Content:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			}
			_ = page.Create()
		}
	}
}
