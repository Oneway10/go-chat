package main

import (
	"chat/dal/mysql"
	"context"
	"gorm.io/gen"
)

func main() {
	mysql.Init()
	db := mysql.DB(context.Background())

	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../dal/dao",
		ModelPkgPath:  "../../dal/model",
		Mode:          gen.WithoutContext,
		FieldNullable: true,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("user"))

	g.Execute()
}
