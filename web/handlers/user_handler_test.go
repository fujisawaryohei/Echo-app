package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fujisawaryohei/blog-server/codes"
	mock_repositories "github.com/fujisawaryohei/blog-server/domain/users"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestUserList(t *testing.T) {
	storedUsers := &[]dto.User{
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
		name              string
		prepareListMock   func(mock *mock_repositories.MockUserRepository)
		prepareAuthMockFn func(mock *auth.MockIAuthenticator)
		users             *[]dto.User
		wantCode          int
	}{
		{
			name: "ユーザー一覧を返す",
			prepareListMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().List().Return(storedUsers, nil)
			},
			users:    storedUsers,
			wantCode: http.StatusOK,
		},
		{
			name: "internal server error",
			prepareListMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().List().Return(&[]dto.User{}, errors.New("internal server error"))
			},
			users:    &[]dto.User{},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareListMock(mr)
		authenticator := auth.NewAuthenticator()
		userUsecase := usecases.NewUserUsecase(mr, authenticator)
		UserHandler := NewUserHandler(userUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/admin/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserHandler.List(c)
		if rec.Code != tt.wantCode {
			t.Errorf("TestCase is %s", tt.name)
			t.Errorf("UserList() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestFindUser(t *testing.T) {
	storedUser := &dto.User{
		ID:                   1,
		Name:                 "test",
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	tests := []struct {
		name            string
		prepareFindMock func(mock *mock_repositories.MockUserRepository)
		user            *dto.User
		wantCode        int
	}{
		{
			name: "ユーザー情報を返す",
			prepareFindMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(storedUser, nil)
			},
			user:     storedUser,
			wantCode: http.StatusOK,
		},
		{
			name: "not found error",
			prepareFindMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&dto.User{}, codes.ErrUserNotFound)
			},
			user:     &dto.User{},
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareFindMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindById(gomock.Any()).Return(&dto.User{}, errors.New("internal server error"))
			},
			user:     &dto.User{},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareFindMock(mr)
		authenticator := auth.NewAuthenticator()
		userUsecase := usecases.NewUserUsecase(mr, authenticator)
		UserHandler := NewUserHandler(userUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/admin/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserHandler.Find(c)
		if rec.Code != tt.wantCode {
			t.Errorf("TestCase is %s", tt.name)
			t.Errorf("FindUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestStoreUser(t *testing.T) {
	storedUser := &dto.User{
		ID:                   1,
		Name:                 "test",
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	tests := []struct {
		name                   string
		prepareStoreMock       func(mock *mock_repositories.MockUserRepository)
		prepareFindByEmailMock func(mock *mock_repositories.MockUserRepository)
		prepareAuthMock        func(mock *auth.MockIAuthenticator)
		userJSON               string
		wantCode               int
	}{
		{
			name: "ユーザーを作成する",
			prepareStoreMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Save(gomock.Any()).Return(nil)
			},
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(nil, codes.ErrUserNotFound)
			},
			prepareAuthMock: func(mock *auth.MockIAuthenticator) {
				mock.EXPECT().GenerateToken(gomock.Any()).Return("secret", nil)
			},
			userJSON: `{"name": "test", "email":"test1@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusOK,
		},
		{
			name: "email has already existed",
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(storedUser, nil).AnyTimes()
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "internal server error",
			prepareStoreMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Save(gomock.Any()).Return(codes.ErrInternalServerError)
			},
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(nil, codes.ErrUserNotFound)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		var userUsecase *usecases.UserUseCase

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)
		ma := auth.NewMockIAuthenticator(ctrl)

		// メールアドレスの入力値検証の失敗のテストケースはFindByEmailのみをモックする
		if tt.name == "email has already existed" {
			tt.prepareFindByEmailMock(mr)
		} else {
			tt.prepareStoreMock(mr)
			tt.prepareFindByEmailMock(mr)
		}

		// ユーザーを作成するテストケースではAuthをモックする
		if tt.name == "ユーザーを作成する" {
			tt.prepareAuthMock(ma)
			userUsecase = usecases.NewUserUsecase(mr, ma)
		} else {
			authenticator := auth.NewAuthenticator()
			userUsecase = usecases.NewUserUsecase(mr, authenticator)
		}

		UserHandler := NewUserHandler(userUsecase)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/admin/users", strings.NewReader(tt.userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserHandler.Store(c)
		if rec.Code != tt.wantCode {
			t.Errorf("TestCase is %s", tt.name)
			t.Errorf("StoreUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestUpdate(t *testing.T) {
	storedUser := &dto.User{
		ID:                   1,
		Name:                 "test",
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	tests := []struct {
		name                   string
		prepareUpdateMock      func(mock *mock_repositories.MockUserRepository)
		prepareFindByEmailMock func(mock *mock_repositories.MockUserRepository)
		userJSON               string
		wantCode               int
	}{
		{
			name: "ユーザー情報を更新する",
			prepareUpdateMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(nil, codes.ErrUserNotFound)
			},
			userJSON: `{"name": "test", "email":"test1@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusAccepted,
		},
		{
			name: "email has already existed",
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(storedUser, nil).AnyTimes()
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "not found error",
			prepareUpdateMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(codes.ErrUserNotFound)
			},
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(nil, codes.ErrUserNotFound)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareUpdateMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(codes.ErrInternalServerError)
			},
			prepareFindByEmailMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().FindByEmail(gomock.Any()).Return(nil, codes.ErrUserNotFound)
			},
			userJSON: `{"name": "test", "email":"test@example.com", "password": "password", "password_confirmation": "password" }`,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mr := mock_repositories.NewMockUserRepository(ctrl)

		// メールアドレスの入力値検証の失敗のテストケースはFindByEmailのみをモックする
		if tt.name == "email has already existed" {
			tt.prepareFindByEmailMock(mr)
		} else {
			tt.prepareUpdateMock(mr)
			tt.prepareFindByEmailMock(mr)
		}

		authenticator := auth.NewAuthenticator()
		userUsecase := usecases.NewUserUsecase(mr, authenticator)
		UserHandler := NewUserHandler(userUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPatch, "/admin/users/:id", strings.NewReader(tt.userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserHandler.Update(c)
		if rec.Code != tt.wantCode {
			t.Errorf("TestCase is %s", tt.name)
			t.Errorf("UpdateUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name              string
		prepareDeleteMock func(mock *mock_repositories.MockUserRepository)
		wantCode          int
	}{
		{
			name: "ユーザーを削除する",
			prepareDeleteMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(nil)
			},
			wantCode: http.StatusAccepted,
		},
		{
			name: "not found error",
			prepareDeleteMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(codes.ErrUserNotFound)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "internal server error",
			prepareDeleteMock: func(mock *mock_repositories.MockUserRepository) {
				mock.EXPECT().Delete(gomock.Any()).Return(codes.ErrInternalServerError)
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		mr := mock_repositories.NewMockUserRepository(ctrl)
		tt.prepareDeleteMock(mr)
		authenticator := auth.NewAuthenticator()
		userUsecase := usecases.NewUserUsecase(mr, authenticator)
		UserHandler := NewUserHandler(userUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/admin/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		UserHandler.Delete(c)
		if rec.Code != tt.wantCode {
			t.Errorf("TestCase is %s", tt.name)
			t.Errorf("DeleteUser() code = %d, want = %d", rec.Code, tt.wantCode)
		}
	}
}
