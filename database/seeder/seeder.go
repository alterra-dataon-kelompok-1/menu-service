package seeder

import "github.com/alterra-kelompok-1/menu-service/database"

func Seed() {

	conn := database.GetConnection()

	menuTableSeeder(conn)
	// otherTableSeeder(conn)
}
