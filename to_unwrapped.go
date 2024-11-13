// Copyright 2024 Pavllo Beo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package firestruct

import (
	"cloud.google.com/go/firestore"
	"errors"
	"fmt"
	"strings"
	"time"
)

// FirestoreFlatDataTypes lists Firestore protojson tags without nested data structures
var FirestoreFlatDataTypes = []string{
	strings.ToLower("stringValue"),
	strings.ToLower("booleanValue"),
	strings.ToLower("integerValue"),
	strings.ToLower("doubleValue"),
	strings.ToLower("nullValue"),
	strings.ToLower("bytesValue"),
	strings.ToLower("referenceValue"),
	strings.ToLower("geoPointValue"),
}

// UnwrapFirestoreFields unwraps a map[string]any containing Firestore protojson encoded fields
func UnwrapFirestoreFields(client *firestore.Client, input map[string]any) (map[string]any, error) {
	if input == nil {
		return nil, errors.New("firestruct: nil map contents")
	}

	output := make(map[string]any, len(input))

	for key, value := range input {
		unwrappedValue, err := unwrapValue(client, value)
		if err != nil {
			return nil, fmt.Errorf("firestruct: error processing field %q: %w", key, err)
		}
		output[key] = unwrappedValue
	}

	return output, nil
}

// unwrapValue handles any Firestore value type
func unwrapValue(client *firestore.Client, value any) (any, error) {
	valueMap, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("firestruct: expected map[string]any, got %T", value)
	}

	// Handle each possible value type
	for valueType, fieldValue := range valueMap {
		switch strings.ToLower(strings.TrimSpace(valueType)) {
		case "mapvalue":
			return unwrapMapValue(client, fieldValue)
		case "valuetype":
			return unwrapValue(client, fieldValue)
		case "arrayvalue":
			return unwrapArrayValue(client, fieldValue)
		case "referencevalue":
			if client != nil {
				return client.Doc(fieldValue.(string)), nil
			}
			return fieldValue, nil
		case "timestampvalue":
			if maptime, ok := fieldValue.(map[string]interface{}); ok {
				return mapToTimeString(client, maptime)
			} else {
				return fieldValue, nil
			}
		default:
			if containsString(FirestoreFlatDataTypes, strings.ToLower(strings.TrimSpace(valueType))) {
				return fieldValue, nil
			}
		}
	}

	return nil, fmt.Errorf("firestruct: no valid value type found")
}

func mapToTimeString(client *firestore.Client, data map[string]interface{}) (string, error) {
	// Extract seconds and nanoseconds from the map
	seconds, ok := data["seconds"].(float64)
	if !ok {
		seconds = 0
	}
	nanos, ok := data["nanos"].(float64)
	if !ok {
		nanos = 0
	}

	// Create a time.Time from Unix seconds and nanoseconds
	timestamp := time.Unix(int64(seconds), int64(nanos))

	// Format the time as a string
	formattedTime := timestamp.Format(time.RFC3339) // Change format as needed

	return formattedTime, nil
}

// unwrapMapValue handles map values specifically
func unwrapMapValue(client *firestore.Client, value any) (map[string]any, error) {
	mapValue, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("firestruct: expected map[string]any for mapValue, got %T", value)
	}

	// If empty map
	if len(mapValue) == 0 {
		return make(map[string]any), nil
	}

	// If has fields, process them
	if fields, ok := mapValue["fields"].(map[string]any); ok {
		return UnwrapFirestoreFields(client, fields)
	}

	// If direct map without fields
	return UnwrapFirestoreFields(client, mapValue)
}

// unwrapArrayValue handles array values specifically
func unwrapArrayValue(client *firestore.Client, value any) ([]any, error) {
	arrayValue, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("firestruct: expected map[string]any for arrayValue, got %T", value)
	}

	// Get the values array
	values, ok := arrayValue["values"].([]any)
	if !ok {
		// Empty array or invalid structure
		return make([]any, 0), nil
	}

	result := make([]any, len(values))
	for i, val := range values {
		unwrapped, err := unwrapValue(client, val)
		if err != nil {
			return nil, fmt.Errorf("firestruct: error processing array value at index %d: %w", i, err)
		}
		result[i] = unwrapped
	}

	return result, nil
}

// containsString checks if a string slice contains a specific string
func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
