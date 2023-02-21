package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"time"
)

func doAuth() error {
	//migration
	dbType := cel.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)
	err := copyFilefromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGraceFuly(err)
	}
	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens cascade;"), downFile)
	if err != nil {
		exitGraceFuly(err)
	}
	//run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGraceFuly(err)
	}

	// coppy file over
	err = copyFilefromTemplate("templates/data/user.go.txt", cel.RootPath+"/data/user.go")
	if err != nil {
		exitGraceFuly(err)
	}

	err = copyFilefromTemplate("templates/data/token.go.txt", cel.RootPath+"/data/token.go")
	//copy over middleware
	err = copyFilefromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGraceFuly(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGraceFuly(err)
	}

	color.Yellow("  - users, tokens, and remember_tokens migration created and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("dont forger to add user and token models in data/models.go and to add appropriate middleware to your routes!")

	return nil
}
