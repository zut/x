package xx

import "github.com/gogf/gf/container/gset"

func IntersectSize(a, b []string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	s1 := gset.NewStrSetFrom(a)
	s2 := gset.NewStrSetFrom(b)
	return s1.Intersect(s2).Size()
}
