package models

import (
	"reflect"
	"testing"
	"time"
)

// ... (unchanged)

func TestStockJsonTags(t *testing.T) {
	stock := Stock{}
	expectedJsonTags := map[string]interface{}{
		"ID":        "id",
		"TIMESTAMP": "TIMESTAMP",
		// ... add more fields as needed
		"ISIN": "ISIN",
	}

	// Check if the actual JSON tags match the expected JSON tags
	for field, expectedTag := range expectedJsonTags {
		actualTag := getFieldJsonTag(stock, field)
		if actualTag != expectedTag {
			t.Errorf("Unexpected JSON tag for field %v: got %v, want %v", field, actualTag, expectedTag)
		}
	}
}

// Helper function to get the JSON tag for a field in a struct
func getFieldJsonTag(s interface{}, field string) string {
	// Using reflection to get the JSON tag for the specified field
	fieldType := reflect.TypeOf(s)

	// Handle the error if the field is not found
	fieldStruct, found := fieldType.FieldByName(field)
	if !found {
		return ""
	}

	// Return the JSON tag if it exists, otherwise return an empty string
	return fieldStruct.Tag.Get("json")
}
