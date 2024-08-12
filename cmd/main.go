package main

import (
	"os"

	"github.com/jonathatargino/fifo-vs-aging/internal/algorithm"
	"github.com/jonathatargino/fifo-vs-aging/internal/pipeline"
)

func main() {
	alging := algorithm.NewAging()
	// fifo := algorithm.NewFifo()

	algingNamed := pipeline.AlgorithmNamed{Algorithm: alging, Name: "Aging"}

	file, err := os.Create("cmd/result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pipeline.NewPipeline(
		pipeline.WithAlgoritms(algingNamed),
		pipeline.WithGeneratorParams(10, 20),
		pipeline.WithWriter(file),
	).Run()
}
