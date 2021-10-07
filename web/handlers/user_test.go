package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/database"
	mock_repositories "github.com/fujisawaryohei/blog-server/domain/mock-repositories"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestUserList(t *testing.T) {
	storedUsers := &[]database.User{
		{
			ID:                   1,
			Name:                 "test",
			Email:                "test@example.com",
			Password:             "password",
			PasswordConfirmation: "password",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			ID:                   2,
			Name:                 "test2",
			Email:                "test2@example.com",
			Password:             "password",
			PasswordConfirmation: "password",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
	}
	tests := []struct {
		name          string
		prepareMockFn func(mock *mock_repositories.MockUserRepository)
		users         *[]database.User
		wantCode      int
	}{
		{
			name: "ユーザー一覧を返す",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().List().Return(storedUsers, nil)
			},
			users:    storedUsers,
			wantCode: http.StatusOK,
		},
		{
			name: "internal server error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().List().Return(&[]database.User{}, errors.New("internal server error"))
			},
			users:    &[]database.User{},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareMockFn(mr)
		userUsecase := usecases.NewUserUsecase(mr)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserList(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("UserList() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestFindUser(t *testing.T) {
	storedUser := &database.User{
		ID:                   1,
		Name:                 "test",
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	tests := []struct {
		name          string
		prepareMockFn func(mock *mock_repositories.MockUserRepository)
		user          *database.User
		wantCode      int
	}{
		{
			name: "ユーザー情報を返す",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(storedUser, nil)
			},
			user:     storedUser,
			wantCode: 200,
		},
		{
			name: "Not Found Error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&database.User{}, codes.ErrUserNotFound)
			},
			user:     &database.User{},
			wantCode: 404,
		},
		{
			name: "Internal Server Error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&database.User{}, errors.New("internal server error"))
			},
			user:     &database.User{},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareMockFn(mr)
		userUsecase := usecases.NewUserUsecase(mr)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		FindUser(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("FindUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}
