package String

import "golang.org/x/text/width"

// DBC2SBC 半角转全角.
func (*String)DbcToSbc(s string) string {
	return width.Widen.String(s)
}
