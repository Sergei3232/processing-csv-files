package main

import (
	"fmt"
	"github.com/Sergei3232/processing-csv-files/internal/app/structTags"
	"github.com/Sergei3232/processing-csv-files/internal/postgres"
	"log"
)

func main() {
	config := structTags.NenConfig()

	ImageLoader, err := postgres.NewDBConnect(config.ImageLoader)
	if err != nil {
		log.Panicln(err)
	}

	FileStorage, err := postgres.NewDBConnect(config.FileStorage)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(ImageLoader, FileStorage)

}
