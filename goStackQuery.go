package goStackQuery

import (
	"fmt"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

const (
	queryURL string = "http://stackoverflow.com/search?q=[go]+"
)

// Query - function that either outputs a link to the search using switch 0,
// or opens a link via switch 1.
func Query(err error, switcher int) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)

	switch {

	case switcher == 0:
		fmt.Fprintln(os.Stderr, "Search: "+queryURL+replaceSpace(err.Error()))

	case switcher == 1:
		webbrowser.Open(queryURL + err.Error())

	default:
		fmt.Fprintf(os.Stderr, "UNKNOWN switcher (need 0 or 1): %v\n", switcher)
	}

	os.Exit(1)
}

// replace " " to "%20"
func replaceSpace(errStr string) string {
	var result string

	replaceStr := strings.Split(errStr, " ")

	for i := range replaceStr {
		result += replaceStr[i] + "%20"
	}
	result = strings.TrimRight(result, `%20`)

	return result
}
