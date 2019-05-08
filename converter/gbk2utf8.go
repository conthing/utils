package converter

import (
	"bytes"
    "io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
)

//Gb18030ToUtf8 GB18030 to UTF-8
func Gb18030ToUtf8(s string) (string, error) {
    reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GB18030.NewDecoder())
    d, err := ioutil.ReadAll(reader)
    if err != nil {
        return s, err
    }
    return string(d), nil
}
