package xx

import (
	"fmt"
	"sort"
)

func SortMapByKey(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	//for _, k := range keys {
	//	fmt.Println(k, m[k])
	//}
	return keys
}

func SortMapByValue(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })
	return keys
}

// max to min
func SortMapByValueInt(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })
	return keys
}

func SortMapByValueIntMinToMax(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] < m[keys[j]] })
	return keys
}
