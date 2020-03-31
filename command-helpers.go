package clicmdflags

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func (thisRef *Command) flagCmdForExecuteAndPopulateTheirFlags(args []string) {
	for index := 0; index < len(args); index++ {
		if isFlag, flagName := thisRef.isFlag(args[index]); isFlag { // case 1 - FLAG
			nextIndex := index + 1
			if nextIndex < len(args) { // if not reached last item
				// case 1 - next is another FLAG
				if nextIsFlag, _ := thisRef.isFlag(args[nextIndex]); nextIsFlag {
					// if this is a `bool` flag and the value is missing then set the value to `true`
					if thisRef.getFlag(flagName).Kind() == reflect.Bool {
						thisRef.setFlagValue(flagName, "true")
					}

					continue
				}

				// case 2 - next is a COMMAND
				if isCommand, _ := thisRef.isCommand(args[nextIndex]); isCommand {
					continue
				}

				// case 3 - this is a flag value, then set the value to the current `Flags` struct
				flagValue := args[nextIndex]
				thisRef.setFlagValue(flagName, flagValue)
				index = nextIndex
			} else {
				// INFO: nothing to do as default value is already set by the user in the struct
			}
		} else if isCommand, subCommand := thisRef.isCommand(args[index]); isCommand { // case 2 - COMMAND
			argsToProcess := []string{}
			nextIndex := index + 1
			if nextIndex < len(args) {
				argsToProcess = args[nextIndex:]
			}
			subCommand.flagedForExecute = true
			subCommand.flagCmdForExecuteAndPopulateTheirFlags(argsToProcess)
		} else { // case 2 - UNKNOWN command or flag
			// ignore
		}
	}
}

func (thisRef *Command) isCommand(arg string) (bool, *Command) {
	if thisRef.Name == arg {
		return true, thisRef
	}

	for _, c := range thisRef.subCommands {
		if c.Name == arg {
			return true, c
		}
	}

	return false, nil
}

func (thisRef *Command) isFlag(arg string) (bool, string) {
	if strings.HasPrefix(arg, flagPattern) {
		return true, strings.Replace(arg, flagPattern, "", 1)
	}

	return false, ""
}

func (thisRef *Command) getRealRuntimeStruct(structByValOrRef reflect.Value) (reflect.Value, error) {
	// 1. interface{} underlying data is value   and receiver is value
	// 2. interface{} underlying data is value   and receiver is pointer
	// 3. interface{} underlying data is pointer and receiver is value
	// 4. interface{} underlying data is pointer and receiver is pointer

	// Using reflection we can determine the underling value of our interface.
	// Using reflection we can generate the alternate data type to our current type.
	// If the data passed in was a value we need to generate a pointer to it.

	if isNil(structByValOrRef) || thisRef.Flags == nil {
		return reflect.New(reflect.TypeOf("")), errors.New("THE THING IS NULL")
	}

	isPointer := (structByValOrRef.Type().Kind() == reflect.Ptr)

	var runtimeStructRef reflect.Value // represents the run-time data
	if isPointer {
		// acquire value referenced by pointer
		runtimeStructRef = structByValOrRef.Elem()
	} else {
		// set the element's value once there is a pointer to the correct type
		pointerToRuntimeType := reflect.New(structByValOrRef.Type())
		pointerToRuntimeType.Elem().Set(structByValOrRef)
		runtimeStructRef = pointerToRuntimeType.Elem()
	}

	return runtimeStructRef, nil
}

func (thisRef *Command) setFlagValue(flagName string, flagValue string) {
	value := reflect.ValueOf(thisRef.Flags)
	runtimeStructRef, err := thisRef.getRealRuntimeStruct(value)
	if err != nil {
		return
	}

	isPointer := (value.Kind() == reflect.Ptr)

	// update the struct field value
	updateHappened := false
	for i := 0; i < runtimeStructRef.NumField(); i++ {
		attrFlagName := runtimeStructRef.Type().Field(i).Tag.Get("flagName")
		if attrFlagName == flagName {
			field := runtimeStructRef.Field(i)
			thisRef.setFieldValue(field, flagValue)
			updateHappened = true
			break
		}
	}

	// if this is not a pointer then this is a copy, we must put it back in the original
	if updateHappened && !isPointer {
		thisRef.Flags = runtimeStructRef.Interface()
	}
}

func (thisRef *Command) getFlag(flagName string) reflect.Value {
	var result reflect.Value

	runtimeStructRef, err := thisRef.getRealRuntimeStruct(reflect.ValueOf(thisRef.Flags))
	if err != nil {
		return result
	}

	for i := 0; i < runtimeStructRef.NumField(); i++ {
		attrFlagName := runtimeStructRef.Type().Field(i).Tag.Get("flagName")
		if attrFlagName == flagName {
			result = runtimeStructRef.Field(i)
			break
		}
	}

	return result
}

func (thisRef *Command) setFieldValue(field reflect.Value, valueAsString string) {
	switch field.Kind() {
	case reflect.Invalid:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		fallthrough
	case reflect.Array:
		fallthrough
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Struct:
		fallthrough
	case reflect.UnsafePointer:
		// INFO: ignore

	case reflect.Bool:
		value, _ := strconv.ParseBool(valueAsString)
		field.SetBool(value)

	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value, _ := strconv.ParseInt(valueAsString, 10, 64)
		field.SetInt(value)

	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value, _ := strconv.ParseUint(valueAsString, 10, 64)
		field.SetUint(value)

	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, _ := strconv.ParseFloat(valueAsString, 64)
		field.SetFloat(value)

	case reflect.String:
		field.SetString(valueAsString)
	}
}

func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}

func (thisRef *Command) getLastSubcommandFlagedForExecute() *Command {
	if len(thisRef.subCommands) <= 0 {
		return thisRef
	}

	for _, c := range thisRef.subCommands {
		if c.flagedForExecute && c != helpCmd {
			return c.getLastSubcommandFlagedForExecute()
		}
	}

	return thisRef
}

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// ~~~~ DEBUG Helpers
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

func structToString(s interface{}) string {
	d, _ := json.Marshal(s)

	return string(d)
}

// DEBUGDumpCommandFlags -
func DEBUGDumpCommandFlags(command *Command) {
	fmt.Println(structToString(command.Flags))

	for _, cmd := range command.subCommands {
		if cmd != helpCmd {
			DEBUGDumpCommandFlags(cmd)
		}
	}
}
