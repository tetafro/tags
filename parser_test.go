package tags

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var (
		testPtr  = &testOneField{Int: 1}
		testFunc = func() {}
		testChan = make(chan int)
	)

	testCases := []struct {
		name   string
		obj    interface{}
		tags   []string
		values []interface{}
	}{
		{
			name:   "struct with one field",
			obj:    testOneField{Int: 10},
			tags:   []string{"int"},
			values: []interface{}{int(10)},
		},
		{
			name:   "pointer to struct with one field",
			obj:    &testOneField{Int: 10},
			tags:   []string{"int"},
			values: []interface{}{int(10)},
		},
		{
			name: "struct with many fields",
			obj: testManyFields{
				Empty:      1,
				Dash:       2,
				Int:        3,
				String:     "hello",
				Float64:    1.5,
				Array:      [1]int{1},
				Chan:       testChan,
				Func:       testFunc,
				Map:        map[int]int{1: 2},
				Ptr:        testPtr,
				Slice:      []int{1, 2, 3},
				Struct:     testOneField{Int: 5},
				unexported: 10,
			},
			tags: []string{
				"int", "string", "float64",
				"array", "chan", "func",
				"map", "ptr", "slice",
				"struct",
			},
			values: []interface{}{
				int(3), "hello", float64(1.5),
				[1]int{1}, testChan, testFunc,
				map[int]int{1: 2}, testPtr, []int{1, 2, 3},
				testOneField{Int: 5},
			},
		},
		{
			name: "struct with no suitable fields",
			obj: testNoSuitableFields{
				Empty:      1,
				Dash:       2,
				unexported: 10,
			},
			tags:   nil,
			values: nil,
		},
		{
			name: "anonymouse struct",
			obj: struct {
				Int int `example:"int"`
			}{
				Int: 10,
			},
			tags:   []string{"int"},
			values: []interface{}{int(10)},
		},
		{
			name:   "struct with no fields",
			obj:    testNoFields{},
			tags:   nil,
			values: nil,
		},
		{
			name:   "pointer to struct with no fields",
			obj:    testNoFields{},
			tags:   nil,
			values: nil,
		},
		{
			name:   "integer",
			obj:    int(10),
			tags:   nil,
			values: nil,
		},
		{
			name:   "slice",
			obj:    []int{10},
			tags:   nil,
			values: nil,
		},
		{
			name:   "map",
			obj:    map[int]int{10: 10},
			tags:   nil,
			values: nil,
		},
		{
			name:   "func",
			obj:    func() {},
			tags:   nil,
			values: nil,
		},
		{
			name:   "nil",
			obj:    nil,
			tags:   nil,
			values: nil,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tags, values := Parse(tt.obj, testTag)

			if len(tt.tags) != len(tags) {
				t.Fatalf(
					"Wrong tags number:\n  expected: %d\n       got: %d",
					len(tt.tags), len(tags),
				)
			}
			for i := 0; i < len(tt.tags); i++ {
				if tt.tags[i] != tags[i] {
					t.Fatalf(
						"Wrong tag:\n  expected: %s\n       got: %s",
						tt.tags[i], tags[i],
					)
				}
			}

			if len(tt.values) != len(values) {
				t.Fatalf(
					"Wrong values number:\n  expected: %d\n       got: %d",
					len(tt.values), len(values),
				)
			}
			for i := 0; i < len(tt.values); i++ {
				// Compare functions (DeepEqual doesn't work for it)
				if reflect.TypeOf(tt.values[i]).Kind() == reflect.Func {
					// Check type
					typ := reflect.TypeOf(values[i])
					if typ.Kind() != reflect.Func {
						t.Fatalf(
							"Wrong value for tag %s:\n  expected: function\n       got: %v",
							tt.tags[i], typ,
						)
					}
					// Check pointers
					p1 := reflect.ValueOf(tt.values[i]).Pointer()
					p2 := reflect.ValueOf(tt.values[i]).Pointer()
					if p1 != p2 {
						t.Fatalf(
							"Wrong value for tag %s:\n  functions are not the same",
							tt.tags[i],
						)
					}
					continue
				}
				// Compare anything else
				if !reflect.DeepEqual(tt.values[i], values[i]) {
					t.Fatalf(
						"Wrong value for tag %s:\n  expected: %v\n       got: %v",
						tt.tags[i], tt.values[i], values[i],
					)
				}
			}
		})
	}
}

const testTag = "example"

type testOneField struct {
	Int int `example:"int"`
}

type testManyFields struct {
	Empty      int
	Dash       int           `example:"-"`
	Int        int           `example:"int"`
	String     string        `example:"string"`
	Float64    float64       `example:"float64"`
	Array      [1]int        `example:"array"`
	Chan       chan int      `example:"chan"`
	Func       func()        `example:"func"`
	Map        map[int]int   `example:"map"`
	Ptr        *testOneField `example:"ptr"`
	Slice      []int         `example:"slice"`
	Struct     testOneField  `example:"struct"`
	unexported int           `example:"unexported"`
}

type testNoSuitableFields struct {
	Dash       int `example:"-"`
	Empty      int
	unexported int `example:"unexported"`
}

type testNoFields struct{}
