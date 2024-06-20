// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/genelet/determined/convert"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the HashiCups client is properly configured.
	// It is also possible to use the HASHICUPS_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = `
  provider "xmft" {
	product = "cft"
	alias = "cft1"
	username = "admin"
	password = "changeit"
	host     = "https://ci8.jda.axwaytest.net:1768"
  }
  provider "xmft" {
	product = "st"
	alias = "st1"
	username = "admin"
	password = "admin*"
	host     = "https://ci8.jda.axwaytest.net:8444"
	additional_attributes = {
	  Env = "test"
	}
  }
  `
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"xmft": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

type ResourceCheck struct {
	Type  string
	Name  string
	Key   string
	Value string
}

func check(typename, name, path string, m map[string]interface{}) []ResourceCheck {
	if path == "" {
		path = ""
	} else {
		path = path + "."
	}
	var l []ResourceCheck
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			l = append(l, check(typename, name, path+k, v)...)
		case bool, int, float64:
			l = append(l, ResourceCheck{typename, name, path + k, fmt.Sprint(v)})
			fmt.Println(typename, name, path+k, v)
		case string:
			l = append(l, ResourceCheck{typename, name, path + k, v})
			fmt.Println(typename, name, path+k, v)
		default:
			panic("unsupported stype :" + path + k + " " + fmt.Sprintf("%T", v))
		}
	}
	return l
}

func checkResources(bs []byte) resource.TestCheckFunc {
	jsonstr, err := convert.HCLToJSON(bs)
	if err != nil {
		panic(err)
	}
	// parse json to map
	var m map[string]interface{}
	err = json.Unmarshal(jsonstr, &m)
	if err != nil {
		panic(err)
	}

	var l []ResourceCheck
	resources := m["resource"].(map[string]interface{})
	for typename, v := range resources {
		m := v.(map[string]interface{})
		for name, attrs := range m {
			attrs := attrs.(map[string]interface{})
			l = append(l, check(typename, name, "", attrs)...)
		}
	}
	var funcs []resource.TestCheckFunc
	for _, r := range l {
		funcs = append(funcs, resource.TestCheckResourceAttr(r.Type+"."+r.Name, r.Key, r.Value))
	}
	return resource.ComposeAggregateTestCheckFunc(funcs...)
}

func init() {
	os.Setenv("TF_ACC", "1")
}
