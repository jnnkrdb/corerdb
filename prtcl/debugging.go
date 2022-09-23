package prtcl

import (
	"log"
	"os"
)

// to activate the debugging print
// you have to set this value to true
var _DEBUG bool = false

// activate or deactivate the debugging-print option
//
// Parameters:
//   - `active` : bool > to activate -> true; to deactivate -> false
func SetDebugging(active bool) {

	_DEBUG = active
}

// get the current state of the debugging-print option
func GetDebugging() bool {

	return _DEBUG
}

// --------------------------------------------------
// the default output-log functions start here

// current logging for default log output
var Log *log.Logger = log.New(os.Stdout, "[RDB] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

// if the _DEBUG value is true, the function will print out the
// given structs and variables, else, nothing will happen.
//
// Parameters:
//   - `object` : interface{} > should be listed like PrintObject(interface{}_1, interface{}_2, interface{}_X, ...)
func PrintObject(object ...any) {

	if _DEBUG {

		p, f := Log.Prefix(), Log.Flags()

		Log.SetPrefix("[DEBUG] ")

		Log.SetFlags(log.LstdFlags)

		for _, item := range object {

			log.Printf("%#v\n", item)
		}

		Log.SetPrefix(p)

		Log.SetFlags(f)
	}
}
