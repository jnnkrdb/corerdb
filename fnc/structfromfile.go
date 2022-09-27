package fnc

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jnnkrdb/corerdb/prtcl"
	"gopkg.in/yaml.v2"
)

// load a struct from a file
//
// Parameters:
//   - `datatype` : string > declares the file type: ["json", "yaml"]
//   - `file` : string > path to the jsonfile, which contains the object, the json-tags are necessary
//   - `objpointer` : interface{} > pointer to the object, which should be loaded from json (*interface{})
func LoadStructFromFile(datatype, file string, objpointer interface{}) error {

	bytes, err := os.ReadFile(file)

	if err == nil {

		switch datatype {

		case "json":

			err = json.Unmarshal(bytes, objpointer)

		case "yaml":

			err = yaml.Unmarshal(bytes, objpointer)

		default:

			err = errors.New("datatype unsupported, please use 'json'or 'yaml'")
		}
	}

	if err != nil {

		prtcl.PrintObject(datatype, file, objpointer, string(bytes), err)
	}

	return err
}
