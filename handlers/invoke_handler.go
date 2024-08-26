package handlers

import (
	"Funcy/registry"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// RequestPayload represents the input JSON structure for the API.
type RequestPayload struct {
	Type   string        `json:"type"`   // "function"
	Func   string        `json:"func"`   // The function name to invoke
	Params []interface{} `json:"params"` // Parameters to pass to the function
}

// ResponsePayload represents the output JSON structure for the API.
type ResponsePayload struct {
	Result []interface{} `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}

// InvokeFunctionHandler handles HTTP requests to dynamically invoke functions.
func InvokeFunctionHandler(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := DynamicCallFunction(payload.Func, payload.Params...)
	response := ResponsePayload{}

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Result = results
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DynamicCallFunction dynamically calls a standalone function.
func DynamicCallFunction(functionName string, params ...interface{}) ([]interface{}, error) {
	fn, err := registry.GetFunction(functionName)
	if err != nil {
		return nil, err
	}

	// Prepare input parameters
	inputs, err := prepareInputs(fn, params)
	if err != nil {
		return nil, err
	}

	// Call the function
	results := fn.Call(inputs)

	// Handle the results
	return handleResults(results)
}

// prepareInputs converts and prepares the input parameters to match the function signature.
func prepareInputs(fn reflect.Value, params []interface{}) ([]reflect.Value, error) {
	fnType := fn.Type()
	inputs := make([]reflect.Value, len(params))

	for i, param := range params {
		expectedType := fnType.In(i)
		paramValue := reflect.ValueOf(param)

		// If the expected type is a struct, we need to manually convert the map[string]interface{} to the struct
		if expectedType.Kind() == reflect.Struct && paramValue.Kind() == reflect.Map {
			// Create a new instance of the expected struct type
			newStruct := reflect.New(expectedType).Elem()

			// Convert map[string]interface{} to struct by setting the fields
			err := mapToStruct(param.(map[string]interface{}), newStruct)
			if err != nil {
				return nil, fmt.Errorf("error converting map to struct: %v", err)
			}

			inputs[i] = newStruct
		} else if paramValue.Type().ConvertibleTo(expectedType) {
			// Handle basic type conversions (e.g., float64 to int)
			inputs[i] = paramValue.Convert(expectedType)
		} else {
			return nil, fmt.Errorf("cannot convert parameter %d from %s to %s", i+1, paramValue.Type(), expectedType)
		}
	}

	return inputs, nil
}

// mapToStruct maps a map[string]interface{} to a struct using reflection
func mapToStruct(data map[string]interface{}, result reflect.Value) error {
	for key, value := range data {
		// Get the field in the struct
		field := result.FieldByName(key)
		if !field.IsValid() {
			return fmt.Errorf("no such field: %s in struct", key)
		}

		// Check if the field can be set
		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", key)
		}

		// Set the value for the field, converting types if necessary
		val := reflect.ValueOf(value)
		if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		} else {
			return fmt.Errorf("cannot convert %v to %s", val.Type(), field.Type())
		}
	}

	return nil
}

// handleResults processes the return values from a function call and checks for errors.
func handleResults(results []reflect.Value) ([]interface{}, error) {
	outputs := make([]interface{}, len(results))

	for i, result := range results {
		outputs[i] = result.Interface()
	}

	// Check if the last result is an error
	if len(results) > 0 {
		if err, ok := outputs[len(outputs)-1].(error); ok && err != nil {
			return outputs[:len(outputs)-1], err
		}
	}

	return outputs, nil
}
