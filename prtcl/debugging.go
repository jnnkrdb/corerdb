package prtcl

import (
	"encoding/json"
	"io"
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

// --------------------------------------------------
// the debugging functions start here

// the default logging output for DEBUG-MODE
var _debugOutput *log.Logger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)

// set an outputwriter to the debugging log
//
// Parameters:
//   - `w` : io.Writer > contains the writer, to which the output should be send
func SetDebugOutput(w io.Writer) {

	_debugOutput.SetOutput(w)
}

// get the current writers
func GetDebugOutput() io.Writer {

	return _debugOutput.Writer()
}

// if the _DEBUG value is true, the function will print out the
// given structs and variables, else, nothing will happen.
//
// Parameters:
//   - `object` : interface{} > should be listed like PrintObject(interface{}_1, interface{}_2, interface{}_X, ...)
func PrintObject(object ...any) {

	if _DEBUG {

		for _, item := range object {

			_debugOutput.Printf("%#v\n", item)
		}
	}
}

// if the _DEBUG value is true, the function will print out the
// given structs and variables in pretty json-format, else, nothing will happen.
// if an error occurs, the json-pretty-print will be skipped and PrintObject(object ...any) will
// be used.
//
// Parameters:
//   - `object` : interface{} > should be listed like PrintObject(interface{}_1, interface{}_2, interface{}_X, ...)
func PrintJSON(object ...any) {

	if _DEBUG {

		for _, item := range object {

			if iJSON, err := json.MarshalIndent(item, "", "  "); err != nil {

				PrintObject(item, err)

			} else {

				_debugOutput.Printf("\n%s\n", string(iJSON))
			}
		}
	}
}
