package fnc

// compares the given string with the given list, if the list contains the
// string, the returnvalue will be true
//
// Parameters:
//   - `_string` : string > contains the string, which will be looked for in the list
//   - `_list` : []string > list of string, maybe contains the given string
func StringInList(_string string, _list []string) bool {
	for index := range _list {
		if _list[index] != "" {
			if _list[index] == _string {
				return true
			}
		}
	}
	return false
}
