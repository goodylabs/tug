package utils

import (
	"sort"
	"strings"
)

func SortOptions(options []string) {
	sort.Strings(options)
}

func FilterExclude(items []string, exclude string) []string {
	var result []string
	for _, item := range items {
		if !strings.Contains(item, exclude) {
			result = append(result, item)
		}
	}
	return result
}

func FilterInclude(items []string, exclude string) []string {
	var result []string
	for _, item := range items {
		if strings.Contains(item, exclude) {
			result = append(result, item)
		}
	}
	return result
}
