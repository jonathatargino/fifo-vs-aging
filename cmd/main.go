package main

import (
	"os"

	"github.com/jonathatargino/fifo-vs-aging/internal/algorithm"
	"github.com/jonathatargino/fifo-vs-aging/internal/pipeline"
)

func main() {
	alging := algorithm.NewAging()
	fifo := algorithm.NewFifo()

	algingNamed := pipeline.AlgorithmNamed{Algorithm: alging, Name: "Aging"}
	fifoNamed := pipeline.AlgorithmNamed{Algorithm: fifo, Name: "Fifo"}

	file, err := os.Create("cmd/result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pipeline.NewPipeline(
		pipeline.WithAlgoritms(algingNamed, fifoNamed),
		pipeline.WithGeneratorParams(5, 10),
		pipeline.WithWriter(file),
	).Run()
}
