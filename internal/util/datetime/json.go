package datetime

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Local string

func (t *Local) UnmarshalJSON(b []byte) error {
	arr := strings.Split(string(b), ".")
	seconds, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return err
	}
	var nanoseconds int64
	if len(arr) > 1 {
		nanoseconds, err = strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return err
		}
		if count := utf8.RuneCountInString(arr[1]); count < 9 {
			nanoseconds = nanoseconds * int64(math.Pow(10, float64(9-count)))
		}
	}
	unixTime := time.Unix(seconds, nanoseconds)
	pst := time.FixedZone("UTC-8", -8*60*60)
	*t = Local(unixTime.In(pst).Format(time.RFC1123))
	return nil
}
