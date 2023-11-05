package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=db.etsbbvehqrcwtxbuject.supabase.co user=postgres password=Saha@6462++ dbname=postgres port=5432"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	dbSQL.Close()
}
