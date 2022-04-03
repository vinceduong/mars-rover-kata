package domain

import (
	"reflect"
	"testing"
)

func TestNewPlateau(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	emptyMap := make(map[int]Position)
	tests := []struct {
		name      string
		args      args
		want      *Plateau
		wantPanic bool
	}{
		{
			name:      "NewPlateau with 5 and 5",
			args:      args{5, 5},
			want:      &Plateau{5, 5, emptyMap},
			wantPanic: false,
		},
		{
			name:      "NewPlateau with 0 and 5",
			args:      args{0, 5},
			want:      &Plateau{5, 5, emptyMap},
			wantPanic: true,
		},
		{
			name:      "NewPlateau with 5 and 0",
			args:      args{5, 0},
			want:      &Plateau{5, 5, emptyMap},
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

func TestPlateau_SpawnRover(t *testing.T) {
	type fields struct {
		height   int
		width    int
		roverPos map[int]Position
	}
	type args struct {
		id  int
		pos Position
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPanic bool
	}{
		{
			name:      "Add a rover with Position 0 0",
			fields:    fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:      args{id: 1, pos: Position{0, 0, East}},
			wantPanic: false,
		},
		{
			name:      "Add a rover with Position 5 5",
			fields:    fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:      args{id: 1, pos: Position{5, 5, North}},
			wantPanic: true,
		},
		{
			name:      "Add a rover with Position -1 -1",
			fields:    fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:      args{id: 1, pos: Position{-1, -1, South}},
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
			p := &Plateau{
				height:   tt.fields.height,
				width:    tt.fields.width,
				roverPos: tt.fields.roverPos,
			}
			p.SpawnRover(tt.args.id, tt.args.pos)
		})
	}
}
