package pipeline

import (
	"io"

	"github.com/jonathatargino/fifo-vs-aging/internal/algorithm"
	"github.com/jonathatargino/fifo-vs-aging/internal/generator"
)

type Option func(*Pipeline)

type Pipeline struct {
	Writer     io.Writer
	Algorithms []algorithm.Algorithm
	Input      []int
}

func WithAlgoritms(algs ...algorithm.Algorithm) Option {
	return func(p *Pipeline) {
		p.Algorithms = algs
	}
}

func WithWriter(w io.Writer) Option {
	return func(p *Pipeline) {
		p.Writer = w
	}
}

func WithGeneratorParams(numberOfPages, numberOfFrames int) Option {
	return func(p *Pipeline) {
		p.Input = generator.NewPageReferences(numberOfFrames, numberOfFrames)
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
	for _, algorithm := range p.Algorithms {

		p.Writer.Write([]byte("Input: "))
		p.Writer.Write(intsToBytes(p.Input))
		p.Writer.Write([]byte("\n"))

		// TODO Let dinamic
		result := algorithm.Run(p.Input, 10)

		p.Writer.Write([]byte{byte(result)})
	}
}

func intsToBytes(ints []int) []byte {
	bytes := make([]byte, len(ints))
	for i, v := range ints {
		bytes[i] = byte(v)
	}
	return bytes
}
