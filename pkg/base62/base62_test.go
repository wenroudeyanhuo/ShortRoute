// @program:     ShotTener
// @file:        base62_test.go.go
// @author:      16574
// @create:      2025-12-15 12:32
// @description:

package base62

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt2String(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{seq: 6347}, want: "1En", wantErr: true},
		{name: "test2", args: args{seq: 62}, want: "10", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args.seq); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint64
		wantErr    bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{seq: "1En"}, wantResult: 6347, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := String2Int(tt.args.seq); gotResult != tt.wantResult {
				t.Errorf("String2Int() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
