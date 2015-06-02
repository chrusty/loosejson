package loosejson

import (
	"fmt"
	"testing"
)

const (
	jsonBools           = `{"boolAsBool": true, "boolAsFloat": 1.2, "boolAsString": "true", "boolAsBoolPtr": true}`
	jsonFloats          = `{"floatAsBool": true, "floatAsFloat": 3.14, "floatAsString": "3.14", "floatAsFloatPtr": 3.14}`
	jsonInts            = `{"intAsBool": true, "intAsFloat": 1.0, "intAsString": "1", "intAsIntPtr": 1}`
	jsonStrings         = `{"stringAsBool": true, "stringAsFloat": 3.14, "stringAsString": "one", "stringAsStringPtr": "one"}`
	jsonCapitalisations = `{"boolWIthLowerCase": "true", "floatWithLowerCase": "37.75", "IntWithUpperCase": "123", "StringWithUpperCase": 99.9999}`
)

type TestObject struct {
	BoolAsBool          bool     `json:"boolAsBool"`
	BoolAsFloat         bool     `json:"boolAsFloat"`
	BoolAsString        bool     `json:"boolAsString"`
	BoolAsBoolPtr       *bool    `json:"boolAsBoolPtr"`
	FloatAsBool         float32  `json:"floatAsBool"`
	FloatAsFloat        float32  `json:"floatAsFloat"`
	FloatAsString       float32  `json:"floatAsString"`
	FloatAsFloatPtr     *float32 `json:"floatAsFloatPtr"`
	IntAsBool           int32    `json:"intAsBool"`
	IntAsFloat          int32    `json:"intAsFloat"`
	IntAsString         int32    `json:"intAsString"`
	IntAsIntPtr         *int32   `json:"intAsIntPtr"`
	StringAsBool        string   `json:"stringAsBool"`
	StringAsFloat       string   `json:"stringAsFloat"`
	StringAsString      string   `json:"stringAsString"`
	StringAsStringPtr   *string  `json:"stringAsStringPtr"`
	BoolWIthLowerCase   bool
	FloatWithLowerCase  float32
	IntWithUpperCase    int32
	StringWithUpperCase string
	IgnoreField         []byte `json:"-"`
}

func TestBools(t *testing.T) {
	// Make a test struct:
	testStruct := TestObject{}

	// Unmarshal:
	err := Unmarshal([]byte(jsonBools), &testStruct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Test bools:
	if testStruct.BoolAsBool != true {
		fmt.Println("Unable to unmarshal bool into bool!")
		t.Fail()
	}
	if testStruct.BoolAsFloat != true {
		fmt.Println("Unable to unmarshal float into bool!")
		t.Fail()
	}
	if testStruct.BoolAsString != true {
		fmt.Println("Unable to unmarshal string into bool!")
		t.Fail()
	}
	if *testStruct.BoolAsBoolPtr != true {
		fmt.Println("Unable to unmarshal bool into *bool!")
		t.Fail()
	}
}

func TestFloats(t *testing.T) {
	// Make a test struct:
	testStruct := TestObject{}

	// Unmarshal:
	err := Unmarshal([]byte(jsonFloats), &testStruct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Test floats:
	if testStruct.FloatAsBool != 1 {
		fmt.Println("Unable to unmarshal bool into float32!")
		t.Fail()
	}
	if testStruct.FloatAsFloat != 3.14 {
		fmt.Println("Unable to unmarshal float into float32!")
		t.Fail()
	}
	if testStruct.FloatAsString != 3.14 {
		fmt.Println("Unable to unmarshal string into float32!")
		t.Fail()
	}
	if *testStruct.FloatAsFloatPtr != 3.14 {
		fmt.Println("Unable to unmarshal float into *float32!")
		t.Fail()
	}
}

func TestInts(t *testing.T) {
	// Make a test struct:
	testStruct := TestObject{}

	// Unmarshal:
	err := Unmarshal([]byte(jsonInts), &testStruct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Test ints:
	if testStruct.IntAsBool != 1 {
		fmt.Println("Unable to unmarshal bool into int32!")
		t.Fail()
	}
	if testStruct.IntAsFloat != 1 {
		fmt.Println("Unable to unmarshal float into int32!")
		t.Fail()
	}
	if testStruct.IntAsString != 1 {
		fmt.Println("Unable to unmarshal string into int32!")
		t.Fail()
	}
	if *testStruct.IntAsIntPtr != 1 {
		fmt.Println("Unable to unmarshal float into *int32!")
		t.Fail()
	}
}

func TestStrings(t *testing.T) {
	// Make a test struct:
	testStruct := TestObject{}

	// Unmarshal:
	err := Unmarshal([]byte(jsonStrings), &testStruct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Test strings:
	if testStruct.StringAsBool != "true" {
		fmt.Println("Unable to unmarshal bool into string!")
		t.Fail()
	}
	if testStruct.StringAsFloat != "3.14" {
		fmt.Println("Unable to unmarshal float into string!")
		t.Fail()
	}
	if testStruct.StringAsString != "one" {
		fmt.Println("Unable to unmarshal string into string!")
		t.Fail()
	}
	if *testStruct.StringAsStringPtr != "one" {
		fmt.Println("Unable to unmarshal string into *string!")
		t.Fail()
	}
}

func TestJsonFieldNames(t *testing.T) {
	// Make a test struct:
	testStruct := TestObject{}

	// Unmarshal:
	err := Unmarshal([]byte(jsonCapitalisations), &testStruct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Test strings:
	if testStruct.BoolWIthLowerCase != true {
		fmt.Println("Unable to unmarshal bool with assumed (camel-case) JSON field-name!")
		t.Fail()
	}
	if testStruct.FloatWithLowerCase != 37.75 {
		fmt.Println("Unable to unmarshal float with assumed (camel-case) JSON field-name!")
		t.Fail()
	}
	if testStruct.IntWithUpperCase != 123 {
		fmt.Println("Unable to unmarshal float with assumed (capitalised) JSON field-name!")
		t.Fail()
	}
	if testStruct.StringWithUpperCase != "99.9999" {
		fmt.Println("Unable to unmarshal string with assumed (capitalised) JSON field-name!")
		t.Fail()
	}
}
