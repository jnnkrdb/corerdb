package fnc

import "gopkg.in/yaml.v2"

// JSONToYAML converts json to yaml
//
// Parameters:
//   - `j` : []byte > unparsed json byte, converted from a string to a byte-array
func JSONToYAML(j []byte) (yamlBytes []byte, err error) {
	var jsonObj interface{}
	if err = yaml.Unmarshal(j, &jsonObj); err == nil {
		return yaml.Marshal(jsonObj)
	}
	return
}
