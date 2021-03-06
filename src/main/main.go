package main

import (
	"github.com/DapperBlondie/movie-api/src/repo"
	zerolog "github.com/rs/zerolog/log"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const MONGO_DSN = "localhost"

func main() {
	myMongo, err := repo.CreateSession(MONGO_DSN)
	if err != nil {
		zerolog.Fatal().Msg(err.Error())
		return
	}
	defer myMongo.MSession.Close()
	myMongo.AddDataBase("appdb")
	myMongo.AddCollection("movies")

	NewConfig(myMongo)

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           chiRoutes(),
		ReadTimeout:       time.Second * 8,
		ReadHeaderTimeout: time.Second * 5,
		WriteTimeout:      time.Second * 8,
		IdleTimeout:       time.Second * 6,
	}

	zerolog.Log().Msg("Listening and Serving on localhost:8080 ...")
	err = srv.ListenAndServe()
	if err != nil {
		zerolog.Fatal().Msg(err.Error())
		return
	}
	return
}
