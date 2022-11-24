package fnc

import (
	"regexp"

	"github.com/jnnkrdb/corerdb/prtcl"
)

// compare a string with a list of regexp-like strings, if the
// list contains an expression, which matches with the given string, the
// return value will be true
//
// if an error occurs, the returnvalue will be false
//
// Parameters:
//   - `compare` : string > string, which should be compared to the list
//   - `listOfRegexp` : []string > list of regexp-like strings
func FindStringInRegexpList(compare string, listOfRegexp []string) bool {

	for _, regex := range listOfRegexp {

		if match, err := regexp.MatchString(regex, compare); err != nil {

			prtcl.Log.Println("error comparing regexp with string:", err)

			prtcl.PrintObject(regex, compare, err)

		} else {

			if match {

				return true
			}
		}
	}

	return false
}
