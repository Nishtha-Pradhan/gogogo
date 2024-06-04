package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue,omitempty"`
	NullIntValue    *int      `json:"nullIntValue"`
	EmptyString     string    `json:"emptyString,omitempty"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {

	fmt.Println("************** Generating JSON ****************\n")

	/**
	Generating JSON.
	*/

	// Using map to generate JSON.
	data := map[string]interface{}{
		"intValue": 1234,
		"boolValue": true,
		"stringValue": "hello!",
		"dateValue":   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"nullStringValue": nil,
		"nullIntValue":    nil,
	}


	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("This is the json data: %s\n", jsonData)

	// Using struct to generate JSON.
	otherInt := 4321
	structData := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}

	jsonStructData, structErr := json.Marshal(structData)
	if structErr != nil {
		fmt.Printf("Could not marshal struct json: %s\n", structErr)
		return
	}

	fmt.Printf("This is the struct json data: %s\n\n", jsonStructData)

	/**
	Parsing the generated JSON
	*/

	fmt.Println("************** Parsing JSON ****************\n")

	// Parsing JSON using map.
	var parsedMapJsonData map[string]interface{}
	mapErr := json.Unmarshal([]byte(jsonData), &parsedMapJsonData)
	// The jsonData variable is being passed to json.Unmarshal as a []byte because the function requires a []byte type and jsonData is initially defined as a string type. This works because a string in Go can be translated to a []byte, and vice versa.
	if mapErr != nil {
		fmt.Printf("Could not parse json: %s\n", mapErr)
		return
	}

	fmt.Printf("json map: %v\n", parsedMapJsonData)

	// Fetching the desired value from parsed Go data:
	rawDateValue, ok := parsedMapJsonData["dateValue"]
	if !ok {
		fmt.Printf("dateValue does not exist\n")
		return
	}
	mapDateValue, ok := rawDateValue.(string)

	if !ok {
		fmt.Printf("dateValue is not a string\n")
		return
	}
	fmt.Printf("Date Value: %s\n", mapDateValue)


	// Parsing JSON using struct
	var parsedStructJsonData myJSON
	structErr2 := json.Unmarshal([]byte(jsonStructData), &parsedStructJsonData)
	if structErr2 != nil {
		fmt.Printf("Could not unmarshal struct json: %s\n", structErr2)
		return
	}

	fmt.Printf("json struct: %#v\n", parsedStructJsonData)
	fmt.Printf("Date Value: %#v\n", parsedStructJsonData.DateValue)
	fmt.Printf("Object Value: %#v\n", parsedStructJsonData.ObjectValue)


}
