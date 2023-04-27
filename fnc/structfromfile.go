package fnc

import (
	"encoding/json"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// load a struct from a file
//
// Parameters:
//   - `datatype` : string > declares the file type: ["json", "yaml"]
//   - `file` : string > path to the jsonfile, which contains the object, the json-tags are necessary
//   - `objpointer` : interface{} > pointer to the object, which should be loaded from json (*interface{})
func LoadStructFromFile(datatype, file string, objpointer interface{}) (err error) {
	var b []byte
	if b, err = os.ReadFile(file); err == nil {
		switch datatype {
		case "json":
			err = json.Unmarshal(b, objpointer)
		case "yaml":
			err = yaml.Unmarshal(b, objpointer)
		default:
			err = errors.New("datatype unsupported, please use 'json'or 'yaml'")
		}
	}
	return err
}
