package fnc

import "encoding/base64"

// this function just skips the error, which the base64-package throws
//
// Parameters:
//   - `b64` : string > base64 encoded string, which should be unencoded
func UnencodeB64(b64 string) string {

	if str, err := base64.StdEncoding.DecodeString(b64); err != nil {

		return ""

	} else {

		return string(str)
	}
}

// encode something to base64, this function only exists, because there is a
// function to unencode base64 strings... it is not necessary, you can just
// use base64.StdEncoding.EncodeToString([]byte(text))
//
// Parameters:
//   - `text` : string > unencoded string, which should be abse64 encoded
func EncodeB64(text string) string {

	return base64.StdEncoding.EncodeToString([]byte(text))
}
