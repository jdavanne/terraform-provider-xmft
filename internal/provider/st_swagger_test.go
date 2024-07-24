package provider

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"terraform-provider-xmft/internal/tfhelper"

	"gopkg.in/yaml.v2"
)

func Keys[T comparable](m map[T]interface{}) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func get(t *testing.T, client *http.Client, refs map[string]interface{}, currentFile string, path string, obj interface{}) interface{} {
	parts := strings.Split(path, ".")
	res, _ := getParts(t, client, refs, currentFile, false, parts, obj)
	return res
}

func getNoFollow(t *testing.T, client *http.Client, refs map[string]interface{}, currentFile string, path string, obj interface{}) interface{} {
	parts := strings.Split(path, ".")
	res, _ := getParts(t, client, refs, currentFile, false, parts, obj)
	return res
}

func getParts(t *testing.T,
	client *http.Client,
	refs map[string]interface{},
	currentFile string,
	followRef bool,
	path []string, obj interface{},
) (interface{}, string) {
	for _, p := range path {
		for {
			switch v := obj.(type) {
			case map[string]interface{}:
				if followRef && v["$ref"] != nil {
					// fmt.Println(path, "$ref", fmt.Sprint(v["$ref"]))
					obj, currentFile = sref(t, client, refs, currentFile, fmt.Sprint(v["$ref"]))
					continue
				}
				// fmt.Println("  ", p, Keys(v))
				obj = v[p]
			case map[interface{}]interface{}:
				if followRef && v["$ref"] != nil {
					// fmt.Println(path, "$ref", fmt.Sprint(v["$ref"]))
					obj, currentFile = sref(t, client, refs, currentFile, fmt.Sprint(v["$ref"]))
					continue
				}
				obj = nil
				// fmt.Println("  ", p, Keys(v))
				for k, v := range v {
					k2 := fmt.Sprint(k)
					if k2 == p {
						obj = v
					}
				}
			case nil:
				// fmt.Println("  ", path, obj, "-->nil")
				return nil, currentFile
			default:
				panic("unsupported type:" + fmt.Sprintf("%T", v))
			}
			break
		}
	}
	return obj, currentFile
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func fetch(client *http.Client, uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://ci8.jda.axwaytest.net:8444"+uri, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin*"))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("fetch " + uri + " status code:" + fmt.Sprint(resp.StatusCode) + " " + string(body))
	}

	return body, nil
}

func update(client *http.Client, uri string, path string) ([]byte, error) {
	doUpdate := false

	var data []byte
	var err error
	if doUpdate {
		data, err = fetch(client, uri)
		if err != nil {
			return nil, err
		}
	}
	old, err := os.ReadFile(path)

	if !doUpdate {
		return old, err
	} else {
		if err == nil {
			if !doUpdate || string(old) == string(data) {
				return old, nil
			}
		}
	}

	if doUpdate {
		err = os.WriteFile(path, data, 0o644)
		if err != nil {
			panic(err)
		}
	}
	return data, nil
}

func loadYml(t *testing.T, client *http.Client, k string) map[string]interface{} {
	swagger, err := update(client, "/api/v2.0/docs/"+k, "./testdata/"+k)
	if err != nil {
		t.Errorf("Error fetching %s : %v", k, err)
		t.FailNow()
	}

	var obj map[string]interface{}
	err = yaml.Unmarshal(swagger, &obj)
	if err != nil {
		t.Errorf("Error unmarshalling YAML %s : %v", k, err)
		t.FailNow()
	}
	return obj
}

func sref(t *testing.T, client *http.Client, refs map[string]interface{}, currentFile, path string) (interface{}, string) {
	parts := strings.Split(path, "#")

	filename := currentFile
	if parts[0] != "" {
		filename = parts[0]
	}

	if refs[filename] == nil {
		refs[filename] = loadYml(t, client, filename)
	}
	r := refs[filename]
	pathParts := strings.Split(parts[1], "/")
	res, _ := getParts(t, client, refs, currentFile, false, pathParts[1:], r)
	return res, filename
}

