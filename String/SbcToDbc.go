package String

import "golang.org/x/text/width"

// SBC2DBC 全角转半角.
func (*String)SbcToDbc(s string) string {
	return width.Narrow.String(s)
}
