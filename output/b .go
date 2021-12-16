// Code generated by schema-generate. DO NOT EDIT.

package output

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// AdditionalProperties 
type AdditionalProperties struct {
  Property1 string `json:"property1,omitempty"`
  Property2 *Address `json:"property2,omitempty"`
  Property3 map[string]int `json:"property3,omitempty"`
  Property4 map[string]string `json:"property4,omitempty"`
  Property5 map[string]*NotSoAnonymous `json:"property5,omitempty"`
  Property6 map[string]*Property6Item `json:"property6,omitempty"`
  Property7 map[string]map[string]*Anonymous1 `json:"property7,omitempty"`
}

// Address 
type Address struct {
  City string `json:"city,omitempty"`
}

// Anonymous1 
type Anonymous1 struct {
  Subproperty1 int `json:"subproperty1,omitempty"`
}

// NotSoAnonymous 
type NotSoAnonymous struct {
  Subproperty1 int `json:"subproperty1,omitempty"`
}

// PackageList List of packages marked as abandoned for this repository, the mark can be boolean or a package name/URL pointing to a recommended alternative.
type PackageList struct {
}

// Property6Item 
type Property6Item struct {
  Subproperty1 int `json:"subproperty1,omitempty"`
}

// Root 
type Root struct {

  // List of packages marked as abandoned for this repository, the mark can be boolean or a package name/URL pointing to a recommended alternative.
  Abandoned *PackageList `json:"abandoned,omitempty"`

  // Repository name.
  Name string `json:"name"`
}

func (strct *Root) BMarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "abandoned" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"abandoned\": ")
	if tmp, err := json.Marshal(strct.Abandoned); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Root) BUnmarshalJSON(b []byte) error {
    nameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "abandoned":
            if err := json.Unmarshal([]byte(v), &strct.Abandoned); err != nil {
                return err
             }
        case "name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            nameReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if name (a required property) was received
    if !nameReceived {
        return errors.New("\"name\" is required but was not present")
    }
    return nil
}