func describeSchema(t *testing.T, client *http.Client, refs map[string]interface{}, currentFile, path string, schema interface{}) []map[string]string {
	fields := make([]map[string]string, 0)
	debug := false
	switch v := schema.(type) {
	case map[interface{}]interface{}:
		typ := v["type"]
		switch typ {
		case "array":
			// fmt.Println(path, "array")
			items := v["items"]
			fields = append(fields, describeSchema(t, client, refs, currentFile, path+".#", items)...)
		case "string", "integer", "number", "boolean":
			// fmt.Println(path, typ, v)
			fields = append(fields, map[string]string{
				"apiPath":     path,
				"type":        fmt.Sprint(typ),
				"default":     fmt.Sprint(v["default"]),
				"description": fmt.Sprint(v["description"]),
				"enum":        fmt.Sprint(v["enum"]),
				"format":      fmt.Sprint(v["format"]),
				"maxLength":   fmt.Sprint(v["maxLength"]),
				"minLength":   fmt.Sprint(v["minLength"]),
			})
		default:
			found := false

			if !found {
				if v["type"] == "object" || v["properties"] != nil {
					found = true
					properties := v["properties"]
					if properties != nil {
						for k, v2 := range properties.(map[interface{}]interface{}) {
							fields = append(fields, describeSchema(t, client, refs, currentFile, path+"."+fmt.Sprint(k), v2)...)
						}
					} else {
						// fmt.Println(path, "map")
						fields = append(fields, map[string]string{
							"apiPath":     path,
							"type":        "map",
							"description": fmt.Sprint(v["description"]),
						})
					}
				}
			}

			if !found {
				ref := v["$ref"]
				if ref != nil {
					found = true
					f, nc := sref(t, client, refs, currentFile, fmt.Sprint(ref))
					detail := ""
					if debug {
						detail = fmt.Sprint(".$ref(", ref, ")")
					}

					fields = append(fields, describeSchema(t, client, refs, nc, path+detail, f)...)
					// fmt.Println(path, "$ref", ref)
				}
			}

			if !found {
				allOf := v["allOf"]
				if allOf != nil {
					found = true
					for i, v2 := range allOf.([]interface{}) {
						detail := ""
						if debug {
							detail = fmt.Sprint(".allOf(", i, ")")
						}
						fields = append(fields, describeSchema(t, client, refs, currentFile, path+detail, v2)...)
					}
				}
			}

			if !found {
				discriminator := v["discriminator"]

				if discriminator != nil {
					found = true

					propertyName := fmt.Sprint(v["discriminator"].(map[interface{}]interface{})["propertyName"])
					mapping := v["discriminator"].(map[interface{}]interface{})["mapping"]

					for k, ref := range mapping.(map[interface{}]interface{}) {
						// val := fmt.Sprint(get(t, client, refs, currentFile, "type.properties.default", v2))
						f, nc := sref(t, client, refs, currentFile, fmt.Sprint(ref))
						fields = append(fields, describeSchema(t, client, refs, nc, path+".["+propertyName+"="+fmt.Sprint(k)+"]", f)...)
					}

				}
			}

			if !found {
				anyOf := v["anyOf"] // for schedules only ????
				if anyOf != nil {
					found = true
					for i, v2 := range anyOf.([]interface{}) {
						detail := ""
						if debug {
							detail = fmt.Sprint(".anyOf(", i, ")")
						}
						// typ, _ := getParts(t, client, refs, currentFile, true, []string{"properties", "type", "type"}, v2)
						def, _ := getParts(t, client, refs, currentFile, true, []string{"properties", "type", "default"}, v2)
						detail = fmt.Sprint(".anyOf(", i, ")")
						if def != nil {
							detail = fmt.Sprint(".[type=", def, "]")
						}
						fields = append(fields, describeSchema(t, client, refs, currentFile, path+detail, v2)...)
					}
				}
			}

			if !found {
				// FIXME: required
				if v["required"] != nil || v["default"] != nil || v["description"] != nil {
					found = true
					fmt.Println(path, "oups")
				}
			}
			if !found {
				panic("unsupported type:" + fmt.Sprintf("%s : %T %v", path, v, v))
			}
		}
	default:
		panic("unsupported type:" + fmt.Sprintf("%s %T %v", path, v, v))
		// fmt.Println(path, v)
	}
	return fields
}

func walkSTSwaggerResources(t *testing.T) []map[string]string {
	// open and parse yaml file named ./st_swagger.yaml
	// read the data from the file
	// convert the data to a struct
	// print the struct
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	refs := make(map[string]interface{})

	currentFile := "swagger.yaml"
	stSwagger := loadYml(t, client, currentFile)
	fields := make([]map[string]string, 0)

	discarded := make([]string, 0)
main:
	for k, p := range stSwagger["paths"].(map[interface{}]interface{}) {
		uri := fmt.Sprint(k)

		if strings.HasSuffix(uri, "/operations") {
			discarded = append(discarded, uri+"(operations)")
			continue
		}

		for _, skip := range []string{
			"/accountSetup",
			"/sessions",
			"/events",
			"/logs/transfers",
			"/logs/audit",
			"/logs/server",
			"/servers",
			"/certificates/requests",
			"/configurations/clusterManagement/",
		} {
			if strings.HasPrefix(uri, skip) {
				discarded = append(discarded, uri+"(skip_prefix)")
				continue main
			}
		}

		hasNoSkip := false
		for _, noskip := range []string{"/configurations/sentinel", "/configurations/fileArchiving"} {
			if strings.HasPrefix(uri, noskip) {
				hasNoSkip = true
			}
		}

		if !hasNoSkip && !strings.HasSuffix(uri, "}") {
			discarded = append(discarded, uri+"(no end curly)")
			continue main
		}

		// fmt.Println(uri)
		obj := get(t, client, refs, currentFile, "get.responses.200.content.application/json.schema", p)
		if obj == nil {
			continue
		}

		found := false
		ref := getNoFollow(t, client, refs, currentFile, "$ref", obj)
		// ref := obj["$ref"]
		if ref != nil {
			// fmt.Println(uri, "$ref", ref)
			found = true

			f, nc := sref(t, client, refs, currentFile, fmt.Sprint(ref))
			fields = append(fields, describeSchema(t, client, refs, nc, uri, f)...)
			// fmt.Println(uri, f)

		}

		if !found {
			propertyName := fmt.Sprint(get(t, client, refs, currentFile, "discriminator.propertyName", obj))
			refa := get(t, client, refs, currentFile, "discriminator.mapping", obj)
			if refa != nil {
				// fmt.Println(uri, "discriminator.mapping", refa)
				found = true

				refa2 := refa.(map[interface{}]interface{})
				for ki, ref := range refa2 {
					k := fmt.Sprint(ki)
					fmt.Println(uri, "discriminator.mapping", k, ref)

					f, nc := sref(t, client, refs, currentFile, fmt.Sprint(ref))
					fields = append(fields, describeSchema(t, client, refs, nc, uri+"["+propertyName+"="+k+"]", f)...)
				}
			}
		}

		if !found {
			refa := get(t, client, refs, currentFile, "items", obj)
			if refa != nil {
				continue // skip array
			}
		}

		if !found {
			panic("no ref found for:" + uri + " " + fmt.Sprint(obj))
		}

	}
	fmt.Println("n", len(fields))

	for _, f := range fields {
		fmt.Println(f["apiPath"], f["type"], f["default"])
	}

	fmt.Println("discarded", len(discarded))
	sort.Slice(discarded, func(i, j int) bool {
		return discarded[i] < discarded[j]
	})
	for _, d := range discarded {
		fmt.Println("discarded", d)
	}

	fmt.Println("n", len(fields))

	return fields
}

