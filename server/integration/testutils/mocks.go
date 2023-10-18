package testutils

import "github.com/curio-research/keystone/server"

type MockErrorHandler struct {
	errorsByPlayerID []string
}

func NewMockErrorHandler() *MockErrorHandler {
	return &MockErrorHandler{errorsByPlayerID: []string{}}
}

func (m *MockErrorHandler) FormatMessage(id int, errorMessage string) *server.NetworkMessage {
	m.errorsByPlayerID = append(m.errorsByPlayerID, errorMessage)
	return &server.NetworkMessage{}
}

func (m *MockErrorHandler) LastError() string {
	return m.errorsByPlayerID[len(m.errorsByPlayerID)-1]
}

func (m *MockErrorHandler) ErrorCount() int {
	return len(m.errorsByPlayerID)
}
