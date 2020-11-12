package main

import (
	"fruitshop/internal/httpsrv"
	"log"
	"time"
)

func main() {
	start := time.Now()

	s, err := httpsrv.New("sqlite3", "fruitshop.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	if err := s.Run(":8080", start); err != nil {
		log.Fatalln(err)
	}
}
