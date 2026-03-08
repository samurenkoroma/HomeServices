package main

import (
	"context"
	"log"
	"net/http"
	c "samurenkoroma/services/internal/application/bootstrap"
	"samurenkoroma/services/internal/infrastructure/configs"
)

func main() {
	conf := configs.LoadConfig()
	//database := db.NewDb(conf)
	//application := app.NewApplication(conf, database)

	ctx := context.Background()

	app, err := c.Build(ctx, conf.Db.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", app.HTTPHandler); err != nil {
		log.Fatal(err)
	}

}
