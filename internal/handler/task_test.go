package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	entity "todo-app/internal/entity"
	"todo-app/internal/service"
	service_mocks "todo-app/internal/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Task(t *testing.T) {
	type args struct {
		ctx   context.Context
		input entity.InsertInput
	}

	type mockBehavior func(r *service_mocks.MockTasks, args args)

	tests := []struct {
		name                 string
		method               string
		path                 string
		args                 args
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "create new task",
			method: http.MethodPost,
			path:   "/tasks",
			args: args{
				ctx:   context.Background(),
				input: entity.InsertInput{Title: "title 1", Description: "some text"},
			},
			inputBody: `{"title":"title 1","description":"some text"}`,
			mockBehavior: func(r *service_mocks.MockTasks, args args) {
				r.EXPECT().Create(args.ctx, args.input).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"created id":1}`,
		},
		{
			name:                 "Bad request error",
			method:               http.MethodPost,
			path:                 "/tasks",
			args:                 args{},
			inputBody:            `"title":3456`,
			mockBehavior:         func(r *service_mocks.MockTasks, args args) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"json: cannot unmarshal string into Go value of type entity.InsertInput"}`,
		},
		{
			name:   "getting list of tasks",
			method: http.MethodGet,
			path:   "/tasks",
			args: args{
				ctx: context.Background(),
			},
			mockBehavior: func(r *service_mocks.MockTasks, args args) {
				r.EXPECT().GetAll(args.ctx).Return([]entity.Task{{Id: 1, Title: "title 1"}}, nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"title":"title 1","description":"","status":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task_server := service_mocks.NewMockTasks(c)
			test.mockBehavior(task_server, test.args)

			services := &service.Service{Tasks: task_server}
			handler := Handler{services}

			app := handler.InitRoutes()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, test.path, bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			app.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())

		})
	}
}
