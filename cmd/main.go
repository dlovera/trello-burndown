package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dlovera/trello-burndown/pkg/server"
	"github.com/dlovera/trello-burndown/pkg/trello"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

func init() {
	binaryPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(binaryPath)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	go server.Start()
	trello.Start()
}
