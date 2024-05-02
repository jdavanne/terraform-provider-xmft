package tfhelper

import (
	"errors"
	"slices"
	"strings"
)

const (
	tfsdkTagName  = "tfsdk"
	helperTagName = "helper"
)

// ModelToSchema converts the fields in the Terraform resource to the corresponding attributes in the XCO MFT object.
// It takes the resource name and the resource object as input.
// The function returns a schema.Schema object that contains the attributes of the resource.
// helper attributes :
//   required
//   computed
//   default
//   modified
//   sensitive
//   state
//   nope

var supportedAttributes = []string{"required", "computed", "sensitive", "state", "optional", "noread", "nowrite", "fold", "readmap", "default", "default:", "elementtype:"}

func checkSupportedAttribute(flag string) bool {
	for _, v := range supportedAttributes {
		if v == flag {
			return true
		}
		if v[len(v)-1] == ':' && strings.HasPrefix(flag, v) {
			return true
		}
	}
	return false
}

func checkSupportedAttributes(location, flags string) error {
	f := strings.Split(flags, ",")
	for i, v := range f {
		if i == 0 {
			if v != "" && checkSupportedAttribute(v) {
				return errors.New("conflicting name with helper flag" + location + " : '" + v + "'")
			}
		} else if v != "" && !checkSupportedAttribute(v) {
			return errors.New("unsupported helper flag " + location + " : " + v)
		}
	}
	return nil
}

func mustCheckSupportedAttributes(location, flags string) {
	if err := checkSupportedAttributes(location, flags); err != nil {
		panic(err.Error())
	}
}

func FlagsHas(flags, flag string) bool {
	f := strings.Split(flags, ",")
	return slices.Contains(f, flag)
}

func FlagsGet(flags, flag string) (string, bool) {
	f := strings.Split(flags, ",")
	for _, v := range f {
		if strings.HasPrefix(v, flag) {
			if strings.HasPrefix(v, flag+":") {
				val := strings.Split(v, ":")[1]
				return val, true
			} else {
				val := ""
				return val, true
			}
		}
	}
	return "", false
}

func FlagsTfsdkGetName(tfsdkFlags string) string {
	f := strings.Split(tfsdkFlags, ",")
	return f[0]
}

func FlagsHelperName(tfsdkFlags, helperFlags string) string {
	f := strings.Split(helperFlags, ",")
	if f[0] != "" {
		return f[0]
	}
	return FlagsTfsdkGetName(tfsdkFlags)
}
