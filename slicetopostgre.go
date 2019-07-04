package slicetopostgre

import (
	"fmt"
	"strconv"
	"strings"
)

// slice to a postgres array format
// support a slice of string, int and fmt.Stringer
func Array(value interface{}) string {
	var aux string
	var check = func(aux string, value interface{}) (ret string) {
		if aux != "" {
			aux += ","
		}
		ret = aux + Array(value)
		return
	}
	switch value.(type) {
	case []fmt.Stringer:
		for _, v := range value.([]fmt.Stringer) {
			aux = check(aux, v)
		}
		return "{" + aux + "}"
	case []interface{}:
		for _, v := range value.([]interface{}) {
			aux = check(aux, v)
		}
		return "{" + aux + "}"
	case []string:
		for _, v := range value.([]string) {
			aux = check(aux, v)
		}
		return "{" + aux + "}"
	case []int:
		for _, v := range value.([]int) {
			aux = check(aux, v)
		}
		return "{" + aux + "}"
	case string:
		aux := value.(string)
		aux = strings.Replace(aux, `\`, `\\`, -1)
		aux = strings.Replace(aux, `"`, `\"`, -1)
		return `"` + aux + `"`
	case int:
		return strconv.Itoa(value.(int))
	case fmt.Stringer:
		v := value.(fmt.Stringer)
		return Array(v.String())
	}
	return ""
}

