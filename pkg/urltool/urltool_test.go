// @program:     ShotTener
// @file:        urltool_test.go.go
// @author:      16574
// @create:      2025-12-15 10:25
// @description:

package urltool

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		Url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{Url: "https://www.yuque.com/wenroudeyanhuo-goab3/"}, want: "wenroudeyanhuo-goab3", wantErr: false},
		{name: "无效示例", args: args{Url: "xxxx/123456"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.Url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
