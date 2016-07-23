package trace

import (
	"fmt"
	"io"
)
// Tracer is the interface that describes an object capable of
// tracing events throughtout code.

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte([]byte("\n")))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func(t *nilTracer) Trace(a ...interface{}) {}

// Off create a Tracer that will ignore calls to Trace
func Off() Tracer {
	return &nilTracer{}
}
