package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {

	switch arg2 {
	case "migration":
		dbType := cel.DB.DataType
		if arg3 == "" {
			exitGraceFuly(errors.New("you must give the migration a name"))
		}
		filename := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := cel.RootPath + "/migrations/" + filename + "." + dbType + ".up.sql"
		downFile := cel.RootPath + "/migrations/" + filename + "." + dbType + ".down.sql"

		err := copyFilefromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGraceFuly(err)
		}

		err = copyFilefromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGraceFuly(err)
		}
	case "auth":
		err := doAuth()
		if err != nil {
			exitGraceFuly(err)
		}

	}
	return nil
}
