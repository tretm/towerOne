package main

import (
	"fmt"
	"sync"

	httpserver "tower/api/httpServer"
	"tower/api/typesModels"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ini, err := httpserver.NewInit("settings.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := httpserver.New()
	tower := typesModels.NewTower()
	err = server.Init(ini, tower)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	server.Run(&wg)
	wg.Wait()
}
