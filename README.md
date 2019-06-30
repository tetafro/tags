# Tags

[![CircleCI](https://circleci.com/gh/tetafro/tags.svg?style=shield)](https://circleci.com/gh/tetafro/tags)
[![Codecov](https://codecov.io/gh/tetafro/tags/branch/master/graph/badge.svg)](https://codecov.io/gh/tetafro/tags)
[![Go Report](https://goreportcard.com/badge/github.com/tetafro/tags)](https://goreportcard.com/report/github.com/tetafro/tags)

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
