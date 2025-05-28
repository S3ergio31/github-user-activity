package domain

import (
	"encoding/json"
	"log"
)

func MapToStruct[T any](m map[string]any) T {
	var result T
	jsonBytes, err := json.Marshal(m)

	handleError(err)

	err = json.Unmarshal(jsonBytes, &result)

	handleError(err)

	return result
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("mapToStruct error: %v", err)
	}
}
