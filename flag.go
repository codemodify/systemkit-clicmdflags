package clicmdflags

import (
	"reflect"
	"strings"
)

var flagPatterns = []string{
	"--",
	"-",
}

// {flagRequired} and {flagDefault} are MUTUALLY EXCLUSIVE
const flagName = "flagName"               //
const flagRequired = "flagRequired"       // required - needs inpuut from user
const flagDefault = "flagDefault"         // default - would be the value if not set
const flagDescription = "flagDescription" //
const flagHidden = "flagHidden"           //

type flag struct {
	name                     string
	typeName                 string
	isRequired               string
	defaultValue             string
	defaultValueWasRequested bool
	description              string
	isHidden                 string

	wasSet bool
}

func isArgFlag(arg string) (bool, string) {
	for _, flagPattern := range flagPatterns {
		if strings.HasPrefix(arg, flagPattern) {
			return true, strings.Replace(arg, flagPattern, "", 1)
		}
	}

	return false, ""
}

func (thisRef *Command) getDefinedFlags() []flag {
	result := []flag{}

	runtimeStructRef, err := getRealRuntimeStruct(reflect.ValueOf(thisRef.Flags))
	if err != nil {
		return result
	}

	for i := 0; i < runtimeStructRef.NumField(); i++ {
		newFlag := flag{
			name:         runtimeStructRef.Type().Field(i).Tag.Get(flagName),
			typeName:     runtimeStructRef.Type().Field(i).Type.Name(),
			isRequired:   runtimeStructRef.Type().Field(i).Tag.Get(flagRequired),
			defaultValue: runtimeStructRef.Type().Field(i).Tag.Get(flagDefault),
			description:  runtimeStructRef.Type().Field(i).Tag.Get(flagDescription),
			isHidden:     runtimeStructRef.Type().Field(i).Tag.Get(flagHidden),
		}

		if _, tagFound := runtimeStructRef.Type().Field(i).Tag.Lookup(flagDefault); tagFound {
			newFlag.defaultValueWasRequested = true
		}

		if len(strings.TrimSpace(newFlag.isRequired)) <= 0 {
			newFlag.isRequired = "false"
		}

		result = append(result, newFlag)
	}

	return result
}

func (thisRef *Command) getRequriedFlags() []flag {
	result := []flag{}

	for _, definedFlag := range thisRef.getDefinedFlags() {
		if definedFlag.isRequired == "true" {
			result = append(result, definedFlag)
		}
	}

	return result
}
