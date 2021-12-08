package actions

import (
	"errors"
	"testing"

	"github.com/elvenworks/redis-conector/internal/domain"
	"github.com/go-redis/redismock/v8"
)

func Test_clusterClient_ClusterSetMessage(t *testing.T) {

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

			client, mock := redismock.NewClusterMock()

			mock.ExpectSet(tt.args.m.Key, tt.args.m.Message, 0).SetVal("")

			c := &clusterClient{
				RedisClusterClient: client,
			}
			if err := c.ClusterSetMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("clusterClient.ClusterSetMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clusterClient_ClusterGetMessage(t *testing.T) {
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

			client, mock := redismock.NewClusterMock()

			if tt.wantErr != false {
				mock.ExpectGet(tt.args.m.Key).SetErr(errors.New("error"))
			}

			mock.ExpectGet(tt.args.m.Key).SetVal(tt.want)

			c := &clusterClient{
				RedisClusterClient: client,
			}
			got, err := c.ClusterGetMessage(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("clusterClient.ClusterGetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("clusterClient.ClusterGetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
