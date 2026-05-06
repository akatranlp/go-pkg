package main

import (
	"fmt"
	"strconv"
	"strings"
)

type GenConfig struct {
	NameSuffix    string
	InNameSuffix  string
	OutNameSuffix string

	InGenerics     string
	InVars         string
	InVarsWithType []string

	OutGenerics     string
	OutVars         string
	OutVarsWithType []string

	Generics string

	ReturnType string

	AccWithType []string
	DefWithType []string

	Acc string
	Def string

	InSeqType    string
	InIterType   string
	InStreamType string

	OutSeqType    string
	OutIterType   string
	OutStreamType string
}

func CreateConfig(in, out []string, stream StreamType) GenConfig {
	var c GenConfig

	if len(in) > 1 || len(out) > 1 {
		c.NameSuffix = fmt.Sprintf("%d%d", len(in), len(out))
	}
	if len(in) > 1 {
		c.InNameSuffix = strconv.Itoa(len(in))
	}
	var acc []string
	var def []string
	if len(out) > 1 {
		c.OutNameSuffix = strconv.Itoa(len(out))
		for i, v := range out {
			acc = append(acc, fmt.Sprintf("acc%d", i+1))
			c.AccWithType = append(c.AccWithType, fmt.Sprintf("acc%d %s", i+1, v))
			def = append(def, fmt.Sprintf("def%d", i+1))
			c.DefWithType = append(c.DefWithType, fmt.Sprintf("def%d %s", i+1, v))
		}
	} else if len(out) == 1 {
		acc = []string{"acc"}
		c.AccWithType = []string{"acc " + out[0]}
		def = []string{"def"}
		c.DefWithType = []string{"def " + out[0]}
	}
	c.Acc = strings.Join(acc, ", ")
	c.Def = strings.Join(def, ", ")

	c.InGenerics = strings.Join(in, ", ")
	c.InVars = strings.ToLower(c.InGenerics)
	for _, v := range in {
		c.InVarsWithType = append(c.InVarsWithType, fmt.Sprintf("%s %s", strings.ToLower(v), v))
	}

	c.OutGenerics = strings.Join(out, ", ")
	c.OutVars = strings.ToLower(c.OutGenerics)
	for _, v := range out {
		c.OutVarsWithType = append(c.OutVarsWithType, fmt.Sprintf("%s %s", strings.ToLower(v), v))
	}

	c.Generics = fmt.Sprintf("%s, %s", c.InGenerics, c.OutGenerics)

	if len(out) > 1 {
		c.ReturnType = "(" + c.OutGenerics + ")"
	} else {
		c.ReturnType = c.OutGenerics
	}

	iterFmt := "iter.Seq%s[%s]"
	streamFmt := "Stream%s[%s]"

	c.InIterType = fmt.Sprintf(iterFmt, c.InNameSuffix, c.InGenerics)
	c.InStreamType = fmt.Sprintf(streamFmt, c.InNameSuffix, c.InGenerics)

	c.OutIterType = fmt.Sprintf(iterFmt, c.OutNameSuffix, c.OutGenerics)
	c.OutStreamType = fmt.Sprintf(streamFmt, c.OutNameSuffix, c.OutGenerics)

	if stream == Iter {
		c.InSeqType = c.InIterType
		c.OutSeqType = c.OutIterType
	} else {
		c.InSeqType = c.InStreamType
		c.OutSeqType = c.OutStreamType
	}

	return c
}
