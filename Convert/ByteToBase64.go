package Convert

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// byte转base64
func (*Convert)ByteToBase64(content []byte, ext ...string) string {
	if len(ext) > 0 {
		extType := strings.ToLower(ext[0])
		return fmt.Sprintf("data:image/%s;base64,%s", extType, base64.StdEncoding.EncodeToString(content))
	} else {
		return fmt.Sprintf("%s", base64.StdEncoding.EncodeToString(content))
	}
}
