package goand

import (
	"fmt"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Goand struct {
	AppName string
	Debug   bool
	Version string
}

func (g *Goand) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migration", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	//read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}
	return nil
}

func (g *Goand) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Goand) checkDotEnv(path string) error {
	err := g.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}
