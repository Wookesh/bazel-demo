package greeter

import "testing"

func TestGreeter_Greet(t *testing.T) {
	type fields struct {
		who string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "user success",
			fields: fields{
				who: "user",
			},
			want: "hello user",
		},
		{
			// TODO: fix
			name: "empty success",
			fields: fields{
				who: "",
			},
			want: "hello ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeter{
				who: tt.fields.who,
			}
			if got := g.Greet(); got != tt.want {
				t.Errorf("Greet() returned: '%v', want: '%v'", got, tt.want)
			}
		})
	}
}
