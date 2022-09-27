package dir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jnnkrdb/corerdb/prtcl"
)

// collect all files from a given path and return an array of "Files"
//
// Parameters:
//   - `dir` : string > directory, in which the files will be searched
//   - `excludes` : ...string > a collection of strings, if a collected subpath contains one of the excludes, it will be skipped
func GetFiles(dir string, excludes ...string) (files []File, e error) {

	prtcl.Log.Println("read files from path", dir)

	var currDir string = ""

	e = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			prtcl.PrintObject(dir, excludes, files, currDir, path, info, err)

		} else {

			if !exclude(path, excludes) {

				if !info.IsDir() {

					files = append(files, File{
						Name:      info.Name(),
						Dir:       strings.ReplaceAll(path, "/"+info.Name(), ""),
						SearchDir: dir,
						Content:   "",
					})
				}
			}
		}

		return err
	})

	return
}

// exclude function, handles the excludes for the filewwalk. if "dir" contains
// an item from the "exclude" array, the returnvalue will be true
//
// Parameters:
//   - `dir` : string > directory, in which the files will be searched
//   - `exclude` : []string > a collection of strings which should be excluded
func exclude(dir string, exclude []string) bool {

	if len(exclude) > 0 {

		for _, exc := range exclude {

			if strings.Contains(dir, exc) {

				return true
			}
		}
	}

	return false
}
