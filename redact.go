package sensitive_logging

import (
	"encoding/json"
)

const (
	redactedValue = "*****"
)

func redactedSensitiveData(sensitiveFields []string, data any) any {
	mapV, err := convertTo[map[string]any](data)
	if err == nil {
		redactedSensitiveDataFromMap(sensitiveFields, mapV)
		return mapV
	}

	sliceV, err := convertTo[[]any](data)
	if err == nil {
		for i, v := range sliceV {
			sliceV[i] = redactedSensitiveData(sensitiveFields, v)
		}
		return sliceV
	}

	return data
}

func redactedSensitiveDataFromMap(sensitiveFields []string, vMap map[string]any) {
	for k, v := range vMap {
		if contains(sensitiveFields, k) {
			vMap[k] = redactedValue
		}
		if nv, ok := v.(map[string]any); ok {
			redactedSensitiveDataFromMap(sensitiveFields, nv)
		}
		if nv, ok := v.([]any); ok {
			for i, item := range nv {
				nv[i] = redactedSensitiveData(sensitiveFields, item)
			}
		}
	}
}

func contains[T comparable](arr []T, item T) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

func convertTo[T any](input any) (T, error) {
	var emptyT T
	if input == nil {
		return emptyT, nil
	}
	if v, ok := input.(T); ok {
		return v, nil
	}

	bytes, err := json.Marshal(input)
	if err != nil {
		return emptyT, nil
	}
	var v T
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return emptyT, err
	}
	return v, nil
}
