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

func (u URL) QueryInt(paramName string) (int, error) {
	query := u.Query().Get(paramName)
	// Checks if query param exists at all via param name. If it does, check for blank value
	if _, ok := u.Query()[paramName]; ok && query == "" {
		log.Print(fmt.Sprintf("Empty %q query param", paramName))
		return 0, ParamError{Key: paramName, Value: query}
	} else if !ok { // Doesn't exist so run "1" as a default thru `Atoi`
		query = "1"
	}
	intQuery, err := strconv.Atoi(query)
	if err != nil {
		log.Printf("%q query param int conversion failed as %d due to %v", paramName, intQuery, err)
		return 0, ParamError{Key: paramName, Value: query}
	}
	return intQuery, nil
}
