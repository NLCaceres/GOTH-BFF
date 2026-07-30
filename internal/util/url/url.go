package url

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type URL struct {
	url.URL
}

func New(u url.URL) URL {
	return URL{u}
}

func (u URL) QueryInt(paramName string) (int, error) {
	query := u.Query()
	value := query.Get(paramName)
	// Checks if query param exists at all via param name. If it does, check for blank value
	if _, ok := query[paramName]; ok && value == "" {
		log.Print(fmt.Sprintf("Empty %q query param", paramName))
		return 0, ParamError{Key: paramName, Value: value}
	} else if !ok { // Doesn't exist so run "1" as a default thru `Atoi`
		value = "1"
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("%q query param int conversion failed as %d due to %v", paramName, intValue, err)
		return 0, ParamError{Key: paramName, Value: value}
	}
	return intValue, nil
}

func (u URL) QueryIntOr(paramName string, defaultValue int) int {
	value, err := u.QueryInt(paramName)
	if err != nil {
		return defaultValue
	}
	return value
}
