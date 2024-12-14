package main

import (
	"chat/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao",
		ModelPkgPath: "./model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.ServerHost,
		config.Mysql.ServerPort,
		config.Mysql.Database)
	gormdb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	g.UseDB(gormdb)

	allModel := g.GenerateModelAs("user", "User")
	g.ApplyBasic(allModel)
	g.Execute()
}
