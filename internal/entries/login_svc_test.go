package entries

import (
	"github.com/stretchr/testify/assert"
	"pornrangers/entities"
	"pornrangers/mocks"
	"pornrangers/pkg/testutils"
	"testing"
)

func TestLoginSvc_Login(t *testing.T) {
	testutils.LoadEnv()
	mockRepo := mocks.NewLoginRepoInterface(t)
	mockUser := entities.UserLogin{Id: 1, Name: "Test User", Password: "$2a$10$eKRwYoRIcMcPTM/nDyB0tOF52HNPwjfqWcCa01L6qTawQApJrtfie"} // bcrypt hash for "password"
	mockRepo.EXPECT().
		GetUserByEmail("test@example.com").Return(mockUser, nil)

	loginSvc := NewLoginSvc(mockRepo)

	data, err := loginSvc.Login("test@example.com", "hello123")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", data.Email)
	assert.Equal(t, "Test User", data.Name)
	assert.NotEmpty(t, data.AccessToken)
}
