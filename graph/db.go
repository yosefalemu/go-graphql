package graph

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=shy-rice-76046207.us-west-2.aws.neon.tech user=neondb_owner password=XbU8Y0eDEwnN dbname=gohasura_db_9704180 sslmode=require options='project=shy-rice-76046207'"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
