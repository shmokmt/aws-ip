package main

import (
	"context"
	"log"
	"os"

	"github.com/shmokmt/aws-ip/internal/gen"
)

func main() {
	if err := gen.Run(context.Background()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
