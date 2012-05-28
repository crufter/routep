// A routep package egy nagyon egyszerű (ezáltal gyors) módja annak, hogy egy url-t összehasunlítsunk egy "url template"-tel
// és kiextraháljuk belőle a változókat.
//
// Működése: 
package routep

import(
	"strings"
)

const(
	not_matching	= "value not matching with template"
	value_longer	= "value is longer than maximum allowed by template"
)

// Rövidebb és hosszabb urlt is megenged, mint a template ha a fix részek megegyeznek.
//
//	"/person/{name}/hobby/{hobby}", "/person/john/hobby/cycling" 	=> map[name: john, hobby: cycling]
//
//	"/person/{name}/hobby/{hobby}", "/person/john"					=> map[name:john]
//
//	"/person/{name}",				"/person/john/hobby/cycling"	=> map[name:john]
func Comp(tpl, str string) (map[string]string, string) {
	t := strings.Split(tpl, "/")
	s := strings.Split(str, "/")
	ret := make(map[string]string)
	for i, v := range t {
		//if i == len(s)-1 {
		//	v = strings.Split(v, "?")[0]
		//}
		if v != "" && string(v[0]) == "{" && string(v[len(v)-1]) == "}" && i < len(s) {
			key := v[1:len(v)-1]
			ret[key] = s[i]
		} else if i < len(s) && s[i] != v {	// ha van
			return ret, not_matching
		}
	}
	return ret, ""
}

// Hosszabb urlt nem enged meg, mint a template.
//
//	"/person/{name}/hobby/{hobby}", "/person/john/hobby/cycling" 	=> map[name: john, hobby: cycling]
//
//	"/person/{name}/hobby/{hobby}", "/person/john"					=> map[name:john]
//
//	"/person/{name}",				"/person/john/hobby/cycling"	=> error
func CompStrict(tpl, str string) (map[string]string, string) {
	if strings.Count(tpl, "/") < strings.Count(str, "/") {
		return nil, value_longer
	}
	return Comp(tpl, str)
}