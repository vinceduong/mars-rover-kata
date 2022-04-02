package main

import (
	"reflect"
	"testing"
)

func TestNewPlateau(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name      string
		args      args
		want      *Plateau
		wantPanic bool
	}{
		{
			name:      "NewPlateau with 5 and 5",
			args:      args{5, 5},
			want:      &Plateau{5, 5},
			wantPanic: false,
		},
		{
			name:      "NewPlateau with 5 and 5",
			args:      args{0, 5},
			want:      &Plateau{5, 5},
			wantPanic: true,
		},
		{
			name:      "NewPlateau with 5 and 5",
			args:      args{5, 0},
			want:      &Plateau{5, 5},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SequenceInt() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := NewPlateau(tt.args.height, tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlateau() = %v, want %v", got, tt.want)
			}
		})
	}
}
