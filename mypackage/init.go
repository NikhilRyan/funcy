package mypackage

import (
	"Funcy/registry"
	"fmt"
	"reflect"
)

// Function1 Example functions with complex return types
func Function1(param string) ([]byte, Details, error) {
	if param == "" {
		return nil, Details{}, fmt.Errorf("param cannot be empty")
	}
	return []byte(param), Details{Description: "Example details", Value: 42}, nil
}

func Function2(param string, num int) (Details, error) {
	if num <= 0 {
		return Details{}, fmt.Errorf("num must be greater than zero")
	}
	return Details{Description: param, Value: num}, nil
}

// Function3 Fully qualify the 'Data' and 'Details' types in function signatures
func Function3(data Data) (Details, error) {
	if data.ID == 0 {
		return Details{}, fmt.Errorf("ID cannot be zero")
	}
	return Details{Description: data.Name, Value: data.ID}, nil
}

func Function4(param1, param2 string, data Data) (string, Details, error) {
	if param1 == "" || param2 == "" {
		return "", Details{}, fmt.Errorf("params cannot be empty")
	}
	return param1 + param2, Details{Description: data.Name, Value: data.ID}, nil
}

func init() {
	// Register functions from this package in the registry
	fmt.Println("Registering functions from mypackage")
	registry.RegisterFunction("mypackage.Function1", reflect.ValueOf(Function1))
	registry.RegisterFunction("mypackage.Function2", reflect.ValueOf(Function2))
	registry.RegisterFunction("mypackage.Function3", reflect.ValueOf(Function3))
	registry.RegisterFunction("mypackage.Function4", reflect.ValueOf(Function4))
}
