package clicmdflags

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func (thisRef *Command) flagNeededCommandsForExecuteAndPopulateTheirFlags(args []string) error {
	var errorToReturn error
	updateErrorToReturn := func(err error) {
		if errorToReturn == nil {
			errorToReturn = err
		}
	}

	definedFlags := thisRef.getDefinedFlags()
	setFlagWithDefaultValue := func(flagName string) {
		for i := 0; i < len(definedFlags); i++ {
			if definedFlags[i].name == flagName {
				thisRef.setFlagValue(flagName, definedFlags[i].defaultValue)
				break
			}
		}
	}
	// set all flags to their defined defaults
	for _, definedFlag := range definedFlags {
		if definedFlag.defaultValueWasRequested {
			setFlagWithDefaultValue(definedFlag.name)
		}
	}

	requriedFlags := thisRef.getRequriedFlags()
	setRequriedFlagAsSet := func(flagName string) {
		for i := 0; i < len(requriedFlags); i++ {
			if requriedFlags[i].name == flagName {
				requriedFlags[i].wasSet = true
				break
			}
		}
	}

	subCommandFoundAlready := false
	for index := 0; index < len(args); index++ {

		// case 1 - FLAG
		if isFlag, flagName := isArgFlag(args[index]); isFlag {

			// This is a flag AND it's not last item
			nextIndex := index + 1
			if nextIndex < len(args) {

				// case 1 - next is another FLAG
				if nextIsFlag, _ := isArgFlag(args[nextIndex]); nextIsFlag {
					// if this is a `bool` flag and the value is missing then set the value to `true`
					setFlagWithDefaultValue(flagName)
					if thisRef.getFlag(flagName).Kind() == reflect.Bool {
						thisRef.setFlagValue(flagName, "true")
						setRequriedFlagAsSet(flagName)
					}

					continue
				}

				// case 2 - next is a COMMAND
				if isCommand, _ := isArgCommand(args[nextIndex], thisRef); isCommand {

					// if this is a `bool` flag and the value is missing then set the value to `true`
					setFlagWithDefaultValue(flagName)
					if thisRef.getFlag(flagName).Kind() == reflect.Bool {
						thisRef.setFlagValue(flagName, "true")
						setRequriedFlagAsSet(flagName)
					}

					continue
				}

				// case 3 - this is a flag value, then set the value to the current `Flags` struct
				flagValue := args[nextIndex]
				thisRef.setFlagValue(flagName, flagValue)
				setRequriedFlagAsSet(flagName)
				index = nextIndex
				continue
			}

			// This is a flag AND it's the last item
			// if this is a `bool` flag and the value is missing then set the value to `true`
			setFlagWithDefaultValue(flagName)
			if thisRef.getFlag(flagName).Kind() == reflect.Bool {
				thisRef.setFlagValue(flagName, "true")
				setRequriedFlagAsSet(flagName)
			}

			continue
		}

		// case 2 - COMMAND
		if isCommand, subCommand := isArgCommand(args[index], thisRef); isCommand && !subCommandFoundAlready {
			subCommandFoundAlready = true

			argsToProcess := []string{}
			nextIndex := index + 1
			if nextIndex < len(args) {
				argsToProcess = args[nextIndex:]
			}

			// commandPath := thisRef.getCommandPath()
			subCommand.flagedForExecute = true
			if err := subCommand.flagNeededCommandsForExecuteAndPopulateTheirFlags(argsToProcess); err != nil {
				updateErrorToReturn(err)
			}

			continue
		}

		// case 3 - uknown COMMAND or FLAG
		// ignore
	}

	// check that all required flags are set
	for _, rf := range requriedFlags {
		if !rf.wasSet {
			updateErrorToReturn(fmt.Errorf("Missing required flag [%s]", rf.name))
			break
		}
	}

	return errorToReturn
}

func isArgCommand(arg string, command *Command) (bool, *Command) {
	if command.Name == arg {
		return true, command
	}

	for _, c := range command.subCommands {
		if c.Name == arg {
			return true, c
		}
	}

	for _, c := range command.subCommands {
		if c.PassThrough {
			for _, subC := range c.subCommands {
				if subC.Name == arg {
					return true, subC
				}
			}
		}
	}

	return false, nil
}

func getRealRuntimeStruct(structByValOrRef reflect.Value) (reflect.Value, error) {
	// 1. interface{} underlying data is value   and receiver is value
	// 2. interface{} underlying data is value   and receiver is pointer
	// 3. interface{} underlying data is pointer and receiver is value
	// 4. interface{} underlying data is pointer and receiver is pointer

	// Using reflection we can determine the underling value of our interface.
	// Using reflection we can generate the alternate data type to our current type.
	// If the data passed in was a value we need to generate a pointer to it.

	if !structByValOrRef.IsValid() ||
		(reflect.ValueOf(structByValOrRef).Kind() == reflect.Ptr && reflect.ValueOf(structByValOrRef).IsNil()) {

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
	runtimeStructRef, err := getRealRuntimeStruct(value)
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
			setFieldValue(field, flagValue)
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

	runtimeStructRef, err := getRealRuntimeStruct(reflect.ValueOf(thisRef.Flags))
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

func setFieldValue(field reflect.Value, valueAsString string) {
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

func (thisRef *Command) getLastSubcommandFlagedForExecute() *Command {
	if len(thisRef.subCommands) <= 0 {
		return thisRef
	}

	for _, c := range thisRef.subCommands {
		if c.flagedForExecute {
			return c.getLastSubcommandFlagedForExecute()
		}
	}

	for _, c := range thisRef.subCommands {
		if c.PassThrough {
			for _, subC := range c.subCommands {
				if subC.flagedForExecute {
					return c.getLastSubcommandFlagedForExecute()
				}
			}
		}
	}

	return thisRef
}

func (thisRef *Command) getRootCommand() *Command {
	rootCommand := thisRef
	for {
		if rootCommand.parentCommand == nil {
			break
		}

		rootCommand = rootCommand.parentCommand
	}

	return rootCommand
}

func (thisRef *Command) getCommandPath() []string {
	path := []string{
		thisRef.Name,
	}

	cmd := thisRef.parentCommand
	for {
		if cmd == nil {
			break
		}

		path = append(path, cmd.Name)
	}

	return path
}

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// ~~~~ DEBUG Helpers
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

func structToString(s interface{}) string {
	d, _ := json.Marshal(s)

	return string(d)
}

// DEBUGDumpCommandFlags -
func DEBUGDumpCommandFlags(command *Command) {
	fmt.Println(structToString(command.Flags))

	for _, cmd := range command.subCommands {
		DEBUGDumpCommandFlags(cmd)
	}
}
