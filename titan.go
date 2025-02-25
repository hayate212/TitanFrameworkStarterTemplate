package main

import (
	"log"

	"./app"
	t "github.com/hayate212/TitanFramework"
)

func main() {
	var config t.TitanConfig
	err := t.ConfigLoad(&config)
	if err != nil {
		log.Fatal(err)
	}
	w := t.NewWorker(t.WorkerConfig{
		Address:        config.Address,
		Port:           config.Port,
		MaxRequestSize: config.MaxRequestSize,
	})
	app.Init(w)
	w.Run()
}
