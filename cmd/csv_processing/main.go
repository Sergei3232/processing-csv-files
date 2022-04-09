package main

import (
	"fmt"
	"github.com/Sergei3232/processing-csv-files/internal/app/structTags"
)

func main() {

	config := structTags.NenConfig()
	fmt.Println(config)

}
