# go-json-transform

A powerful and flexible JSON transformation library for Go, inspired by node-json-transform. This package helps you transform JSON data structures from one format to another using predefined mapping rules.

## Features

- Transform JSON data based on custom mapping rules
- Support for nested object transformations
- Array mapping and transformation
- Custom value operations
- Default value handling
- Conditional transformations

## Installation

```bash
go get github.com/keda-github/go-json-transform
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/keda-github/go-json-transform/transform"
)

func main() {
    // Input data
    input := map[string]interface{}{
        "name": "John Doe",
        "age":  30,
        "contact": map[string]interface{}{
            "email": "john@example.com",
            "phone": "1234567890",
        },
    }

    // Define mapping
    mapping := transform.Mapping{
        Map: map[string]interface{}{
            "fullName":     "name",
            "userAge":      "age",
            "emailAddress": "contact.email",
        },
    }

    // Create transformer
    transformer := transform.New()

    // Transform data
    result, err := transformer.Transform(input, mapping)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", result)
}
```

### Output
```json
{
    "fullName": "John Doe",
    "userAge": 30,
    "emailAddress": "john@example.com"
}
```

### Advanced Usage

#### Array Mapping

```go
input := map[string]interface{}{
    "users": []map[string]interface{}{
        {"name": "John", "age": 30},
        {"name": "Jane", "age": 25},
    },
}

mapping := transform.Mapping{
    Map: map[string]interface{}{
        "people": transform.ArrayMapping{
            Path: "users",
            Mapping: map[string]interface{}{
                "fullName": "name",
                "userAge":  "age",
            },
        },
    },
}
```

#### Default Values

```go
mapping := transform.Mapping{
    Map: map[string]interface{}{
        "status": transform.Field{
            Path:         "status",
            DefaultValue: "active",
        },
    },
}
```

#### Custom Operations

```go
mapping := transform.Mapping{
    Map: map[string]interface{}{
        "fullName": transform.Operation{
            Fields: []string{"firstName", "lastName"},
            Operation: func(values ...interface{}) interface{} {
                return fmt.Sprintf("%s %s", values[0], values[1])
            },
        },
    },
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
