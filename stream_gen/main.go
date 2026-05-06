package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

type M = map[string]any

var mapTemplateStr = `
package {{ .name }}

import (
	"iter"
)

`
var mapTemplate *template.Template

func init() {
	mapTemplate = template.Must(template.New("mapHeader").Parse(mapTemplateStr))
}

func GenIts(packageName string, config Config, stream StreamType) error {
	var err error

	if stream != Iter {
		streamWriter := config.GetWriter("stream")
		if err = mapTemplate.Execute(streamWriter, M{"name": packageName}); err != nil {
			return err
		}
		GenStreams(streamWriter, config.s, stream)
	}

	mapWriter := config.GetWriter("map")
	if err = mapTemplate.Execute(mapWriter, M{"name": packageName}); err != nil {
		return err
	}
	GenMapFuncs(mapWriter, config.m, stream)

	filterWriter := config.GetWriter("filter")
	if err = mapTemplate.Execute(filterWriter, M{"name": packageName}); err != nil {
		return err
	}
	GenFilterFuncs(filterWriter, config.f, stream)

	reduceWriter := config.GetWriter("reduce")
	if err = mapTemplate.Execute(reduceWriter, M{"name": packageName}); err != nil {
		return err
	}
	GenReduceFuncs(reduceWriter, config.r, stream)

	return nil
}

var genericTypeNames = []string{"T", "U", "V", "W"}

func GenMapFuncs(w io.Writer, m MapConfig, stream StreamType) error {
	for _, c := range m.Matrix {
		inTypesC, outTypesC := c[0], c[1]
		inTypes := genericTypeNames[:inTypesC]
		outTypes := genericTypeNames[inTypesC : inTypesC+outTypesC]
		GenMapFuncType(w, inTypes, outTypes, stream)
	}

	for _, c := range m.Matrix {
		fmt.Fprintln(w)
		inTypesC, outTypesC := c[0], c[1]
		inTypes := genericTypeNames[:inTypesC]
		outTypes := genericTypeNames[inTypesC : inTypesC+outTypesC]
		GenMapFunc(w, inTypes, outTypes, stream)
		fmt.Fprintln(w)

		GenFlatMapFunc(w, inTypes, outTypes, stream)
	}
	return nil
}

func GenReduceFuncs(w io.Writer, r ReduceConfig, stream StreamType) error {
	for _, c := range r.Matrix {
		inTypesC, outTypesC := c[0], c[1]
		inTypes := genericTypeNames[:inTypesC]
		outTypes := genericTypeNames[inTypesC : inTypesC+outTypesC]
		GenReduceFuncType(w, inTypes, outTypes, stream)
	}

	for _, c := range r.Matrix {
		fmt.Fprintln(w)
		inTypesC, outTypesC := c[0], c[1]
		inTypes := genericTypeNames[:inTypesC]
		outTypes := genericTypeNames[inTypesC : inTypesC+outTypesC]
		GenReduceFunc(w, inTypes, outTypes, stream)
		fmt.Fprintln(w)

		GenFoldFunc(w, inTypes, outTypes, stream)
	}
	return nil
}

func GenFilterFuncs(w io.Writer, f FilterConfig, stream StreamType) error {
	for _, c := range f.Matrix {
		inTypes := genericTypeNames[:c]
		GenFilterFuncType(w, inTypes, stream)
	}

	for _, c := range f.Matrix {
		fmt.Fprintln(w)
		inTypes := genericTypeNames[:c]
		GenFilterFunc(w, inTypes, stream)
		fmt.Fprintln(w)

		GenAllFunc(w, inTypes, stream)
		fmt.Fprintln(w)

		GenAnyFunc(w, inTypes, stream)
		fmt.Fprintln(w)

		GenNoneFunc(w, inTypes, stream)
	}
	return nil
}

type StreamType int

const (
	Iter = iota
	Stream
	StreamMember
)

type MapConfig struct {
	Matrix [][2]int
}

type FilterConfig struct {
	Matrix []int
}

type ReduceConfig struct {
	Matrix [][2]int
}

type Config struct {
	GetWriter func(string) io.Writer
	m         MapConfig
	f         FilterConfig
	r         ReduceConfig
	s         StreamConfig
}

func main() {
	var err error

	var buf strings.Builder
	var getWriter = func(name string) io.Writer {
		buf.WriteString("--------------------")
		buf.WriteString(name)
		return &buf
	}

	// config := Config{
	// 	m:         MapConfig{Matrix: [][2]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
	// 	f:         FilterConfig{Matrix: []int{1, 2}},
	// 	r:         ReduceConfig{Matrix: [][2]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
	// 	s:         StreamConfig{Matrix: []int{1, 2}},
	// 	GetWriter: getWriter,
	// }

	config := Config{
		m:         MapConfig{Matrix: [][2]int{{1, 1}, {2, 1}}},
		f:         FilterConfig{Matrix: []int{1, 2}},
		r:         ReduceConfig{Matrix: [][2]int{{1, 1}, {2, 1}}},
		s:         StreamConfig{Matrix: []int{1, 2}},
		GetWriter: getWriter,
	}

	if err = GenIts("main", config, StreamMember); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
