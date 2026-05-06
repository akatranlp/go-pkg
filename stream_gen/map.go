package main

import (
	"fmt"
	"io"
)

func GenMapFuncType(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)
	fmt.Fprintf(w, "type Map%sFunc[%s any] = func(%s) %s\n",
		c.NameSuffix, c.Generics, c.InGenerics, c.ReturnType)
	return nil
}

func GenMapFunc(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) Map%s[%s any](fn Map%sFunc[%s]) %s {\n",
			c.InSeqType, c.OutNameSuffix, c.OutGenerics, c.NameSuffix, c.Generics, c.OutSeqType)
	} else {
		fmt.Fprintf(w, "func Map%s[%s any](s %s, fn Map%sFunc[%s]) %s {\n",
			c.NameSuffix, c.Generics, c.InSeqType, c.NameSuffix, c.Generics, c.OutSeqType)
	}
	fmt.Fprintf(w, "	return func(yield func(%s) bool) {\n", c.OutGenerics)
	fmt.Fprintf(w, "		for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "			if !yield(fn(%s)) { return }\n", c.InVars)
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	return nil
}

func GenFlatMapFunc(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	var extraMapFuncName string
	if len(in) > 1 {
		extraMapFuncName = fmt.Sprintf("%d1", len(in))
	}

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) FlatMap%s[%s any](fn Map%sFunc[%s, %s]) %s {\n",
			c.InSeqType, c.OutNameSuffix, c.OutGenerics, extraMapFuncName, c.InGenerics, c.OutSeqType, c.OutSeqType)
	} else {
		fmt.Fprintf(w, "func FlatMap%s[%s any](s %s, fn Map%sFunc[%s, %s]) %s {\n",
			c.NameSuffix, c.Generics, c.InSeqType, extraMapFuncName, c.InGenerics, c.OutSeqType, c.OutSeqType)
	}
	fmt.Fprintf(w, "	return func(yield func(%s) bool) {\n", c.OutGenerics)
	fmt.Fprintf(w, "		for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "			for %s := range fn(%s) {\n", c.OutVars, c.InVars)
	fmt.Fprintf(w, "				if !yield(%s) { return }\n", c.OutVars)
	fmt.Fprintln(w, "			}")
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	return nil
}
