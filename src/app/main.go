package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"visionServiceGo/src/model"
	"visionServiceGo/src/orm"
	"visionServiceGo/src/service"
)

func main() {
	dsn := buildDsnFromEnv()
	fmt.Println(dsn)
	db := orm.NewVisionORM(openDb(dsn))
	app := service.NewService(db)
	err := app.Run(os.Getenv("PORT"))
	if err != nil {
		os.Exit(42)
	}
}

func buildDsnFromEnv() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}

func openDb(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&model.Vision{})
	return db
}
