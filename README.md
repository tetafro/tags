# Tags

Simple package for parsing struct tags.

## Usage

```go
type Example struct {
	Field1 int    `tag:"field1"`
	Field2 string `tag:"field2"`
}

ex := Example{Field1: 10, Field2: "hello"}
tags, values := tags.Parse(ex, "tag") // []string{"field1", "field2"}, []interface{}{10, "hello"}
```
