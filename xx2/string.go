package xx2

import (
	"github.com/flytam/filenamify"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/zut/x/xlog"
	"unicode"
)

// SubStringHan 返回一个字符串的前 n 个字符，其中中文字符的长度为 2：
func SubStringHan(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	var length int
	for i := range runes {
		if length >= n {
			return string(runes[:i])
		}
		if unicode.Is(unicode.Scripts["Han"], runes[i]) {
			length += 2
		} else {
			length++
		}
	}
	return string(runes)
}

func SafeFilename(i string) string {
	i2, err := filenamify.Filenamify(i, filenamify.Options{
		Replacement: "!",
		MaxLength:   200, // linux win max 256
	})
	if err != nil {
		xlog.Warning(i, err)
		//i2, _ = gregex.ReplaceString(`\W`, "_", i)
		i2, _ = gregex.ReplaceString(`[<>:"/\\|?*\x00-\x1F]`, "_", i)
	}
	//后缀被干掉
	s, err := gregex.MatchString(`\.[a-zA-Z0-9]{1,10}$`, i)
	if err == nil && len(s) > 0 {
		suffix := s[0]
		if gstr.SubStr(i2, len(i2)-len(suffix)) != suffix {
			i2 += suffix
		}
	}
	return i2
}
