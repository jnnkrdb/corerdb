package dir

import (
	"os"

	"github.com/jnnkrdb/corerdb/prtcl"
)

// Struct to group information about a file in the directory.
// The struct can be exported as a json-object.
//
//	{
//		"name":"<filename>",
//		"directory":"<path/to/file>",
//		"searcheddirectory":"path/which/got/scanned",
//		"content":"<filecontent, if loaded>"
//	}
type File struct {
	// name of the actual file, without the path
	Name string `json:"name"`
	// full path, without the filename
	Dir string `json:"directory"`
	// path, given to the parameter [dir string] from files.GetFiles(dir string)
	SearchDir string `json:"searchdirectory"`
	// content which gets persisted, after Read() gets called
	Content string `json:"content"`
}

// get the full name of the file
func (f File) FullName() string {

	return f.Dir + "/" + f.Name
}

// get the files content
func (f *File) Read() string {

	f.Content = ""

	if filebyte, err := os.ReadFile(f.FullName()); err != nil {

		prtcl.PrintObject(f, filebyte, err)

	} else {

		f.Content = string(filebyte)
	}

	return f.Content
}
