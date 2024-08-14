package entries

import (
	"github.com/stretchr/testify/assert"
	"pornrangers/mocks"
	"testing"
)

func TestRegisterSvc_CheckEmail(t *testing.T) {
	mockRepo := mocks.NewRegisterRepoInterface(t)
	mockRepo.EXPECT().
		CheckEmail("test@example.com").Return(false, nil)

	registerSvc := NewRegisterSvc(mockRepo)

	err := registerSvc.CheckEmail("test@example.com")
	assert.NoError(t, err)
}
