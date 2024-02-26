package types

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_TreeCmd(t *testing.T) {

	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}

	type args struct {
		pair    Pair
		command Command
	}

	var ttt, v = "name", "password"

	c, e := NewClient(&HostURL, &ttt, &v)

	if e != nil {
		return
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pair
		wantErr bool
	}{
		struct {
			name    string
			fields  fields
			args    args
			want    *Pair
			wantErr bool
		}{
			name: "First test",
			args: args{
				pair: Pair{
					V: "Test",
				},
				command: Command{
					Name: []string{"tree"},
					Flags: []KV{
						KV{
							K: "value",
							V: "test",
						},
						KV{
							K: "action",
							V: "1",
						},
						KV{
							K: "id",
							V: "0",
						},
						KV{
							K: "output",
							V: "byte",
						},
					},
				},
			},
			want: &Pair{
				V: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.TreeCmd(tt.args.command)

			if (err != nil) != tt.wantErr {

				t.Errorf("Client.TreeCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.TreeCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
