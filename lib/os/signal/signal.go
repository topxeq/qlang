package signal

import (
	"os/signal"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/signal",

	"Ignore":  signal.Ignore,
	"Ignored": signal.Ignored,
	"Notify":  signal.Notify,
	"Reset":   signal.Reset,
	"Stop":    signal.Stop,
}
