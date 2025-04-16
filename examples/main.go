package main

import (
    "encoding/json"
    "fmt"
    "github.com/keda-github/go-json-transform/transform"
)

func main() {
    // Example input data
    input := map[string]interface{}{
        "user": map[string]interface{}{
            "firstName": "John",
            "lastName":  "Doe",
            "age":       30,
            "contact": map[string]interface{}{
                "email": "john@example.com",
                "phone": "1234567890",
            },
            "addresses": []interface{}{
                map[string]interface{}{
                    "type":    "home",
                    "street":  "123 Main St",
                    "city":    "Boston",
                    "country": "USA",
                },
                map[string]interface{}{
                    "type":    "work",
                    "street":  "456 Corp Ave",
                    "city":    "New York",
                    "country": "USA",
                },
            },
        },
    }

    // Define mapping
    mapping := transform.Mapping{
        Map: map[string]interface{}{
            "fullName": transform.Operation{
                Fields: []string{"user.firstName", "user.lastName"},
                Operation: func(values ...interface{}) interface{} {
                    return fmt.Sprintf("%s %s", values[0], values[1])
                },
            },
            "age":     "user.age",
            "contact": transform.Field{
                Path: "user.contact.email",
            },
            "addresses": transform.ArrayMapping{
                Path: "user.addresses",
                Mapping: map[string]interface{}{
                    "addressType": "type",
                    "location":    "city",
                    "country":    "country",
                },
            },
        },
    }

    // Transform data
    transformer := transform.New()
    result, err := transformer.Transform(input, mapping)
    if err != nil {
        panic(err)
    }

    // Pretty print the result
    jsonResult, _ := json.MarshalIndent(result, "", "    ")
    fmt.Println(string(jsonResult))
} 