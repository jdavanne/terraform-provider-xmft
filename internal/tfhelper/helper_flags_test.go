package tfhelper

import (
	"strings"
	"testing"
)

func Test_checkSupportedAttributes(t *testing.T) {
	t.Parallel()
	type args struct {
		flags string
		err   string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "ok-one", args: args{flags: ",optional", err: ""}},
		{name: "ok-multiple", args: args{flags: ",optional,default", err: ""}},
		{name: "ok-with-name", args: args{flags: "anyname,optional,default:", err: ""}},
		{name: "ok-with-bool-value", args: args{flags: ",optional,default:true", err: ""}},
		{name: "ok-with-int-value", args: args{flags: ",optional,default:1234", err: ""}},
		{name: "ok-with-empty-value", args: args{flags: ",optional,default:", err: ""}},

		{name: "ko-name-single-flag", args: args{flags: "required", err: "conflicting name with helper flag"}},
		{name: "ko-name-single-attr", args: args{flags: "default", err: "conflicting name with helper flag"}},
		{name: "ko-name-single-empty-attr", args: args{flags: "default:", err: "conflicting name with helper flag"}},
		{name: "ko-name-single-bool-attr", args: args{flags: "default:false", err: "conflicting name with helper flag"}},
		{name: "ko-name-single-int-attr", args: args{flags: "default:1", err: "conflicting name with helper flag"}},
		{name: "ko-name-multiple", args: args{flags: "required,optional", err: "conflicting name with helper flag"}},
		{name: "ko-name-multiple-empty-attr", args: args{flags: "default:,optional", err: "conflicting name with helper flag"}},
		{name: "ko-name-multiple-int-attr", args: args{flags: "default:1234", err: "conflicting name with helper flag"}},
		{name: "ko-name-multiple-bool-attr", args: args{flags: "default:true,optional", err: "conflicting name with helper flag"}},

		{name: "ko-unknown-flag", args: args{flags: ",optional,zouzou", err: "unsupported helper flag"}},
		{name: "ko-unknown-flag", args: args{flags: ",optional,default1234", err: "unsupported helper flag"}},

		{name: "ko-unknown-flag-with-name-1", args: args{flags: "anyname,zou", err: "unsupported helper flag"}},
		{name: "ko-unknown-flag-with-name-2", args: args{flags: "anyname,optional,zou", err: "unsupported helper flag"}},
		{name: "ko-unknown-flag-with-name-3", args: args{flags: "anyname,optional,required,zou", err: "unsupported helper flag"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := checkSupportedAttributes("any-location", tt.args.flags)
			if (val == nil && tt.args.err != "") || (val != nil && !strings.Contains(val.Error(), tt.args.err)) {
				t.Errorf("checkSupportedAttributes() = %v, want %v", val, tt.args.err)
			}
		})
	}
}

func TestFlagsGet(t *testing.T) {
	t.Parallel()

	type args struct {
		flags string
		flag  string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{name: "ok-one", args: args{flags: ",default", flag: "default"}, want: "", want1: true},
		{name: "ok-multiple-last-flag", args: args{flags: ",optional,default", flag: "default"}, want: "", want1: true},
		{name: "ok-multiple-last-empty", args: args{flags: ",optional,default:", flag: "default"}, want: "", want1: true},
		{name: "ok-multiple-last-int", args: args{flags: ",optional,default:1234", flag: "default"}, want: "1234", want1: true},
		{name: "ok-multiple-last-bool-true", args: args{flags: ",optional,default:true", flag: "default"}, want: "true", want1: true},
		{name: "ok-multiple-last-bool-false", args: args{flags: ",optional,default:false", flag: "default"}, want: "false", want1: true},
		{name: "ok-multiple-last-string", args: args{flags: ",optional,default:zou", flag: "default"}, want: "zou", want1: true},

		{name: "ok-multiple-middle-flag", args: args{flags: ",optional,default,computed", flag: "default"}, want: "", want1: true},
		{name: "ok-multiple-middle-empty", args: args{flags: ",optional,default:,computed", flag: "default"}, want: "", want1: true},
		{name: "ok-multiple-middle-int", args: args{flags: ",optional,default:1234,computed", flag: "default"}, want: "1234", want1: true},
		{name: "ok-multiple-middle-bool-true", args: args{flags: ",optional,default:true,computed", flag: "default"}, want: "true", want1: true},
		{name: "ok-multiple-middle-bool-false", args: args{flags: ",optional,default:false,computed", flag: "default"}, want: "false", want1: true},
		{name: "ok-multiple-middle-string", args: args{flags: ",optional,default:zou,computed", flag: "default"}, want: "zou", want1: true},

		{name: "ko-unknown-flag", args: args{flags: ",optional", flag: "zouzou"}, want: "", want1: false},
		{name: "ko-unknown-flag", args: args{flags: ",optional,default", flag: "zouzou"}, want: "", want1: false},
		{name: "ko-unknown-flag-with-name-1", args: args{flags: "anyname,zou", flag: "zouzou"}, want: "", want1: false},
		{name: "ko-unknown-flag-with-name-2", args: args{flags: "anyname,optional,zou", flag: "zouzou"}, want: "", want1: false},
		{name: "ko-unknown-flag-with-name-3", args: args{flags: "anyname,optional,required,zou", flag: "zouzou"}, want: "", want1: false},
		{name: "ko-unknown-flag-with-name-2", args: args{flags: "anyname,optional,zou:", flag: "zouzou"}, want: "", want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FlagsGet(tt.args.flags, tt.args.flag)
			if got != tt.want {
				t.Errorf("FlagsGet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FlagsGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
