package redis

import (
	"crypto/tls"
	"reflect"
	"testing"
	"time"

	"github.com/elvenworks/redis-conector/internal/delivery/actions"
	"github.com/go-redis/redis/v8"
)

func TestInitRedisCluster(t *testing.T) {
	type args struct {
		config InputClusterConfig
	}
	tests := []struct {
		name string
		args args
		want *RedisCluster
	}{
		{
			name: "Init No TLS",
			args: args{
				config: InputClusterConfig{
					Hosts:      []string{"localhost:6379"},
					MaxRetries: 1,
					Timeout:    time.Duration(1) * time.Second,
				},
			},
			want: &RedisCluster{
				ConfigCluster: &redis.ClusterOptions{
					Addrs:       []string{"localhost:6379"},
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
				},
				clusterActions: nil,
			},
		},
		{
			name: "Init TLS",
			args: args{
				config: InputClusterConfig{
					Hosts:      []string{"localhost:6379"},
					Password:   "1",
					MaxRetries: 1,
					Timeout:    time.Duration(1) * time.Second,
					TLS:        true,
				},
			},
			want: &RedisCluster{
				ConfigCluster: &redis.ClusterOptions{
					Addrs:       []string{"localhost:6379"},
					Password:    "1",
					MaxRetries:  1,
					DialTimeout: time.Duration(1) * time.Second,
					TLSConfig:   &tls.Config{},
				},
				clusterActions: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRedisCluster(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRedisCluster() = %v, want %v", got, tt.want)
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

			rc := &RedisCluster{
				ConfigCluster:  tt.fields.ConfigCluster,
				clusterActions: tt.fields.clusterActions,
			}
			if err := rc.ClusterSetMessage(tt.args.key, tt.args.message); (err != nil) != tt.wantErr {
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

			rc := &RedisCluster{
				ConfigCluster:  tt.fields.ConfigCluster,
				clusterActions: tt.fields.clusterActions,
			}
			got, err := rc.ClusterGetMessage(tt.args.key)
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
