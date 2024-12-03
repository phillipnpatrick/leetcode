package Solutions

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: path = "/home/"
			// Output: "/home"
			name: "test01",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			// Input: path = "/home//foo/"
			// Output: "/home/foo"
			name: "test02",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			// Input: path = "/home/user/Documents/../Pictures"
			// Output: "/home/user/Pictures"
			name: "test03",
			args: args{
				path: "/home/user/Documents/../Pictures",
			},
			want: "/home/user/Pictures",
		},
		{
			// Input: path = "/../"
			// Output: "/"
			name: "test04",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			// Input: path = "/.../a/../b/c/../d/./"
			// Output: "/.../b/d"
			name: "test05",
			args: args{
				path: "/.../a/../b/c/../d/./",
			},
			want: "/.../b/d",
		},
		{
			name: "test06",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
		{
			name: "test07",
			args: args{
				path: `/a//b////c/d//././/..`,
			},
			want: "/a/b/c",
		},
		{
			name: "test08",
			args: args{
				path: `/..hidden`,
			},
			want: "/..hidden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
