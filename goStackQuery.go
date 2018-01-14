package goStackQuery

import (
	"fmt"
	"net/http"
	"os"

	"github.com/toqueteos/webbrowser"
)

const (
	queryURL string = "http://stackoverflow.com/search?q=[go]+"
)

// Query -
func Query(err error, switcher int) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)

	switch {

	case switcher == 0:
		resp, err := http.Get(queryURL + err.Error())
		if err != nil {
			fmt.Fprintf(os.Stderr, "http.Get ERROR: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

	case switcher == 1:
		webbrowser.Open(queryURL + err.Error())

	default:
		fmt.Fprintf(os.Stderr, "UNCKNOWN switcher (need 0 or 1): %v\n", switcher)
		os.Exit(1)
	}

	os.Exit(1)
}