package main

import (
	"context"

	"github.com/dougdoenges/flexion-coding-challenge/internal/client"
)

func main() {
	ctx := context.Background()

	client.Run(ctx)
}