func walkSTProviderResources(t *testing.T) []map[string]string {
	fields := make([]map[string]string, 0)
	for _, res := range localProviderResources {
		stRes, ok := res().(*stResource)
		if !ok {
			continue
		}

		// fmt.Println(stRes.name, stRes.kind, stRes.uriReplace)
		uri := stRes.uriReplace[9:] + stRes.swaggerDiscriminator
		fields = append(fields, tfhelper.ModelFlatten(nil, stRes.name, stRes.name, uri, false, reflect.TypeOf(reflect.ValueOf(stRes.obj).Elem().Interface()))...)
	}
	fmt.Println("n", len(fields))

	for _, f := range fields {
		fmt.Println(f["apiPath"], f["type"], f["default"])
	}

	fmt.Println("n", len(fields))

	return fields
}

func TestResources(t *testing.T) {
	providerResources := walkSTProviderResources(t)
	swaggerResources := walkSTSwaggerResources(t)

	sort.Slice(providerResources, func(i, j int) bool {
		return providerResources[i]["apiPath"] < providerResources[j]["apiPath"]
	})

	sort.Slice(swaggerResources, func(i, j int) bool {
		return swaggerResources[i]["apiPath"] < swaggerResources[j]["apiPath"]
	})

	fmt.Println("provider n", len(providerResources))
	fmt.Println("swagger n", len(swaggerResources))

	providerResourcesUnique := make(map[string]map[string]string)
	for _, f := range providerResources {
		providerResourcesUnique[f["apiPath"]] = f
	}

	swaggerResourcesUnique := make(map[string]map[string]string)
	for _, f := range swaggerResources {
		swaggerResourcesUnique[f["apiPath"]] = f
	}

	totalprovider := 0
	okProvider := 0
loop1:
	for _, f := range providerResources {
		for _, skip := range []string{".schedules.", ".transferConfigurations."} {
			if strings.Contains(f["apiPath"], skip) {
				continue loop1
			}
		}
		totalprovider++
		if _, ok := swaggerResourcesUnique[f["apiPath"]]; !ok {
			fmt.Println("missing provider", f["apiPath"], f["type"], f["default"])
		} else {
			fmt.Println("ok      provider", f["apiPath"], f["type"], f["default"])
			okProvider++
		}
	}

	fieldDescriptions := make(map[string]string)
	totalswagger := 0
	okSwagger := 0
loop2:
	for _, f := range swaggerResources {
		for _, skip := range []string{".schedules.", ".transferConfigurations."} {
			if strings.Contains(f["apiPath"], skip) {
				continue loop2
			}
		}
		totalswagger++
		if _, ok := providerResourcesUnique[f["apiPath"]]; !ok {
			fmt.Println("missing swagger", f["apiPath"], f["type"], f["default"])
		} else {
			fmt.Println("ok      swagger", f["apiPath"], f["type"], f["default"], f["description"])
			fieldDescriptions[f["apiPath"]] = f["description"]
			okSwagger++
		}
	}

	fmt.Println("okProvider", okProvider, totalprovider, len(providerResources))
	fmt.Println("okSwagger", okSwagger, totalswagger, len(swaggerResources))

	fieldDescriptionsJson, err := json.MarshalIndent(fieldDescriptions, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./provider_st_field_descriptions.json", fieldDescriptionsJson, 0o644)
	if err != nil {
		panic(err)
	}
}
