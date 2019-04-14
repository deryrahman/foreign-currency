package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/deryrahman/foreign-currency/api/http"
	"github.com/deryrahman/foreign-currency/app"
	"github.com/deryrahman/foreign-currency/app/currency"
	"github.com/deryrahman/foreign-currency/app/rate"
	"github.com/deryrahman/foreign-currency/app/track"
	"github.com/deryrahman/foreign-currency/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// init config
	configuration, err := config.ParseJSON("config.json")
	if err != nil {
		log.Fatalf("couldn't parse config. %s\n", err.Error())
	}
	server := configuration.Server
	database := configuration.DatabaseDev

	// init db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", database.User, database.Password, database.Host, database.Port, database.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("%s\n%s", err.Error(), "failed to connect database")
	}
	defer db.Close()

	// init schema
	db.AutoMigrate(&app.Currency{}, &app.Rate{})
	db.Model(&app.Rate{}).AddForeignKey("currency_id", "currencies(id)", "CASCADE", "CASCADE")

	// init repo
	currencyRepo := currency.CreateRDBMSRepo(db)
	rateRepo := rate.CreateRDBMSRepo(db)

	// init service
	rateService := rate.CreateService(rateRepo, currencyRepo)
	trackService := track.CreateService(rateRepo, currencyRepo)

	// init handler
	handler := h.CreateHTTPHandler(rateService, trackService)

	// ready to serve
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/rates", handler.GetRates).Methods("GET")
	r.HandleFunc("/api/v1/rates", handler.PostRates).Methods("POST")
	r.HandleFunc("/api/v1/tracks", handler.GetTracks).Methods("GET")
	r.HandleFunc("/api/v1/tracks", handler.PostTracks).Methods("POST")
	r.HandleFunc("/api/v1/tracks", handler.DeleteTracks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server.Host, server.Port), r))
}
