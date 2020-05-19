package log

import (
	"log"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "log",

	"LUTC":          log.LUTC,
	"Ldate":         log.Ldate,
	"Llongfile":     log.Llongfile,
	"Lmicroseconds": log.Lmicroseconds,
	"Lmsgprefix":    log.Lmsgprefix,
	"Lshortfile":    log.Lshortfile,
	"LstdFlags":     log.LstdFlags,
	"Ltime":         log.Ltime,

	"Fatal":     log.Fatal,
	"Fatalf":    log.Fatalf,
	"Fatalln":   log.Fatalln,
	"Flags":     log.Flags,
	"Output":    log.Output,
	"Panic":     log.Panic,
	"Panicf":    log.Panicf,
	"Panicln":   log.Panicln,
	"Prefix":    log.Prefix,
	"Print":     log.Print,
	"Printf":    log.Printf,
	"Println":   log.Println,
	"SetFlags":  log.SetFlags,
	"SetOutput": log.SetOutput,
	"SetPrefix": log.SetPrefix,
	"Writer":    log.Writer,

	"New": log.New,
}
