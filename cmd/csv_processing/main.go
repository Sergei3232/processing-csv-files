package main

import (
	"fmt"
	"github.com/Sergei3232/processing-csv-files/internal/app/config"
	"github.com/Sergei3232/processing-csv-files/internal/app/queue/postgres"
	"log"
)

const fileInt = "files/int"

func main() {
	configs := config.NenConfig()

	ImageLoader, err := postgres.NewDBConnect(configs.ImageLoader)
	if err != nil {
		log.Panicln(err)
	}

	FileStorage, err := postgres.NewDBConnect(configs.FileStorage)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(ImageLoader, FileStorage)

}
