package toolbak

import (
	toolEncrypt "github.com/pangshu/tookbak/Encrypt"
	toolArray "github.com/pangshu/toolbak/Array"
	toolConvert "github.com/pangshu/toolbak/Convert"
	toolDate "github.com/pangshu/toolbak/Date"
	toolFile "github.com/pangshu/toolbak/File"
	toolHtml "github.com/pangshu/toolbak/Html"
	toolLog "github.com/pangshu/toolbak/Log"
	toolNet "github.com/pangshu/toolbak/Net"
	toolNumber "github.com/pangshu/toolbak/Number"
	toolOs "github.com/pangshu/toolbak/Os"
	toolRand "github.com/pangshu/toolbak/Rand"
	toolRegexp "github.com/pangshu/toolbak/Regexp"
	toolString "github.com/pangshu/toolbak/String"
	toolUrl "github.com/pangshu/toolbak/Url"
)

var (
	Url toolUrl.Url
	Os toolOs.Os
	Net toolNet.Net
	Html toolHtml.Html
	Log toolLog.Log
	File toolFile.File
	Date toolDate.Date
	Array toolArray.Array
	String  toolString.String
	Regexp  toolRegexp.Regexp
	Number  toolNumber.Number
	Convert toolConvert.Convert
	Rand toolRand.Rand
	Encrypt toolEncrypt.Encrypt
)
