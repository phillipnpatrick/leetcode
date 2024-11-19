package Solutions

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// Input: s = "ab-cd"
			// Output: "dc-ba"
			name: "test01",
			args: args{
				s: "ab-cd",
			},
			want: "dc-ba",
		},
		{
			// Input: s = "a-bC-dEf-ghIj"
			// Output: "j-Ih-gfE-dCba"
			name: "test02",
			args: args{
				s: "a-bC-dEf-ghIj",
			},
			want: "j-Ih-gfE-dCba",
		},
		{
			// Input: s = "Test1ng-Leet=code-Q!"
			// Output: "Qedo1ct-eeLg=ntse-T!"
			name: "test03",
			args: args{
				s: "Test1ng-Leet=code-Q!",
			},
			want: "Qedo1ct-eeLg=ntse-T!",
		},
		{
			name: "test04",
			args: args{
				s: "z<*zj",
			},
			want: "j<*zz",
		},
		{
			name: "test05",
			args: args{
				s: "-S2,_",
			},
			want: "-S2,_",
		},
		{
			name: "test06",
			args: args{
				s: "mv']4",
			},
			want: "vm']4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
