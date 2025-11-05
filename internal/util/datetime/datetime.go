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
	count := utf8.RuneCountInString(arr[1])
	nanoseconds, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return err
	}
	if count < 9 {
		nanoseconds = nanoseconds * int64(math.Pow(10, float64(9-count)))
	}
	*t = Local(time.Unix(seconds, nanoseconds).String())
	return nil
}
