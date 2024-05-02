package tfhelper

import (
	"strings"
)

func ResolveURI(uri string, map1 map[string]interface{}) string {
	for k, v := range map1 {
		switch v := v.(type) {
		case string:
			uri = strings.Replace(uri, "{"+k+"}", v, -1)
		}
	}
	return uri
}
