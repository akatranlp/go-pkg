package main

import (
	"fmt"
	"io"
)

func GenFilterFuncType(w io.Writer, in []string, stream StreamType) error {
	c := CreateConfig(in, nil, stream)
	fmt.Fprintf(w, "type Predicate%sFunc[%s any] = func(%s) bool\n",
		c.InNameSuffix, c.InGenerics, c.InGenerics)
	return nil
}

func GenFilterFunc(w io.Writer, in []string, stream StreamType) error {
	c := CreateConfig(in, nil, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) Filter(predicate Predicate%sFunc[%s]) %s {\n",
			c.InSeqType, c.InNameSuffix, c.InGenerics, c.InSeqType)
	} else {
		fmt.Fprintf(w, "func Filter%s[%s any](s %s, predicate Predicate%sFunc[%s]) %s {\n",
			c.InNameSuffix, c.InGenerics, c.InSeqType, c.InNameSuffix, c.InGenerics, c.InSeqType)
	}
	fmt.Fprintf(w, "	return func(yield func(%s) bool) {\n", c.InGenerics)
	fmt.Fprintf(w, "		for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "			if !predicate(%s) { continue }\n", c.InVars)
	fmt.Fprintf(w, "			if !yield(%s) { return }\n", c.InVars)
	fmt.Fprintln(w, "		}")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	return nil
}

func GenAllFunc(w io.Writer, in []string, stream StreamType) error {
	c := CreateConfig(in, nil, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) All(predicate Predicate%sFunc[%s]) bool {\n",
			c.InSeqType, c.InNameSuffix, c.InGenerics)
	} else {
		fmt.Fprintf(w, "func All%s[%s any](s %s, predicate Predicate%sFunc[%s]) bool {\n",
			c.InNameSuffix, c.InGenerics, c.InSeqType, c.InNameSuffix, c.InGenerics)
	}
	fmt.Fprintf(w, "	for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "		if !predicate(%s) { return false }\n", c.InVars)
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	return true")
	fmt.Fprintln(w, "}")
	return nil
}

func GenNoneFunc(w io.Writer, in []string, stream StreamType) error {
	c := CreateConfig(in, nil, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) None(predicate Predicate%sFunc[%s]) bool {\n",
			c.InSeqType, c.InNameSuffix, c.InGenerics)
	} else {
		fmt.Fprintf(w, "func None%s[%s any](s %s, predicate Predicate%sFunc[%s]) bool {\n",
			c.InNameSuffix, c.InGenerics, c.InSeqType, c.InNameSuffix, c.InGenerics)
	}
	fmt.Fprintf(w, "	for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "		if predicate(%s) { return false }\n", c.InVars)
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	return true")
	fmt.Fprintln(w, "}")
	return nil
}

func GenAnyFunc(w io.Writer, in []string, stream StreamType) error {
	c := CreateConfig(in, nil, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) Any(predicate Predicate%sFunc[%s]) bool {\n",
			c.InSeqType, c.InNameSuffix, c.InGenerics)
	} else {
		fmt.Fprintf(w, "func Any%s[%s any](s %s, predicate Predicate%sFunc[%s]) bool {\n",
			c.InNameSuffix, c.InGenerics, c.InSeqType, c.InNameSuffix, c.InGenerics)
	}
	fmt.Fprintf(w, "	for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "		if predicate(%s) { return true }\n", c.InVars)
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "	return false")
	fmt.Fprintln(w, "}")
	return nil
}
