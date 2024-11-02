package parser

import (
	// "fmt"
	"testing"
)

// Define a helper function to create test cases
func createTestRequest(name string, method string, body string) YamlRequest {
	return YamlRequest{
		Name:       name,
		Method:     method,
		RequestBody: body,
	}
}

// Test validateRequests function
func TestValidateRequests(t *testing.T) {
	tests := []struct {
		name      string
		parseData YamlParserConfig
		expectErr bool
	}{
		{
			name: "No Requests",
			parseData: YamlParserConfig{
				Requests: []YamlRequest{},
			},
			expectErr: true,
		},
		{
			name: "Valid Requests",
			parseData: YamlParserConfig{
				Requests: []YamlRequest{
					createTestRequest("Request1", "GET", ""),
				},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequests(tt.parseData)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateRequests() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test requestsExists function
func TestRequestsExists(t *testing.T) {
	tests := []struct {
		name      string
		parseData YamlParserConfig
		expected  bool
	}{
		{
			name: "Empty Requests",
			parseData: YamlParserConfig{
				Requests: []YamlRequest{},
			},
			expected: false,
		},
		{
			name: "Non-empty Requests",
			parseData: YamlParserConfig{
				Requests: []YamlRequest{
					createTestRequest("Request1", "POST", "Body"),
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := requestsExists(tt.parseData)
			if result != tt.expected {
				t.Errorf("requestsExists() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// Test validateRequest function
func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		request     YamlRequest
		requestIndex int
		expectErr   bool
	}{
		{
			name:        "Missing Name",
			request:     createTestRequest("", "POST", "Body"),
			requestIndex: 1,
			expectErr:   true,
		},
		{
			name:        "Missing Method",
			request:     createTestRequest("Request1", "", "Body"),
			requestIndex: 1,
			expectErr:   true,
		},
		{
			name:        "Valid Request",
			request:     createTestRequest("Request1", "GET", ""),
			requestIndex: 1,
			expectErr:   false,
		},
		{
			name:        "Missing Request Body for POST",
			request:     createTestRequest("Request2", "POST", ""),
			requestIndex: 2,
			expectErr:   true,
		},
		{
			name:        "Valid Request with Body",
			request:     createTestRequest("Request3", "POST", "Body"),
			requestIndex: 3,
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequest(tt.requestIndex, tt.request)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateRequest() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test validateRequestMethod function
func TestValidateRequestMethod(t *testing.T) {
	tests := []struct {
		name        string
		requestMethod string
		requestIndex int
		expectErr   bool
	}{
		{
			name:        "Valid Method GET",
			requestMethod: "GET",
			requestIndex: 1,
			expectErr:   false,
		},
		{
			name:        "Invalid Method",
			requestMethod: "INVALID",
			requestIndex: 2,
			expectErr:   true,
		},
		{
			name:        "Valid Method POST",
			requestMethod: "POST",
			requestIndex: 3,
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequestMethod(tt.requestIndex, tt.requestMethod)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateRequestMethod() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
