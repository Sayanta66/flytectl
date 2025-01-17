// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package create

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (ProjectConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (ProjectConfig) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (ProjectConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in ProjectConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg ProjectConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("ProjectConfig", pflag.ExitOnError)
	cmdFlags.StringVar(&projectConfig.ID, fmt.Sprintf("%v%v", prefix, "id"), projectConfig.ID, "id for the project specified as argument.")
	cmdFlags.StringVar(&projectConfig.Name, fmt.Sprintf("%v%v", prefix, "name"), projectConfig.Name, "name for the project specified as argument.")
	cmdFlags.StringVar(&projectConfig.File, fmt.Sprintf("%v%v", prefix, "file"), projectConfig.File, "file for the project definition.")
	cmdFlags.StringVar(&projectConfig.Description, fmt.Sprintf("%v%v", prefix, "description"), projectConfig.Description, "description for the project specified as argument.")
	cmdFlags.StringToStringVar(&projectConfig.Labels, fmt.Sprintf("%v%v", prefix, "labels"), projectConfig.Labels, "labels for the project specified as argument.")
	cmdFlags.BoolVar(&projectConfig.DryRun, fmt.Sprintf("%v%v", prefix, "dryRun"), projectConfig.DryRun, "execute command without making any modifications.")
	return cmdFlags
}
