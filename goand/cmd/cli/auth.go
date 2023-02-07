package main

import (
	"fmt"
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
	err = copyDataToFile([]byte("drop table if exists users cascade"), downFile)
	if err != nil {
		exitGraceFuly(err)
	}
	//run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGraceFuly(err)
	}

	// coppy file over
	return nil
}
