package redis

import (
	"crypto/tls"
	"reflect"
	"testing"
	"time"

	"github.com/elvenworks/redis-conector/internal/delivery/actions"
	"github.com/go-redis/redis/v8"
)

func TestInitRedis(t *testing.T) {

	type args struct {
		config InputConfig
	}
	tests := []struct {
		name string
		args args
		want *Redis
	}{
		{
			name: "Init No TLS",
			args: args{
				config: InputConfig{
					Host:       "localhost:6379",
					DB:         1,
					MaxRetries: 1,
					Timeout:    time.Duration(1) * time.Second,
				},
			},
			want: &Redis{
				Config: &redis.Options{
					Addr:        "localhost:6379",
					DB:          1,
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
				},
				ConfigCluster: &redis.ClusterOptions{
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
				},
				actions:        nil,
				clusterActions: nil,
			},
		},
		{
			name: "Init TLS",
			args: args{
				config: InputConfig{
					Host:       "localhost:6379",
					Password:   "1",
					DB:         1,
					MaxRetries: 1,
					Timeout:    time.Duration(1) * time.Second,
				},
			},
			want: &Redis{
				Config: &redis.Options{
					Addr:        "localhost:6379",
					Password:    "1",
					DB:          1,
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
					TLSConfig:   &tls.Config{},
				},
				ConfigCluster: &redis.ClusterOptions{
					Password:    "1",
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
					TLSConfig:   &tls.Config{},
				},
				actions:        nil,
				clusterActions: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRedis(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_SetMessage(t *testing.T) {
	type fields struct {
		Config  *redis.Options
		actions actions.MockRedisClient
	}
	type args struct {
		key     string
		message interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				key:     "redis",
				message: "redis",
			},
			fields: fields{
				Config:  &redis.Options{},
				actions: actions.MockRedisClient{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.actions.On("SetMessage", tt.fields.Config).Return(tt.wantErr)

			r := &Redis{
				Config:  tt.fields.Config,
				actions: tt.fields.actions,
			}
			if err := r.SetMessage(tt.args.key, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Redis.SetMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_GetMessage(t *testing.T) {
	type fields struct {
		Config  *redis.Options
		actions actions.MockRedisClient
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				Config:  &redis.Options{},
				actions: actions.MockRedisClient{},
			},
			args: args{
				key: "redis",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.actions.On("GetMessage", tt.fields.Config).Return(tt.wantErr)

			r := &Redis{
				Config:  tt.fields.Config,
				actions: tt.fields.actions,
			}
			got, err := r.GetMessage(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.GetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Redis.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedis_ClusterSetMessage(t *testing.T) {
	type fields struct {
		ConfigCluster  *redis.ClusterOptions
		clusterActions actions.MockClusterClient
	}
	type args struct {
		key     string
		message interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				key:     "redis",
				message: "redis",
			},
			fields: fields{
				ConfigCluster:  &redis.ClusterOptions{},
				clusterActions: actions.MockClusterClient{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.clusterActions.On("SetMessage", tt.fields.ConfigCluster).Return(tt.wantErr)

			r := &Redis{
				ConfigCluster:  tt.fields.ConfigCluster,
				clusterActions: tt.fields.clusterActions,
			}
			if err := r.ClusterSetMessage(tt.args.key, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Redis.ClusterSetMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_ClusterGetMessage(t *testing.T) {
	type fields struct {
		ConfigCluster  *redis.ClusterOptions
		clusterActions actions.MockClusterClient
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test",
			fields: fields{
				ConfigCluster:  &redis.ClusterOptions{},
				clusterActions: actions.MockClusterClient{},
			},
			args: args{
				key: "redis",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.clusterActions.On("GetMessage", tt.fields.ConfigCluster).Return(tt.wantErr)

			r := &Redis{
				ConfigCluster:  tt.fields.ConfigCluster,
				clusterActions: tt.fields.clusterActions,
			}
			got, err := r.ClusterGetMessage(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.ClusterGetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Redis.ClusterGetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
