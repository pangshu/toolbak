package String

import (
	"bytes"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// Big5ToUtf8 BIG5转UTF-8编码.
func (*String)Big5ToUtf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	return string(d), e
}
