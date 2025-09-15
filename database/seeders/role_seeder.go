package seeders

import (
	"fmt"

	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/utils"
	"gorm.io/gorm"
)

func SeedRoleEntity(db *gorm.DB) {
	roles := []entities.Role{
		{Name: "admin", Description: utils.StringPtr("Administrator with full access")},
		{Name: "user", Description: utils.StringPtr("Regular user with limited access")},
		{Name: "auditor", Description: utils.StringPtr("Read-only auditor")},
	}

	for _, role := range roles {
		var existing entities.Role
		// check if role already exists (avoid duplicate seeding)
		if err := db.Where("name = ?", role.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					fmt.Println("Error seeding role:", err)
				} else {
					fmt.Println("Seeded role:", role.Name)
				}
			} else {
				fmt.Println("Error checking role:", err)
			}
		}
	}
}
