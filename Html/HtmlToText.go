package Html

import (
	"golang.org/x/net/html"
	"strings"
	"github.com/pangshu/toolbak/String"
)

// HtmlToText 将html转换为纯文本.
func HtmlToText(str string) string {
	var textHtmlExcludeTags = []string{"head", "title", "img", "form", "textarea", "input", "select", "button", "iframe", "script", "style", "option"}
	if str == "" {
		return ""
	}

	domDoc := html.NewTokenizer(strings.NewReader(str))
	previousStartToken := domDoc.Token()
	var text string
loopDom:
	for {
		nx := domDoc.Next()
		switch {
		case nx == html.ErrorToken:
			break loopDom // End of the document
		case nx == html.StartTagToken:
			previousStartToken = domDoc.Token()
		case nx == html.TextToken:
			if chk, _ := String.Dstrpos(previousStartToken.Data, textHtmlExcludeTags, false); chk {
				continue
			}

			text += " " + strings.TrimSpace(html.UnescapeString(string(domDoc.Text())))
		}
	}

	return String.RemoveSpace(text, false)
}
