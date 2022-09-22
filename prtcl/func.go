package prtcl

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Returns Time in Format -> yyyy/mm/dd hh:mm:ss.nnnnnn
func Timestamp() string {

	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%.6s",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		strconv.Itoa(time.Now().Nanosecond()))
}

// service kill function, to stop the service if anything necessary provides an error
//
// Parameters:
//   - `err` : error > the given error is not "nil", the current instance gets killed via os.Exit(1)
func ErrorKill(err error) {

	if err != nil {

		PrintObject(err)

		os.Exit(1)
	}
}
