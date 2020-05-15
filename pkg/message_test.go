package pkg

import (
	"reflect"
	"testing"
)

func Test_mergeStrMap(t *testing.T) {
	type args struct {
		a map[string]string
		b map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			"works with a longer than b",
			args{
				a: map[string]string{"traceID": "abc", "userID": "user-1", "foo": "bar"},
				b: map[string]string{"baz": "foo"},
			},
			map[string]string{"traceID": "abc", "userID": "user-1", "foo": "bar", "baz": "foo"},
		},
		{
			"works with b longer than a",
			args{
				a: map[string]string{"traceID": "abc"},
				b: map[string]string{"foo": "bar", "baz": "foo"},
			},
			map[string]string{"traceID": "abc", "foo": "bar", "baz": "foo"},
		},
		{
			"works with nil b",
			args{
				a: map[string]string{"traceID": "abc"},
				b: nil,
			},
			map[string]string{"traceID": "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStrMap(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeStrMap() = %v, want %v", got, tt.want)
			}
		})
	}
}