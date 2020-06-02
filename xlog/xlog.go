package xlog

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/debug/gdebug"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/labstack/gommon/color"
	"io"
	"time"
)

func (l *Logger) printToWriterColor(now time.Time, std io.Writer, buffer *bytes.Buffer) {
	s := gstr.Replace(buffer.String(), "\n", " ")
	s, _ = gregex.ReplaceString(` Stack: \d+\. .*$`, "", s)
	//s = s + strings.Repeat(" ", 10) + l.GetStackColor() + "\n"
	s = fmt.Sprintf("%-80s %s\n", s, l.GetStackColor())
	//fmt.Println(l.GetStackColor())
	if gregex.IsMatchString(` \[DEBU\] `, s) {
		s = color.White(s)
	} else if gregex.IsMatchString(` \[INFO\] `, s) {
		s = color.Green(s)
	} else if gregex.IsMatchString(` \[WARN\] `, s) {
		s = color.Yellow(s)
	} else if gregex.IsMatchString(` \[ERRO\] `, s) {
		s = color.Red(s)
	} else if gregex.IsMatchString(` \[PANI\] `, s) {
		s = color.Yellow(s, color.U, color.B)
	} else if gregex.IsMatchString(` \[FATA\] `, s) {
		s = color.Red(s, color.U, color.B)
	}
	if l.config.Writer == nil {
		// Output content to disk file.
		if l.config.Path != "" {
			l.printToFile(now, buffer)
		}
		// Allow output to stdout?
		if l.config.StdoutPrint {
			if _, err := std.Write(gconv.Bytes(s)); err != nil {
				panic(err)
			}
		}
	} else {
		if _, err := l.config.Writer.Write(buffer.Bytes()); err != nil {
			panic(err)
		}
	}
}

// GetStack returns the caller stack content,
// the optional parameter <skip> specify the skipped stack offset from the end point.
func (l *Logger) GetStackColor(skip ...int) string {
	stackSkip := l.config.StSkip
	if len(skip) > 0 {
		stackSkip += skip[0]
	}
	filters := []string{"/glog", "/xlog", "/go/pkg"}
	if l.config.StFilter != "" {
		filters = append(filters, l.config.StFilter)
	}
	s := gdebug.StackWithFilters(filters, stackSkip)
	s, _ = gregex.ReplaceString("[\r\t\n]", " ", s)
	//f, _ = gregex.ReplaceString(`^\d+\. +`, "", f)
	s, _ = gregex.ReplaceString(` {2,}`, " ", s)
	//fmt.Println(s)
	//fmt.Println(f)
	//f, _ = gregex.ReplaceString(`(\d+\. .*?[^/]+ /\S*/[^/]*:\d+) {5,}`, "$1 ", f)
	//f, _ = gregex.ReplaceString(`(\d+)\. .*?([^/]+) /\S*/([^/]*:\d+) `, "$1 $2 $3 ", f)
	ss, _ := gregex.MatchAllString(`(\d+)\. .*?([^/]+) /\S*/([^/]*:\d+) `, s)
	s2 := g.SliceStr{}
	for n, aa := range ss {
		if !gregex.IsMatchString(`github.com/zut/`, aa[0]) {
			continue
		}
		switch {
		case n == 0:
			s2 = append(s2, fmt.Sprintf("%s.%s %s ", aa[1], aa[2], aa[3]))
		case n <= 2:
			s2 = append(s2, fmt.Sprintf("%s.%s %s ", aa[1], aa[2], aa[3]))
		}
	}
	//f, _ = gregex.ReplaceString(` /\S*/([^/]*:\d+) `, " $1 ", f)
	//if err != nil {
	//	return "GetStackColor:" + err.Error()
	//}
	rst := "< " + gstr.JoinAny(s2, " ")
	//fmt.Println("rst:", rst)
	return rst
}
