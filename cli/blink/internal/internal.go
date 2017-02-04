package internal

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

// Check checks error, if != nil, throw panic
func Check(e error) {
	if e != nil {
		log.Fatalf("%s", e)
	}
}

// Exit func display an error message on stderr and exit 1
func Exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
