package pkg

func ToAnySlice[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func SetDataField(data *any, key string, value any) {
	var m map[string]any

	if existing, ok := (*data).(map[string]any); ok {
		m = existing
	} else {
		m = make(map[string]any)
		*data = m
	}

	m[key] = value
}
