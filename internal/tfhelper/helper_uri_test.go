package tfhelper

import (
	"testing"
)

func TestResolveURI(t *testing.T) {
	t.Parallel()

	type args struct {
		uri  string
		map1 map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				uri: "/api/resource/{name}/{id}",
				map1: map[string]interface{}{
					"name": "account1",
					"id":   "zouzou",
				},
			},
			want: "/api/resource/account1/zouzou",
		},
		{
			name: "avoid bool",
			args: args{
				uri: "/api/resource/{name}",
				map1: map[string]interface{}{
					"name": "account1",
					"id":   true,
				},
			},
			want: "/api/resource/account1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveURI(tt.args.uri, tt.args.map1); got != tt.want {
				t.Errorf("ResolveURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
