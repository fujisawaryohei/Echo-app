package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
			wantCode: http.StatusOK,
		},
		{
			name: "not found error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&database.User{}, codes.ErrUserNotFound)
			},
			user:     &database.User{},
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&database.User{}, errors.New("internal server error"))
			},
			user:     &database.User{},
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
		req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		FindUser(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("FindUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestStoreUser(t *testing.T) {
	tests := []struct {
		name          string
		prepareMockFn func(mock *mock_repositories.MockUserRepository)
		userJSON      string
		wantCode      int
	}{
		{
			name: "ユーザーを作成する",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Save(gomock.Any()).Return(nil)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusCreated,
		},
		{
			name: "email has already existed",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Save(gomock.Any()).Return(codes.ErrUserEmailAlreadyExisted)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusConflict,
		},
		{
			name: "internal server error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Save(gomock.Any()).Return(codes.ErrInternalServerError)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
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
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tt.userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		StoreUser(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("StoreUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name          string
		prepareMockFn func(mock *mock_repositories.MockUserRepository)
		userJSON      string
		wantCode      int
	}{
		{
			name: "ユーザー情報を更新する",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusOK,
		},
		{
			name: "not found error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(codes.ErrUserNotFound)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(codes.ErrInternalServerError)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
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
		req := httptest.NewRequest(http.MethodPatch, "/users/:id", strings.NewReader(tt.userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UpdateUser(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("UpdateUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name          string
		prepareMockFn func(mock *mock_repositories.MockUserRepository)
		wantCode      int
	}{
		{
			name: "ユーザーを削除する",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "not found error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(codes.ErrUserNotFound)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareMockFn: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(codes.ErrInternalServerError)
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareMockFn(mr)
		userUsecase := usecases.NewUserUsecase(mr)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		DeleteUser(userUsecase)(c)
		if rec.Code != tt.wantCode {
			t.Errorf("DeleteUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}
