package service

import (
	"fmt"
	"testing"
)

// TestCheckGoogle_Error_123 tests the CheckGoogle method when GetGoogle returns an error.
func TestCheckGoogle_Error_123(t *testing.T) {
	// Arrange
	mockClient := new(MockHTTPClientInterface)
	mockClient.On("GetGoogle").Return("", fmt.Errorf("mock error"))

	svc := &Service{Cl: mockClient}

	// Act
	svc.CheckGoogle()

	// Assert
	mockClient.AssertExpectations(t)
}

// TestCheckGoogle_Success_456 tests the CheckGoogle method when GetGoogle returns a valid response.
func TestCheckGoogle_Success_456(t *testing.T) {
	// Arrange
	mockClient := new(MockHTTPClientInterface)
	mockClient.On("GetGoogle").Return("mock response", nil)

	svc := &Service{Cl: mockClient}

	// Act
	svc.CheckGoogle()

	// Assert
	mockClient.AssertExpectations(t)
}
