package pipeline

import (
	"fmt"
	"io"

	"github.com/jonathatargino/fifo-vs-aging/internal/algorithm"
	"github.com/jonathatargino/fifo-vs-aging/internal/generator"
)

type Option func(*Pipeline)

type Pipeline struct {
	Writer     io.Writer
	Algorithms []AlgorithmNamed
	Input      []int
}

type AlgorithmNamed struct {
	algorithm.Algorithm
	Name string
}

func WithAlgoritms(algs ...AlgorithmNamed) Option {
	return func(p *Pipeline) {
		p.Algorithms = algs
	}
}

func WithWriter(w io.Writer) Option {
	return func(p *Pipeline) {
		p.Writer = w
	}
}

func WithGeneratorParams(pagesNumber, referencesLength int) Option {
	return func(p *Pipeline) {
		p.Input = generator.NewPageReferences(pagesNumber, referencesLength)
	}
}

func NewPipeline(opts ...Option) *Pipeline {
	p := &Pipeline{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Pipeline) Run() {
	p.Writer.Write([]byte("Input: "))
	p.Writer.Write(intsToBytes(p.Input))
	p.Writer.Write([]byte("\n"))

	for _, algorithm := range p.Algorithms {
		fmt.Fprintf(p.Writer, "Test case for %s:\n", algorithm.Name)

		// TODO Let dinamic
		result := algorithm.Run(p.Input, 2)

		fmt.Fprintf(p.Writer, "Number of Fault Pages: %d\n", result)
	}
}

func intsToBytes(ints []int) []byte {
	var str string
	for _, v := range ints {
		str += fmt.Sprintf("%d ", v)
	}
	return []byte(str)
}
