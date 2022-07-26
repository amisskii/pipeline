/*
Copyright 2022 The Tekton Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/test/diff"
)

func TestParamSpec_SetDefaults(t *testing.T) {
	tests := []struct {
		name            string
		before          *v1.ParamSpec
		defaultsApplied *v1.ParamSpec
	}{{
		name: "inferred string type",
		before: &v1.ParamSpec{
			Name: "parametername",
		},
		defaultsApplied: &v1.ParamSpec{
			Name: "parametername",
			Type: v1.ParamTypeString,
		},
	}, {
		name: "inferred type from default value - array",
		before: &v1.ParamSpec{
			Name: "parametername",
			Default: &v1.ArrayOrString{
				ArrayVal: []string{"array"},
			},
		},
		defaultsApplied: &v1.ParamSpec{
			Name: "parametername",
			Type: v1.ParamTypeArray,
			Default: &v1.ArrayOrString{
				ArrayVal: []string{"array"},
			},
		},
	}, {
		name: "inferred type from default value - string",
		before: &v1.ParamSpec{
			Name: "parametername",
			Default: &v1.ArrayOrString{
				StringVal: "an",
			},
		},
		defaultsApplied: &v1.ParamSpec{
			Name: "parametername",
			Type: v1.ParamTypeString,
			Default: &v1.ArrayOrString{
				StringVal: "an",
			},
		},
	}, {
		name: "inferred type from default value - object",
		before: &v1.ParamSpec{
			Name: "parametername",
			Default: &v1.ArrayOrString{
				ObjectVal: map[string]string{"url": "test", "path": "test"},
			},
		},
		defaultsApplied: &v1.ParamSpec{
			Name: "parametername",
			Type: v1.ParamTypeObject,
			Default: &v1.ArrayOrString{
				ObjectVal: map[string]string{"url": "test", "path": "test"},
			},
		},
	}, {
		name: "inferred type from properties - PropertySpec type is not provided",
		before: &v1.ParamSpec{
			Name:       "parametername",
			Properties: map[string]v1.PropertySpec{"key1": {}},
		},
		defaultsApplied: &v1.ParamSpec{
			Name:       "parametername",
			Type:       v1.ParamTypeObject,
			Properties: map[string]v1.PropertySpec{"key1": {Type: "string"}},
		},
	}, {
		name: "inferred type from properties - PropertySpec type is provided",
		before: &v1.ParamSpec{
			Name:       "parametername",
			Properties: map[string]v1.PropertySpec{"key2": {Type: "string"}},
		},
		defaultsApplied: &v1.ParamSpec{
			Name:       "parametername",
			Type:       v1.ParamTypeObject,
			Properties: map[string]v1.PropertySpec{"key2": {Type: "string"}},
		},
	}, {
		name: "fully defined ParamSpec - array",
		before: &v1.ParamSpec{
			Name:        "parametername",
			Type:        v1.ParamTypeArray,
			Description: "a description",
			Default: &v1.ArrayOrString{
				ArrayVal: []string{"array"},
			},
		},
		defaultsApplied: &v1.ParamSpec{
			Name:        "parametername",
			Type:        v1.ParamTypeArray,
			Description: "a description",
			Default: &v1.ArrayOrString{
				ArrayVal: []string{"array"},
			},
		},
	}, {
		name: "fully defined ParamSpec - object",
		before: &v1.ParamSpec{
			Name:        "parametername",
			Type:        v1.ParamTypeObject,
			Description: "a description",
			Default: &v1.ArrayOrString{
				ObjectVal: map[string]string{"url": "test", "path": "test"},
			},
		},
		defaultsApplied: &v1.ParamSpec{
			Name:        "parametername",
			Type:        v1.ParamTypeObject,
			Description: "a description",
			Default: &v1.ArrayOrString{
				ObjectVal: map[string]string{"url": "test", "path": "test"},
			},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			tc.before.SetDefaults(ctx)
			if d := cmp.Diff(tc.before, tc.defaultsApplied); d != "" {
				t.Error(diff.PrintWantGot(d))
			}
		})
	}
}

func TestArrayOrString_ApplyReplacements(t *testing.T) {
	type args struct {
		input              *v1.ArrayOrString
		stringReplacements map[string]string
		arrayReplacements  map[string][]string
		objectReplacements map[string]map[string]string
	}
	tests := []struct {
		name           string
		args           args
		expectedOutput *v1.ArrayOrString
	}{{
		name: "no replacements on array",
		args: args{
			input:              v1.NewArrayOrString("an", "array"),
			stringReplacements: map[string]string{"some": "value", "anotherkey": "value"},
			arrayReplacements:  map[string][]string{"arraykey": {"array", "value"}, "sdfdf": {"sdf", "sdfsd"}},
		},
		expectedOutput: v1.NewArrayOrString("an", "array"),
	}, {
		name: "single string replacement on string",
		args: args{
			input:              v1.NewArrayOrString("$(params.myString1)"),
			stringReplacements: map[string]string{"params.myString1": "value1", "params.myString2": "value2"},
			arrayReplacements:  map[string][]string{"arraykey": {"array", "value"}, "sdfdf": {"asdf", "sdfsd"}},
		},
		expectedOutput: v1.NewArrayOrString("value1"),
	}, {
		name: "multiple string replacements on string",
		args: args{
			input:              v1.NewArrayOrString("astring$(some) asdf $(anotherkey)"),
			stringReplacements: map[string]string{"some": "value", "anotherkey": "value"},
			arrayReplacements:  map[string][]string{"arraykey": {"array", "value"}, "sdfdf": {"asdf", "sdfsd"}},
		},
		expectedOutput: v1.NewArrayOrString("astringvalue asdf value"),
	}, {
		name: "single array replacement",
		args: args{
			input:              v1.NewArrayOrString("firstvalue", "$(arraykey)", "lastvalue"),
			stringReplacements: map[string]string{"some": "value", "anotherkey": "value"},
			arrayReplacements:  map[string][]string{"arraykey": {"array", "value"}, "sdfdf": {"asdf", "sdfsd"}},
		},
		expectedOutput: v1.NewArrayOrString("firstvalue", "array", "value", "lastvalue"),
	}, {
		name: "multiple array replacement",
		args: args{
			input:              v1.NewArrayOrString("firstvalue", "$(arraykey)", "lastvalue", "$(sdfdf)"),
			stringReplacements: map[string]string{"some": "value", "anotherkey": "value"},
			arrayReplacements:  map[string][]string{"arraykey": {"array", "value"}, "sdfdf": {"asdf", "sdfsd"}},
		},
		expectedOutput: v1.NewArrayOrString("firstvalue", "array", "value", "lastvalue", "asdf", "sdfsd"),
	}, {
		name: "empty array replacement without extra elements",
		args: args{
			input:             v1.NewArrayOrString("$(arraykey)"),
			arrayReplacements: map[string][]string{"arraykey": {}},
		},
		expectedOutput: &v1.ArrayOrString{Type: v1.ParamTypeArray, ArrayVal: []string{}},
	}, {
		name: "empty array replacement with extra elements",
		args: args{
			input:              v1.NewArrayOrString("firstvalue", "$(arraykey)", "lastvalue"),
			stringReplacements: map[string]string{"some": "value", "anotherkey": "value"},
			arrayReplacements:  map[string][]string{"arraykey": {}},
		},
		expectedOutput: v1.NewArrayOrString("firstvalue", "lastvalue"),
	}, {
		name: "array replacement on string val",
		args: args{
			input:             v1.NewArrayOrString("$(params.myarray)"),
			arrayReplacements: map[string][]string{"params.myarray": {"a", "b", "c"}},
		},
		expectedOutput: v1.NewArrayOrString("a", "b", "c"),
	}, {
		name: "array star replacement on string val",
		args: args{
			input:             v1.NewArrayOrString("$(params.myarray[*])"),
			arrayReplacements: map[string][]string{"params.myarray": {"a", "b", "c"}},
		},
		expectedOutput: v1.NewArrayOrString("a", "b", "c"),
	}, {
		name: "array indexing replacement on string val",
		args: args{
			input:              v1.NewArrayOrString("$(params.myarray[0])"),
			stringReplacements: map[string]string{"params.myarray[0]": "a", "params.myarray[1]": "b"},
		},
		expectedOutput: v1.NewArrayOrString("a"),
	}, {
		name: "object replacement on string val",
		args: args{
			input: v1.NewArrayOrString("$(params.object)"),
			objectReplacements: map[string]map[string]string{
				"params.object": {
					"url":    "abc.com",
					"commit": "af234",
				},
			},
		},
		expectedOutput: v1.NewObject(map[string]string{
			"url":    "abc.com",
			"commit": "af234",
		}),
	}, {
		name: "object star replacement on string val",
		args: args{
			input: v1.NewArrayOrString("$(params.object[*])"),
			objectReplacements: map[string]map[string]string{
				"params.object": {
					"url":    "abc.com",
					"commit": "af234",
				},
			},
		},
		expectedOutput: v1.NewObject(map[string]string{
			"url":    "abc.com",
			"commit": "af234",
		}),
	}, {
		name: "string replacement on object individual variables",
		args: args{
			input: v1.NewObject(map[string]string{
				"key1": "$(mystring)",
				"key2": "$(anotherObject.key)",
			}),
			stringReplacements: map[string]string{
				"mystring":          "foo",
				"anotherObject.key": "bar",
			},
		},
		expectedOutput: v1.NewObject(map[string]string{
			"key1": "foo",
			"key2": "bar",
		}),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.input.ApplyReplacements(tt.args.stringReplacements, tt.args.arrayReplacements, tt.args.objectReplacements)
			if d := cmp.Diff(tt.expectedOutput, tt.args.input); d != "" {
				t.Errorf("ApplyReplacements() output did not match expected value %s", diff.PrintWantGot(d))
			}
		})
	}
}

type ArrayOrStringHolder struct {
	AOrS v1.ArrayOrString `json:"val"`
}

func TestArrayOrString_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		input  map[string]interface{}
		result v1.ArrayOrString
	}{
		{
			input:  map[string]interface{}{"val": 123},
			result: *v1.NewArrayOrString("123"),
		},
		{
			input:  map[string]interface{}{"val": "123"},
			result: *v1.NewArrayOrString("123"),
		},
		{
			input:  map[string]interface{}{"val": ""},
			result: *v1.NewArrayOrString(""),
		},
		{
			input:  map[string]interface{}{"val": nil},
			result: v1.ArrayOrString{Type: v1.ParamTypeString, ArrayVal: nil},
		},
		{
			input:  map[string]interface{}{"val": []string{}},
			result: v1.ArrayOrString{Type: v1.ParamTypeArray, ArrayVal: []string{}},
		},
		{
			input:  map[string]interface{}{"val": []string{"oneelement"}},
			result: v1.ArrayOrString{Type: v1.ParamTypeArray, ArrayVal: []string{"oneelement"}},
		},
		{
			input:  map[string]interface{}{"val": []string{"multiple", "elements"}},
			result: v1.ArrayOrString{Type: v1.ParamTypeArray, ArrayVal: []string{"multiple", "elements"}},
		},
		{
			input:  map[string]interface{}{"val": map[string]string{"key1": "val1", "key2": "val2"}},
			result: v1.ArrayOrString{Type: v1.ParamTypeObject, ObjectVal: map[string]string{"key1": "val1", "key2": "val2"}},
		},
	}

	for _, c := range cases {
		for _, opts := range []func(enc *json.Encoder){
			// Default encoding
			func(enc *json.Encoder) {},
			// Multiline encoding
			func(enc *json.Encoder) { enc.SetIndent("", "  ") },
		} {
			b := new(bytes.Buffer)
			enc := json.NewEncoder(b)
			opts(enc)
			if err := enc.Encode(c.input); err != nil {
				t.Fatalf("error encoding json: %v", err)
			}

			var result ArrayOrStringHolder
			if err := json.Unmarshal(b.Bytes(), &result); err != nil {
				t.Errorf("Failed to unmarshal input '%v': %v", c.input, err)
			}
			if !reflect.DeepEqual(result.AOrS, c.result) {
				t.Errorf("expected %+v, got %+v", c.result, result)
			}
		}
	}
}

func TestArrayOrString_UnmarshalJSON_Directly(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		expected v1.ArrayOrString
	}{
		{desc: "empty value", input: ``, expected: *v1.NewArrayOrString("")},
		{desc: "int value", input: `1`, expected: *v1.NewArrayOrString("1")},
		{desc: "int array", input: `[1,2,3]`, expected: *v1.NewArrayOrString("[1,2,3]")},
		{desc: "nested array", input: `[1,\"2\",3]`, expected: *v1.NewArrayOrString(`[1,\"2\",3]`)},
		{desc: "string value", input: `hello`, expected: *v1.NewArrayOrString("hello")},
		{desc: "array value", input: `["hello","world"]`, expected: *v1.NewArrayOrString("hello", "world")},
		{desc: "object value", input: `{"hello":"world"}`, expected: *v1.NewObject(map[string]string{"hello": "world"})},
	}

	for _, c := range cases {
		aos := v1.ArrayOrString{}
		if err := aos.UnmarshalJSON([]byte(c.input)); err != nil {
			t.Errorf("Failed to unmarshal input '%v': %v", c.input, err)
		}
		if !reflect.DeepEqual(aos, c.expected) {
			t.Errorf("Failed to unmarshal input '%v': expected %+v, got %+v", c.input, c.expected, aos)
		}
	}
}

func TestArrayOrString_UnmarshalJSON_Error(t *testing.T) {
	cases := []struct {
		desc  string
		input string
	}{
		{desc: "empty value", input: "{\"val\": }"},
		{desc: "wrong beginning value", input: "{\"val\": @}"},
	}

	for _, c := range cases {
		var result ArrayOrStringHolder
		if err := json.Unmarshal([]byte(c.input), &result); err == nil {
			t.Errorf("Should return err but got nil '%v'", c.input)
		}
	}
}

func TestArrayOrString_MarshalJSON(t *testing.T) {
	cases := []struct {
		input  v1.ArrayOrString
		result string
	}{
		{*v1.NewArrayOrString("123"), "{\"val\":\"123\"}"},
		{*v1.NewArrayOrString("123", "1234"), "{\"val\":[\"123\",\"1234\"]}"},
		{*v1.NewArrayOrString("a", "a", "a"), "{\"val\":[\"a\",\"a\",\"a\"]}"},
		{*v1.NewObject(map[string]string{"key1": "var1", "key2": "var2"}), "{\"val\":{\"key1\":\"var1\",\"key2\":\"var2\"}}"},
	}

	for _, c := range cases {
		input := ArrayOrStringHolder{c.input}
		result, err := json.Marshal(&input)
		if err != nil {
			t.Errorf("Failed to marshal input '%v': %v", input, err)
		}
		if string(result) != c.result {
			t.Errorf("Failed to marshal input '%v': expected: %+v, got %q", input, c.result, string(result))
		}
	}
}

func TestArrayReference(t *testing.T) {
	tests := []struct {
		name, p, expectedResult string
	}{{
		name:           "valid array parameter expression with star notation returns param name",
		p:              "$(params.arrayParam[*])",
		expectedResult: "arrayParam",
	}, {
		name:           "invalid array parameter without dollar notation returns the input as is",
		p:              "params.arrayParam[*]",
		expectedResult: "params.arrayParam[*]",
	}}
	for _, tt := range tests {
		if d := cmp.Diff(tt.expectedResult, v1.ArrayReference(tt.p)); d != "" {
			t.Errorf(diff.PrintWantGot(d))
		}
	}
}
