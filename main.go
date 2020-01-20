package main

import (
	"log"
	"mysql_api/databases"
	"mysql_api/jobs"
	"mysql_api/routers"
)

func foreverGo(run func(), routineLimits int) {
	for i := 0; i < routineLimits; i++ {
		go func() {
			for {
				run()
			}
		}()
	}
}

func main() {
	databases.AutoMigrate()

	foreverGo(jobs.UpdateStatus, 1)

	router := routers.Load()
	if err := router.Run(":9020"); err != nil {
		log.Fatalln(err)
	}
}
