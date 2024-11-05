package go_yamltools

import (
	"reflect"

	"gopkg.in/yaml.v2"
)

type NodeValueProcessor func(string, *[]string) bool

func TraverseMapSlice(value interface{}, state *[]string, processor NodeValueProcessor) {
	typeOf := reflect.TypeOf(value)
	if typeOf.Kind() == reflect.String {
		processor(value.(string), state)
	} else if typeOf.String() == "yaml.MapSlice" {
		for _, mapItem := range value.(yaml.MapSlice) {
			TraverseMapSlice(mapItem.Value, state, processor)
		}
	} else if typeOf.Kind() == reflect.Slice {
		for _, s := range value.([]any) {
			TraverseMapSlice(s, state, processor)
		}
	}
}
