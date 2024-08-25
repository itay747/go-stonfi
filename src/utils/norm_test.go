package utils

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNormalizeRequest(t *testing.T) {
	tests := []struct {
		queryParams   url.Values
		expectedQuery url.Values
		name          string
		path          string
		expectedPath  string
		expectError   bool
	}{
		{
			name:          "should insert path parameters and remove them from query",
			path:          "/v1/{foo}",
			queryParams:   url.Values{"foo": []string{"bar"}, "baz": []string{"qux"}},
			expectedPath:  "/v1/bar",
			expectedQuery: url.Values{"baz": []string{"qux"}},
		},
		{
			name:        "should throw an error if path parameter is missing in query",
			path:        "/v1/{foo}",
			queryParams: url.Values{"baz": []string{"qux"}},
			expectError: true,
		},
		{
			name:          "should transform path parameters to URL-safe strings",
			path:          "/v1/{foo}",
			queryParams:   url.Values{"foo": []string{"bar/A+B=C"}},
			expectedPath:  "/v1/bar%2FA%2BB%3DC",
			expectedQuery: url.Values{},
		},
		{
			name:          "should transform query parameters to URL-safe strings",
			path:          "/v1/{foo}",
			queryParams:   url.Values{"foo": []string{"bar"}, "baz": []string{"qux/A+B=C"}},
			expectedPath:  "/v1/bar",
			expectedQuery: url.Values{"baz": []string{"qux%2FA%2BB%3DC"}},
		},
		{
			name:          "should decamelize data in query parameters",
			path:          "/v1/{foo}",
			queryParams:   url.Values{"foo": []string{"0x123"}, "barBaz": []string{"qux"}},
			expectedPath:  "/v1/0x123",
			expectedQuery: url.Values{"bar_baz": []string{"qux"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultPath, resultQuery, err := NormalizeRequest(tt.path, tt.queryParams)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError {
				if resultPath != tt.expectedPath {
					t.Errorf("Expected path %s, got %s", tt.expectedPath, resultPath)
				}
				if url.Values(resultQuery).Encode() != tt.expectedQuery.Encode() {
					t.Errorf("Expected query %s, got %s", tt.expectedQuery.Encode(), url.Values(resultQuery).Encode())
				}
			}
		})
	}
}

func TestNormalizeResponse(t *testing.T) {
	t.Run("should camelcase data in response", func(t *testing.T) {
		input := map[string]interface{}{
			"foo_bar": "baz",
			"bar_baz": "qux",
			"child": map[string]interface{}{
				"foo_bar": "baz",
				"bar_baz": "qux",
			},
		}
		expected := map[string]interface{}{
			"fooBar": "baz",
			"barBaz": "qux",
			"child": map[string]interface{}{
				"fooBar": "baz",
				"barBaz": "qux",
			},
		}
		result := CamelCaseKeys(DenullifyValues(input))
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("should replace all potential null values with default empty strings", func(t *testing.T) {
		input := map[string]interface{}{
			"foo": nil,
			"child": map[string]interface{}{
				"foo": nil,
			},
		}
		expected := map[string]interface{}{
			"foo": "",
			"child": map[string]interface{}{
				"foo": "",
			},
		}
		result := CamelCaseKeys(DenullifyValues(input))
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
