package code 

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) FetchData(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestService(t *testing.T) {
	repo := &RepositoryMock{}
	service := Service{repo: repo}
	repo.On("FetchData", "42").Return("data", nil).Once()

	err := service.Do()

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}