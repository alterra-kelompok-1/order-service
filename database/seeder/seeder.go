package seeder

import (
	"gorm.io/gorm"
)

func Seed(conn *gorm.DB) {
	orderStatusesTableSeeder(conn)
	orderItemStatusesTableSeeder(conn)
}
