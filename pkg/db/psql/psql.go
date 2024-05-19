package psql

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.User{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Course{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.UserCourse{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Lecture{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Exercise{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Comment{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.Rating{}); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&domain.AuthorRequest{}); err != nil {
		return err
	}

	return nil
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal("error getting db", err)
	}

	log.Println("closing db")

	err = db.Close()
	if err != nil {
		log.Fatal("error closing db: ", err)
	}

	log.Println("db closed")
}
