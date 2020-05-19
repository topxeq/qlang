package flag

import (
	"flag"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "flag",

	"ContinueOnError": flag.ContinueOnError,
	"ExitOnError":     flag.ExitOnError,
	"PanicOnError":    flag.PanicOnError,

	"CommandLine": flag.CommandLine,
	"ErrHelp":     flag.ErrHelp,
	"Usage":       flag.Usage,

	"Arg":           flag.Arg,
	"Args":          flag.Args,
	"Bool":          flag.Bool,
	"BoolVar":       flag.BoolVar,
	"Duration":      flag.Duration,
	"DurationVar":   flag.DurationVar,
	"Float64":       flag.Float64,
	"Float64Var":    flag.Float64Var,
	"Int":           flag.Int,
	"Int64":         flag.Int64,
	"Int64Var":      flag.Int64Var,
	"IntVar":        flag.IntVar,
	"NArg":          flag.NArg,
	"NFlag":         flag.NFlag,
	"Parse":         flag.Parse,
	"Parsed":        flag.Parsed,
	"PrintDefaults": flag.PrintDefaults,
	"Set":           flag.Set,
	"String":        flag.String,
	"StringVar":     flag.StringVar,
	"Uint":          flag.Uint,
	"Uint64":        flag.Uint64,
	"Uint64Var":     flag.Uint64Var,
	"UintVar":       flag.UintVar,
	"UnquoteUsage":  flag.UnquoteUsage,
	"Var":           flag.Var,
	"Visit":         flag.Visit,
	"VisitAll":      flag.VisitAll,

	"Flag":    qlang.StructOf((*flag.Flag)(nil)),
	"Lookup":  flag.Lookup,
	"flagset": flag.NewFlagSet,
}
