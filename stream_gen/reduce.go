package main

import (
	"fmt"
	"io"
	"strings"
)

func GenReduceFuncType(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	params := make([]string, 0, len(c.AccWithType)+len(c.InVarsWithType))
	params = append(params, c.AccWithType...)
	params = append(params, c.InVarsWithType...)

	fmt.Fprintf(w, "type Reduce%sFunc[%s any] = func(%s) %s\n",
		c.NameSuffix, c.Generics, strings.Join(params, ", "), c.ReturnType)
	return nil
}

func GenReduceFunc(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) Reduce%s[%s any](fn Reduce%sFunc[%s]) %s {\n",
			c.InSeqType, c.OutNameSuffix, c.OutGenerics, c.NameSuffix, c.Generics, c.ReturnType)
	} else {
		fmt.Fprintf(w, "func Reduce%s[%s any](s %s, fn Reduce%sFunc[%s]) %s {\n",
			c.NameSuffix, c.Generics, c.InSeqType, c.NameSuffix, c.Generics, c.ReturnType)
	}
	for _, v := range c.OutVarsWithType {
		fmt.Fprintf(w, "	var %s\n", v)
	}
	fmt.Fprintf(w, "	for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "		%s = fn(%s, %s)\n", c.OutVars, c.OutVars, c.InVars)
	fmt.Fprintln(w, "	}")
	fmt.Fprintf(w, "	return %s\n", c.OutVars)
	fmt.Fprintln(w, "}")

	return nil
}

func GenFoldFunc(w io.Writer, in, out []string, stream StreamType) error {
	c := CreateConfig(in, out, stream)

	if stream == StreamMember {
		fmt.Fprintf(w, "func (s %s) Fold%s[%s any](%s, fn Reduce%sFunc[%s]) %s {\n",
			c.InSeqType, c.OutNameSuffix, c.OutGenerics, strings.Join(c.DefWithType, ", "), c.NameSuffix, c.Generics, c.ReturnType)
	} else {
		fmt.Fprintf(w, "func Fold%s[%s any](s %s, %s, fn Reduce%sFunc[%s]) %s {\n",
			c.NameSuffix, c.Generics, c.InSeqType, strings.Join(c.DefWithType, ", "), c.NameSuffix, c.Generics, c.ReturnType)
	}
	fmt.Fprintf(w, "	for %s := range s {\n", c.InVars)
	fmt.Fprintf(w, "		%s = fn(%s, %s)\n", c.Def, c.Def, c.InVars)
	fmt.Fprintln(w, "	}")
	fmt.Fprintf(w, "	return %s\n", c.Def)
	fmt.Fprintln(w, "}")

	return nil
}
