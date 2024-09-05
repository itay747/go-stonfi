package utils

import (
	"encoding/json"
	"net/url"
	"regexp"
)
var norm_req_re = regexp.MustCompile(`{([a-zA-Z0-9_]+)}`)
// normalizeRequest prepares the request URL and query parameters.
func NormalizeRequest(path string, options map[string][]string) (string, map[string][]string, error) {
	pathWithParams := norm_req_re.ReplaceAllStringFunc(path, func(m string) string {
		key := m[1 : len(m)-1] // Remove the curly braces
		values, exists := options[key]
		if !exists || len(values) == 0 {
			return m // Return the original match if no substitution is found
		}
		delete(options, key)              // Remove the key from query options after using it
		return url.QueryEscape(values[0]) // Escape the first value corresponding to the key
	})

	// Ensure all values in the options map are URL safe and keys are in snake_case
	for key, values := range options {
		for i, v := range values {
			values[i] = ToUrlSafe(v)
		}
		options[key] = values
	}
	options = DecamelizeKeys(options)

	return pathWithParams, options, nil
}

func NormalizeResponse(response []byte) ([]byte, error) {
	var data interface{}
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}

	data = CamelCaseKeys(data)
	data = DenullifyValues(data)

	return json.Marshal(data)
}
