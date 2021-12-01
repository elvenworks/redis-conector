package actions

import (
	"errors"
	"testing"

	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/go-redis/redismock/v8"
)

func Test_redisClient_SetMessage(t *testing.T) {

	type args struct {
		m domain.InputMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error Marshal",
			args: args{
				m: domain.InputMessage{
					Key:     "redis",
					Message: make(chan int),
				},
			},
			wantErr: true,
		},
		{
			name: "Error Diff Message",
			args: args{
				m: domain.InputMessage{
					Key:     "redis",
					Message: "redis",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client, mock := redismock.NewClientMock()

			mock.ExpectSet(tt.args.m.Key, tt.args.m.Message, 0).SetVal("")

			c := &redisClient{
				RedisClient: client,
			}
			if err := c.SetMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("redisClient.SetMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisClient_GetMessage(t *testing.T) {

	type args struct {
		m domain.InputMessage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get Message",
			args: args{
				m: domain.InputMessage{
					Key: "redis",
				},
			},
			want: "redis",
		},
		{
			name: "Erro",
			args: args{
				m: domain.InputMessage{
					Key: "redis",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client, mock := redismock.NewClientMock()

			if tt.wantErr != false {
				mock.ExpectGet(tt.args.m.Key).SetErr(errors.New("error"))
			}
			mock.ExpectGet(tt.args.m.Key).SetVal(tt.want)

			c := &redisClient{
				RedisClient: client,
			}
			got, err := c.GetMessage(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("redisClient.GetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("redisClient.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
