package utils

// Convert a map of languages into a simple map[string]string
func ExtractLanguages(languages map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for key, value := range languages {
		result[key] = value.(string)
	}
	return result
}

// Convert an array of interfaces into a string array
func ExtractStringArray(data []interface{}) []string {
	var result []string
	for _, value := range data {
		result = append(result, value.(string))
	}
	return result
}

// Extract the first element from a string array
func ExtractFirstString(data []interface{}) string {
	if len(data) > 0 {
		return data[0].(string)
	}
	return ""
}
