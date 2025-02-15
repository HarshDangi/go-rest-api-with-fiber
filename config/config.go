package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/joho/godotenv"
)

var projectDirName = "go-rest-api-with-fiber"

func Config(key string) string {
	reg := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := reg.Find([]byte(cwd))
	err := godotenv.Load(filepath.Join(string(rootPath), `.env`))

	if err != nil {
		fmt.Print("Error loading. env file")
	}

	return os.Getenv(key)
}
