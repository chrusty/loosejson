package loosejson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal into a struct with permissive type-conversion:
func Unmarshal(jsonBytes []byte, structInterface interface{}) error {

	// Something safe to unmarshal the JSON into:
	var mapOfInterfaces map[string]interface{}

	// Unmarshal into the mapOfInterfaces:
	err := json.Unmarshal(jsonBytes, &mapOfInterfaces)
	if err != nil {
		return err
	}

	// Reflect the struct we were given (values and types):
	structValues := reflect.ValueOf(structInterface)
	structTypes := reflect.TypeOf(structInterface)

	// Check that we were given a pointer to something real:
	if structValues.Kind() != reflect.Ptr || structValues.IsNil() {
		return errors.New(fmt.Sprintf("Provided interface is either nil or not a pointer"))
	}

	// Go through each field, and attempt to fill it in from our mapOfInterfaces:
	for i := 0; i < structValues.Elem().NumField(); i++ {

		// Get the field:
		fieldType := structTypes.Elem().Field(i)
		fieldValue := structValues.Elem().Field(i)

		// Split up the JSON tags:
		jsonTags := strings.Split(fieldType.Tag.Get("json"), ",")
		if jsonTags[0] != "-" {

			// Get the JSON field-name (from the tags):
			jsonFieldName := strings.Split(fieldType.Tag.Get("json"), ",")[0]

			// Behave differently according to which type the field is:
			switch fieldType.Type.String() {

			// This struct-field is an int:
			case "int", "int32", "int64", "*int", "*int32", "*int64":
				var jsonValue int64 = 0
				switch mapOfInterfaces[jsonFieldName].(type) {
				case string:
					// Convert a string to an int:
					jsonValue, err = strconv.ParseInt(mapOfInterfaces[jsonFieldName].(string), 0, 64)
				case float32, float64:
					// Convert a float to an int:
					jsonValue = int64(mapOfInterfaces[jsonFieldName].(float64))
				case bool:
					// Convert a bool to an int:
					if mapOfInterfaces[jsonFieldName].(bool) {
						jsonValue = 1
					}
				}
				if err != nil {
					return errors.New(fmt.Sprintf("Can't convert '%v' to int!", mapOfInterfaces[jsonFieldName]))
				} else {
					// See if we're dealing with a pointer:
					if fieldType.Type.Kind() == reflect.Ptr {
						// Set the field value to nil (correct for its type), then set the pointer:
						fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
						fieldValue.Elem().SetInt(jsonValue)
					} else {
						// Set the field to have an int value (directly):
						fieldValue.SetInt(jsonValue)
					}
				}

			// This struct-field is a float:
			case "float32", "float64", "*float32", "*float64":
				var jsonValue float64 = 0.0
				switch mapOfInterfaces[jsonFieldName].(type) {
				case string:
					// Convert a string to a float:
					jsonValue, err = strconv.ParseFloat(mapOfInterfaces[jsonFieldName].(string), 64)
				case float32, float64:
					// Just take a float:
					jsonValue = mapOfInterfaces[jsonFieldName].(float64)
				case bool:
					// Convert a bool to a float:
					if mapOfInterfaces[jsonFieldName].(bool) {
						jsonValue = 1.0
					}
				}
				if err != nil {
					return errors.New(fmt.Sprintf("Can't convert '%v' to float!", mapOfInterfaces[jsonFieldName]))
				} else {
					// See if we're dealing with a pointer:
					if fieldType.Type.Kind() == reflect.Ptr {
						// Set the field value to nil (correct for its type), then set the pointer:
						fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
						fieldValue.Elem().SetFloat(jsonValue)
					} else {
						// Set the field to have a float value (directly):
						fieldValue.SetFloat(jsonValue)
					}
				}

			// This struct-field is a string:
			case "string", "*string":
				var jsonValue string
				switch mapOfInterfaces[jsonFieldName].(type) {
				case string:
					// Just take a string:
					jsonValue = mapOfInterfaces[jsonFieldName].(string)
				case float32, float64:
					// Convert a float to a string:
					jsonValue = strconv.FormatFloat(mapOfInterfaces[jsonFieldName].(float64), 'f', -1, 64)
				case bool:
					// Convert a bool to a string:
					jsonValue = strconv.FormatBool(mapOfInterfaces[jsonFieldName].(bool))
				}
				if err != nil {
					return errors.New(fmt.Sprintf("Can't convert '%v' to string!", mapOfInterfaces[jsonFieldName]))
				} else {
					// See if we're dealing with a pointer:
					if fieldType.Type.Kind() == reflect.Ptr {
						// Set the field value to nil (correct for its type), then set the pointer:
						fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
						fieldValue.Elem().SetString(jsonValue)
					} else {
						// Set the field to have a string value (directly):
						fieldValue.SetString(jsonValue)
					}
				}

			// This struct-field is a bool:
			case "bool", "*bool":
				var jsonValue bool = false
				switch mapOfInterfaces[jsonFieldName].(type) {
				case string:
					// Convert a string to a bool:
					jsonValue, err = strconv.ParseBool(mapOfInterfaces[jsonFieldName].(string))
				case float32, float64:
					// Convert a float to a bool:
					if mapOfInterfaces[jsonFieldName].(float64) > 0.5 {
						jsonValue = true
					}
				case bool:
					// Just take a bool:
					jsonValue = mapOfInterfaces[jsonFieldName].(bool)
				}
				if err != nil {
					return errors.New(fmt.Sprintf("Can't convert '%v' to bool!", mapOfInterfaces[jsonFieldName]))
				} else {
					// See if we're dealing with a pointer:
					if fieldType.Type.Kind() == reflect.Ptr {
						// Set the field value to nil (correct for its type), then set the pointer:
						fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
						fieldValue.Elem().SetBool(jsonValue)
					} else {
						// Set the field to have a bool value (directly):
						fieldValue.SetBool(jsonValue)
					}
				}

			// Or something else:
			default:
				return errors.New(fmt.Sprintf("Can't handle attribute %v (type %v)", fieldType.Name, fieldType.Type))
			}
		}
	}

	// Return:
	return nil
}
