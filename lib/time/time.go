package time

import (
	"time"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
var Exports = map[string]interface{}{
	"_name": "time",

	"ANSIC":       time.ANSIC,
	"April":       time.April,
	"August":      time.August,
	"December":    time.December,
	"February":    time.February,
	"Friday":      time.Friday,
	"Hour":        time.Hour,
	"January":     time.January,
	"July":        time.July,
	"June":        time.June,
	"Kitchen":     time.Kitchen,
	"March":       time.March,
	"May":         time.May,
	"Microsecond": time.Microsecond,
	"Millisecond": time.Millisecond,
	"Minute":      time.Minute,
	"Monday":      time.Monday,
	"Nanosecond":  time.Nanosecond,
	"November":    time.November,
	"October":     time.October,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RubyDate":    time.RubyDate,
	"Saturday":    time.Saturday,
	"Second":      time.Second,
	"September":   time.September,
	"Stamp":       time.Stamp,
	"StampMicro":  time.StampMicro,
	"StampMilli":  time.StampMilli,
	"StampNano":   time.StampNano,
	"Sunday":      time.Sunday,
	"Thursday":    time.Thursday,
	"Tuesday":     time.Tuesday,
	"UnixDate":    time.UnixDate,
	"Wednesday":   time.Wednesday,

	"Duration": qlang.StructOf((*time.Duration)(nil)),

	"Local": time.Local,
	"UTC":   time.UTC,

	"After": time.After,
	"Sleep": time.Sleep,
	"Tick":  time.Tick,

	"FixedZone":              time.FixedZone,
	"LoadLocation":           time.LoadLocation,
	"LoadLocationFromTZData": time.LoadLocationFromTZData,
	"NewTicker":              time.NewTicker,
	"Date":                   time.Date,
	"now":                    time.Now,
	"Now":                    time.Now,
	"Parse":                  time.Parse,
	"ParseInLocation":        time.ParseInLocation,
	"uniUnixx":               time.Unix,
	"unix":                   time.Unix,
	"Unix":                   time.Unix,
	"UnixMilli":              time.UnixMilli,
	"UnixMicro":              time.UnixMicro,
	"NewTimer":               time.NewTimer,
	"AfterFunc":              time.AfterFunc,
}
