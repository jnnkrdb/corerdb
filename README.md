# CoreRDB
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jnnkrdb/corerdb)](https://github.com/jnnkrdb/corerdb)
[![CodeFactor](https://www.codefactor.io/repository/github/jnnkrdb/corerdb/badge)](https://www.codefactor.io/repository/github/jnnkrdb/corerdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/jnnkrdb/corerdb)](https://goreportcard.com/report/github.com/jnnkrdb/corerdb)
[![Github tag](https://badgen.net/github/tag/jnnkrdb/corerdb)](https://github.com/jnnkrdb/corerdb/tags/)
[![GitHub issues](https://badgen.net/github/issues/jnnkrdb/corerdb/)](https://github.com/jnnkrdb/corerdb/issues/)
[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

---
## Description
corerdb is a go package, which contains core functions. It's a small collections of functions, which are used more often in projects and can be exported to a package, but are not as complex to be collected in a specific package, only containing structs, vars or something else, related to the specific function.
## Table of Contents

- [Install](#install)
- [Functionality](#functionality)
  - [Core-Functions](#core-functions)
  - [Encryption](#encryption)
  - [Protocol](#protocol)
  - [Directory](#directory)
- [FIY](#fiy)
---
## Install
Use the go commandline tool to install the package to your go project.
```
go get github.com/jnnkrdb/corerdb
```
---
## Functionality
### Core-Functions
The core function are combined in a package subfolder named "fnc". To import them in a project use:
```go
import "github.com/jnnkrdb/corerdb/fnc"
```
The fnc subpackage contains the following functionalities:
- ``JSONToYAML(j []byte) ([]byte, error)``: This function converts json ([]byte) into yaml ([]byte)
- ``UnencodeB64(b64 string) string``: Converts Base64-strings into unencoded strings.
- ``EncodeB64(text string) string``: Converts strings into Base64-encoded strings.
- ``LoadStructFromFile(datatype, file string, objpointer interface{}) error``: Loads a file into a struct. Currently ``json`` and ``yaml`` are supported. To load the file into the specific struct, the struct needs the specific tags. Only public fields can be loaded. The functions needs the object (of the struct) as a pointer (*struct).
- ``StringInList(_string string, _list []string) bool``: Checks the list, if the list contains `_string` the return value will be `true`.
### Encryption
The encryption subpackage contains only for functions for the encryption. To import them in a project use:
```go
import "github.com/jnnkrdb/corerdb/crypt"
```
The crypt subpackage contains the following functionalities:
- ``GetDefaultPassphrase() string``: Returns the string of the currently defined passphrase. Default is ``"DEFAULT-ENCRYPTION-PASSPHRASE"``.
- ``SetDefaultPassphrase(passphrase string)``: Sets the defined default passphrase to the given string. HIGHLY RECOMMENDED!
- ``EncryptWithDefault(text string) (string, error)``: Encrypts a given string with the default passphrase.
- ``Encrypt(passphrase, text string) (string, error)``: Encrypts a given string with the given passphrase.
- ``DecryptWithDefault(text string) (string, error)``: Decrypts a given string with the default passphrase.
- ``Decrypt(_passphrase, text string) (string, error)``: Decrypts a given string with the given passphrase.
### Protocol
The prtcl subpackage contains debugging and logging architecture. To import them in a project use:
```go
import "github.com/jnnkrdb/corerdb/prtcl"
```
The prtcl subpackage contains the following functionalities:
- ``Timestamp() string``: Returns a timestamp as a string in the format ``yyyy/mm/dd hh:mm:ss.nnnnnn``.
- ``ErrorKill(err error)``: If the given error is not `nil` the program will be terminated with status code `1`.
- ``SetDebugging(active bool)``: Actives the function ``PrintObject()`` and ``PrintJSON()`` with the output of ``prtcl.Log``.
- ``GetDebugging() bool``: Returns the current state of the debugging output.
- ``var Log *log.Logger``: With ``prtcel.Log`` you can access the default logging engine of the prtcl-package. The logging engine is based on the package `log`.
- ``PrintObject(object ...any)``: Prints all given objects to the output of `prtcl.Log`.
### Directory
The dir subpackage contains functionality for file-manageent and directory scanning. To import them in a project use:
```go
import "github.com/jnnkrdb/corerdb/dir"
```
The dir subpackage contains the following functionalities:
- ``GetFiles(dir string, excludes ...string) (files []File, e error)``: This function scans through the given directory and returns an array of type dir.File, a struct which is defined in this package. `excludes` can contain as many strings as wanted. These strings will be used as a pattern, so full pathes to specific files, which contain one of the excludes won't be saved in the dir.File-array.
- ``(f File) FullName() string``: Returns the absolute path and the name of the file defined in the object of dir.File.
- ``(f *File) Read() string``: Reads the content of a specific file, defined in the struct dir.File and returns it as a string. If the function gets called, the content of the file will be stored in the dir.File object.
The struct dir.File:
```go
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
```
---
## FIY
The whole package of `github.com/jnnkrdb/corerdb` uses the `github.com/jnnkrdb/corerdb/prtcl`-Logging architecture. To deactivate the debugging-prints use the `SetDebugging(active bool)` function and set `prtcl.Log = nil`, so the Log-Engine won't be used by this or other pacakges from `github.com/jnnkrdb`.
