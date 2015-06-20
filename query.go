package query

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

var (
	defaultDelimiter = "+"
)

func SetDelimiter(delim string) {
	defaultDelimiter = delim
}

func Bool(req *http.Request, key string, def bool) bool {
	val := req.FormValue(key)
	if len(val) == 0 {
		return def
	}
	f, e := strconv.ParseBool(val)
	if e != nil {
		return def
	}
	return f
}

func Int(req *http.Request, key string, def int) int {
	val := req.FormValue(key)
	if len(val) == 0 {
		return def
	}
	i, e := strconv.Atoi(val)
	if e != nil {
		return def
	}
	return i
}

func Ints(req *http.Request, key string, def []int, delim ...string) []int {
	delim = append(delim, defaultDelimiter)
	val := req.FormValue(key)
	if len(val) == 0 {
		return def
	}
	res := []int{}
	for _, v := range regexp.MustCompile(fmt.Sprintf("[ %s]", delim[0])).Split(val, -1) {
		if i, e := strconv.Atoi(v); e == nil {
			res = append(res, i)
		}
	}
	return res
}

func String(req *http.Request, key string, def string) string {
	val := req.FormValue(key)
	if len(val) == 0 {
		return def
	}
	return val
}

func Strings(req *http.Request, key string, def []string, delim ...string) []string {
	delim = append(delim, defaultDelimiter)
	val := req.FormValue(key)
	if len(val) == 0 {
		return def
	}
	return regexp.MustCompile(fmt.Sprintf("[ %s]", delim[0])).Split(val, -1)
}
