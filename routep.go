// The RouteP package is an overly simplistic way of comparing an URL with an URL template and
// extract variables out of it.
package routep

import (
	"fmt"
	"strings"
)

const (
	not_matching = "value not matching with template"
	value_longer = "value is longer than maximum allowed by template"
)

// It does allow shorter or longer URL than the template if the fix parts match.
//
//	"/person/{name}/hobby/{hobby}", "/person/john/hobby/cycling" 	=> map[name: john, hobby: cycling]
//
//	"/person/{name}/hobby/{hobby}", "/person/john"					=> map[name:john]
//
//	"/person/{name}",				"/person/john/hobby/cycling"	=> map[name:john]
func Comp(tpl, str string) (map[string]string, error) {
	t := strings.Split(tpl, "/")
	s := strings.Split(str, "/")
	ret := make(map[string]string)
	for i, v := range t {
		//if i == len(s)-1 {
		//	v = strings.Split(v, "?")[0]
		//}
		if v != "" && string(v[0]) == "{" && string(v[len(v)-1]) == "}" && i < len(s) {
			key := v[1 : len(v)-1]
			ret[key] = s[i]
		} else if i < len(s) && s[i] != v { // ha van
			return ret, fmt.Errorf(not_matching)
		}
	}
	return ret, nil
}

// Does not allow longer URL than the template.
//
//	"/person/{name}/hobby/{hobby}", "/person/john/hobby/cycling" 	=> map[name: john, hobby: cycling]
//
//	"/person/{name}/hobby/{hobby}", "/person/john"					=> map[name:john]
//
//	"/person/{name}",				"/person/john/hobby/cycling"	=> error
func CompStrict(tpl, str string) (map[string]string, error) {
	if strings.Count(tpl, "/") < strings.Count(str, "/") {
		return nil, fmt.Errorf(value_longer)
	}
	return Comp(tpl, str)
}
