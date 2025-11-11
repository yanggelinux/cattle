package util

import "sort"

//value string 类型的map进行排序
func SortStringMap(m map[string]string) map[string]string {

	sortedKeys := make([]string, 0)
	for k, _ := range m {
		sortedKeys = append(sortedKeys, k)
	}
	nm := make(map[string]string, len(sortedKeys))
	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		nm[key] = m[key]
	}
	return nm
}
