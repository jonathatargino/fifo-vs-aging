package main

import (
	"os"

	"github.com/jonathatargino/fifo-vs-aging/internal/algorithm"
	"github.com/jonathatargino/fifo-vs-aging/internal/pipeline"
)

func main() {
	alging := algorithm.NewAging()
	fifo := algorithm.NewFifo()

	fileResult, err := os.Open("cmd/result.txt")
	if err != nil {
		panic(err)
	}

	pipeline.NewPipeline(
		pipeline.WithAlgoritms(alging, fifo),
		pipeline.WithGeneratorParams(10, 2),
		pipeline.WithWriter(fileResult),
	).Run()
}
