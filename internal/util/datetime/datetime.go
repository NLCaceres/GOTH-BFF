package datetime

// Convenient datetime formats for use in app t.Format or time.Parse
const AppTime = "Jan 2 2006 3:04:05pm MST"
const LogTime = "2 Jan 2006 15:04:05 Z0700"

// Go uses very specific constants to create datetime formats
// Since they can be a little tough to find, here's every last one.
// const (
//     stdLongMonth      = "January"
//     stdMonth          = "Jan"
//     stdNumMonth       = "1"
//     stdZeroMonth      = "01"
//     stdLongWeekDay    = "Monday"
//     stdWeekDay        = "Mon"
//     stdDay            = "2"
//     stdUnderDay       = "_2"
//     stdZeroDay        = "02"
//     stdHour           = "15"
//     stdHour12         = "3"
//     stdZeroHour12     = "03"
//     stdMinute         = "4"
//     stdZeroMinute     = "04"
//     stdSecond         = "5"
//     stdZeroSecond     = "05"
//     stdLongYear       = "2006"
//     stdYear           = "06"
//     stdPM             = "PM"
//     stdpm             = "pm"
//     stdTZ             = "MST"
//     stdISO8601TZ      = "Z0700"  // prints Z for UTC
//     stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
//     stdNumTZ          = "-0700"  // always numeric
//     stdNumShortTZ     = "-07"    // always numeric
//     stdNumColonTZ     = "-07:00" // always numeric
// )
