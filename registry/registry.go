package registry

import (
	"fmt"
	"reflect"
	"sync"
)

// Registry stores the reflect.Type and reflect.Value for structs and functions across the project
type Registry struct {
	mu        sync.RWMutex
	types     map[string]reflect.Type
	functions map[string]reflect.Value
}

// globalRegistry is the shared instance of the registry for the entire project
var globalRegistry = &Registry{
	types:     make(map[string]reflect.Type),
	functions: make(map[string]reflect.Value),
}

// RegisterType registers a struct type globally in the registry
func RegisterType(name string, typ reflect.Type) {
	globalRegistry.mu.Lock()
	defer globalRegistry.mu.Unlock()
	globalRegistry.types[name] = typ
}

// RegisterFunction registers a function globally in the registry
func RegisterFunction(name string, fn reflect.Value) {
	globalRegistry.mu.Lock()
	defer globalRegistry.mu.Unlock()
	globalRegistry.functions[name] = fn
}

// GetType retrieves a type by name from the registry
func GetType(name string) (reflect.Type, error) {
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()
	if typ, ok := globalRegistry.types[name]; ok {
		return typ, nil
	}
	return nil, fmt.Errorf("type %s not found", name)
}

// GetFunction retrieves a function by name from the registry
func GetFunction(name string) (reflect.Value, error) {
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()
	if fn, ok := globalRegistry.functions[name]; ok {
		return fn, nil
	}
	return reflect.Value{}, fmt.Errorf("function %s not found", name)
}
