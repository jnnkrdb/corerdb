package fnc

// compares the given string with the given list, if the list contains the
// string, the returnvalue will be true
//
// Parameters:
//   - `_string` : string > contains the string, which will be looked for in the list
//   - `_list` : []string > list of string, maybe contains the given string
func StringInList(_string string, _list []string) bool {

	for _, str := range _list {

		if str != "" {

			if str == _string {

				return true
			}
		}
	}

	return false
}
