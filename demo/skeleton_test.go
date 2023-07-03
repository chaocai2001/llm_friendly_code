package demo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) RetiveData(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) StoreData(token string, data string) error {
	args := m.Called(token, data)
	return args.Error(0)
}

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(raw string) (string, error) {
	args := m.Called(raw)
	return args.String(0), args.Error(1)
}

type MockTokenCreator struct {
	mock.Mock
}

func (m *MockTokenCreator) CreateToken(data string) string {
	args := m.Called(data)
	return args.String(0)
}

func TestProcessingServiceImpl_Process(t *testing.T) {
	mockStorage := new(MockStorage)
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)

	service := NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	raw := "raw_data"
	processed := "processed_data"
	token := "token"

	mockProcessor.On("Process", raw).Return(processed, nil).Once()
	mockTokenCreator.On("CreateToken", processed).Return(token)
	mockStorage.On("StoreData", token, processed).Return(nil)

	result, err := service.Process(raw)

	assert.NoError(t, err)
	assert.Equal(t, token, result)

	mockProcessor.AssertExpectations(t)
	mockTokenCreator.AssertExpectations(t)
	mockStorage.AssertExpectations(t)

	// Test error case
	mockProcessor.On("Process", raw).Return("", errors.New("processing error")).Once()
	_, err = service.Process(raw)
	assert.Error(t, err)
}

func TestProcessingServiceImpl_Retrive(t *testing.T) {
	mockStorage := new(MockStorage)
	mockProcessor := new(MockProcessor)
	mockTokenCreator := new(MockTokenCreator)

	service := NewProcessingService(mockProcessor, mockTokenCreator, mockStorage)

	token := "token"
	data := "data"

	mockStorage.On("RetiveData", token).Return(data, nil).Once()

	result, err := service.Retrive(token)

	assert.NoError(t, err)
	assert.Equal(t, data, result)

	mockStorage.AssertExpectations(t)

	// Test error case
	mockStorage.On("RetiveData", token).Return("", errors.New("retrieval error")).Once()
	_, err = service.Retrive(token)
	assert.Error(t, err)
}
