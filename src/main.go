package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/server"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"

	"github.com/joho/godotenv"
)

func init() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error load .env file: ", err.Error())
		}
	}

	postgres.InitMigration()

	util.CreateAdminUser()
	util.CreateStorageDirectory()

	util.InitClinics()
	util.InitDoctors()
	util.InitStaffs()
	util.InitDays()
	util.InitSchedules()
	util.InitTimes()
	util.InitPatients()
	util.InitSessions()
}

func main() {
	e := server.CreateServer()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := e.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
