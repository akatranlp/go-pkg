package main

import (
	"fmt"
	"io"
)

var streamTemplateStr = `
package {{ .name }}

import (
	"iter"
)

`

type StreamConfig struct {
	Matrix []int
}

func GenStreams(w io.Writer, s StreamConfig, stream StreamType) {
	for _, c := range s.Matrix {
		inTypes := genericTypeNames[:c]
		outTypes := genericTypeNames[c : c+1]
		GenStreamType(w, inTypes, outTypes, stream)
		fmt.Fprintln(w)
	}
}

func GenStreamType(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	fmt.Fprintf(w, "type Stream%s[%s any] %s\n",
		c.InNameSuffix, c.InGenerics, c.InIterType)
	fmt.Fprintln(w)

	// Of
	fmt.Fprintf(w, "func Of%s[%s any](s %s) %s {\n", c.InNameSuffix, c.InGenerics, c.InIterType, c.InStreamType)
	fmt.Fprintf(w, "	return %s(s)", c.InStreamType)
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	// To
	fmt.Fprintf(w, "func (s %s) To() %s {\n", c.InStreamType, c.InIterType)
	fmt.Fprintf(w, "	return %s(s)\n", c.InIterType)
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	if stream == StreamMember {
		// Collect
		fmt.Fprintf(w, "func (s %s) Collect%s[%s any](fn func(%s) %s) %s {\n", c.InStreamType, c.OutNameSuffix, c.OutGenerics, c.InIterType, c.ReturnType, c.ReturnType)
		fmt.Fprintln(w, "	return fn(s.To())")
		fmt.Fprintln(w, "}")
	}

	return nil
}
