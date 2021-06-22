package types

import "testing"

func TestRedisGetRefreshTokenFromAccessToken(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				key: "foofoofoofooFbarbarbarbar",
			},
			want: "foofoofoofoo",
		},
		{
			name: "empty",
			args: args{
				key: "",
			},
			want: "",
		},
		{
			name: "without F",
			args: args{
				key: "foofoofoofoo%barbarbarbar",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedisGetRefreshTokenFromAccessToken(tt.args.key); got != tt.want {
				t.Errorf("RedisGetRefreshTokenFromAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
