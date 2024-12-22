package cmd

import (
	"log"

	"github.com/section14/workout-tracker/backend/api"
)

func Dev() {
    err := api.LoadStore()
    if err != nil {
        log.Panic(err)
    }

	api.Serve()
}
