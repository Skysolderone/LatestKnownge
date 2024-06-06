package utils

import (
	"sort"
	"strings"
)

func BuildGetParams(params map[string]string) string {
	//urlParams := url.Values{}
	//if params != nil && len(params) > 0 {
	//	for k := range params {
	//		urlParams.Add(k, params[k])
	//	}
	//}
	//return "?" + urlParams.Encode()
	if len(params) == 0 {
		return ""
	}
	return SortParams(params)
}

func SortParams(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		// sorted[i] = k + "=" + url.QueryEscape(params[k])
		sorted[i] = k + "=" + params[k]
		i++
	}
	return strings.Join(sorted, "&")
}
