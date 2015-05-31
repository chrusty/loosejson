# loosejson
Loose JSON unmarshalling (for when your JSON arrives in the wrong type but you still want the data)

## Usage:
The "Unmarshal" method takes the same params as "json.Unmarshal()" (a []byte and a pointer to a struct/interface). Unmarshal() will attempt to convert the types of the JSON attributes into the correct types for your struct.
```
package main

import (
	"encoding/json"
	"fmt"
	loosejson ""
)

const (
	cruftJson = `{"eventType":"cruft","cruftPercent":"88.7","timestamp":"1432761305","isCrufty":true,"cruftiness":"87","description":"Brrrrrr blah blah cruft cruft..."}`
)

type Cruft struct {
	EventType        string  `protobuf:"bytes,1,req,name=eventType" json:"eventType,omitempty"`
	Timestamp        *string `protobuf:"bytes,2,req,name=timestamp" json:"timestamp,omitempty"`
	IsCrufty         *bool   `protobuf:"varint,3,opt,name=isCrufty" json:"isCrufty"`
	Cruftiness       int32   `protobuf:"varint,4,opt,name=cruftiness" json:"cruftiness,omitempty"`
	CruftPercent     float64 `protobuf:"varint,5,opt,name=cruftPercent" json:"cruftPercent,omitempty"`
	Description      string  `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func main() {
	// Make a "Cruft" event:
	cruft := Cruft{}

	// Unmarshal:
	err := loosejson.Unmarshal([]byte(cruftJson), &cruft)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Cruft.EventType:    %v\n", cruft.EventType)
	fmt.Printf("Cruft.Timestamp:    %v\n", cruft.Timestamp)
	fmt.Printf("Cruft.*Timestamp:   %v\n", *cruft.Timestamp)
	fmt.Printf("Cruft.IsCrufty:     %v\n", cruft.IsCrufty)
	fmt.Printf("Cruft.*IsCrufty:    %v\n", *cruft.IsCrufty)
	fmt.Printf("Cruft.Cruftiness:   %v\n", cruft.Cruftiness)
	fmt.Printf("Cruft.Description:  %v\n", cruft.Description)
	fmt.Printf("Cruft.CruftPercent: %v\n", cruft.CruftPercent)

}
```


## Cruft:
* Only accepts flat JSON (map[string]interface{}) - no nested JSON
* Only supports the following types:
  * int, int32, int64, *int, *int32, *int64
  * float32, float64, *float32, *float64
  * string, *string
  * bool, *bool

