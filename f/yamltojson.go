package f

import "gopkg.in/yaml.v2"

// JSONToYAML converts json to yaml
//
// Parameters:
//   - `j` : []byte > unparsed json byte, converted from a string to a byte-array
func JSONToYAML(j []byte) ([]byte, error) {

	var jsonObj interface{}

	if err := yaml.Unmarshal(j, &jsonObj); err != nil {

		return nil, err
	}

	if yamlBytes, err := yaml.Marshal(jsonObj); err != nil {

		return nil, err

	} else {

		return yamlBytes, nil
	}
}
