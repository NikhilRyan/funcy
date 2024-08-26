
# Dynamic Function and Method Invocation in Go

This project implements a dynamic function and method invocation system in Go using reflection and a global registry. It allows users to invoke functions dynamically via HTTP by specifying the function name and parameters in a JSON payload.

## Key Features:
- **Automatic Function Registration**: Functions are registered globally during package initialization (`init()`) and stored in a registry for dynamic access.
- **Dynamic Invocation**: Functions can be invoked dynamically by passing the function name and parameters via HTTP requests.
- **Flexible Input Handling**: The system handles basic types (e.g., `int`, `string`), complex structs, slices, and maps. It converts JSON-encoded parameters (e.g., `map[string]interface{}`) to the appropriate Go types (e.g., custom structs).
- **Error Handling**: The system captures and returns errors that occur during function invocation, allowing for debugging and tracing.

## Project Structure:
```
your_project/
├── main.go               # Entry point of the application
├── go.mod                # Module definition
├── registry/
│   └── registry.go       # Global registry for functions and types
├── mypackage/
│   ├── data.go           # Defines the Data struct
│   ├── details.go        # Defines the Details struct
│   └── init.go           # Registers functions from mypackage
└── handlers/
    └── invoke_handler.go # HTTP handler for dynamic function invocation
```

## Example Functions:

#### `mypackage.Function1`
```go
func Function1(param string) ([]byte, Details, error)
```

#### `mypackage.Function2`
```go
func Function2(param string, num int) (Details, error)
```

#### `mypackage.Function3`
```go
func Function3(data Data) (Details, error)
```

#### `mypackage.Function4`
```go
func Function4(param1, param2 string, data Data) (string, Details, error)
```

## How It Works:

1. **Function Registration**: Functions are registered in the global registry via the `init()` function in each package. For example, the `mypackage` package registers its functions when imported.

2. **Dynamic Invocation**: The user sends a POST request to the `/invoke-function` endpoint with a JSON payload specifying the function name and parameters.

3. **Reflection-Based Invocation**: The system uses reflection to dynamically call the specified function and pass the parameters after converting them to the appropriate types (e.g., converting a `map[string]interface{}` to a `struct`).

4. **Response**: The result of the function invocation is returned as JSON, including any errors encountered.

## Example Payloads:

#### Example for `Function1`:
```json
{
    "type": "function",
    "func": "mypackage.Function1",
    "params": ["example"]
}
```

#### Example for `Function3`:
```json
{
    "type": "function",
    "func": "mypackage.Function3",
    "params": [{"ID": 10, "Name": "Test Data"}]
}
```

## Running the Project:

1. **Install Dependencies**:
   Run `go mod tidy` to install any dependencies specified in the `go.mod` file.

2. **Start the Server**:
   Run `go run main.go` to start the HTTP server.

3. **Invoke Functions**:
   Send POST requests to `http://localhost:8080/invoke-function` with the appropriate JSON payload to dynamically invoke functions.

## Example cURL Request:

```bash
curl -X POST http://localhost:8080/invoke-function -H "Content-Type: application/json" -d '{
    "type": "function",
    "func": "mypackage.Function3",
    "params": [{"ID": 10, "Name": "Test Data"}]
}'
```

## Future Improvements:
- **Struct Method Invocation**: Expand the system to handle dynamic method invocation on struct instances.
- **Additional Type Conversions**: Enhance the system to handle more complex type conversions, including nested structs and more advanced error reporting.
- **Performance Optimizations**: Optimize the reflection-based invocation process to reduce overhead.

## Conclusion:
This project showcases the power of reflection and dynamic invocation in Go, providing a flexible system for invoking functions dynamically based on JSON input. The system is suitable for scenarios like debugging, dynamic dispatch, and plugin architectures.
