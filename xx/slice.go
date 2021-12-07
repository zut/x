package xx

import "github.com/gogf/gf/container/gset"

func IntersectSize(a, b []string) int {
	s1 := gset.NewStrSetFrom(a)
	s2 := gset.NewStrSetFrom(b)
	return s1.Intersect(s2).Size()
}
