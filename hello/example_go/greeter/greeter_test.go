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
			name: "world success",
			fields: fields{
				who: "world",
			},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeter{
				who: tt.fields.who,
			}
			if got := g.Greet(); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
