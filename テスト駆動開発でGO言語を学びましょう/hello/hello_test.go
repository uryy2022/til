package hello

import (
	"fmt"
	"testing"
)

func Test_greetingPrefix(t *testing.T) {
	type args struct {
		language string
		name     string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix string
	}{
		{
			name: "正常系",
			args: args{
				language: "English",
				name:     "",
			},
			wantPrefix: "Hello,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPrefix := greetingPrefix(tt.args.language); gotPrefix != tt.wantPrefix {
				t.Errorf("greetingPrefix() = %v, want %v", gotPrefix, tt.wantPrefix)
			}
		})
	}
}
func ExampleHello() {
	fmt.Println(Hello("", ""))
	// Output: result="Hello, World"

}
