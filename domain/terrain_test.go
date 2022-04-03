package domain

import (
	"reflect"
	"testing"
)

func TestTerrain_SpawnRover(t *testing.T) {
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Add a rover with Position 0 0",
			fields:  fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:    args{id: 1, pos: Position{0, 0, East}},
			wantErr: false,
		},
		{
			name:    "Add a rover with Position 5 5",
			fields:  fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:    args{id: 1, pos: Position{5, 5, North}},
			wantErr: true,
		},
		{
			name:    "Add a rover with Position -1 -1",
			fields:  fields{height: 5, width: 5, roverPos: make(map[int]Position)},
			args:    args{id: 1, pos: Position{-1, -1, South}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Terrain{
				height:   tt.fields.height,
				width:    tt.fields.width,
				roverPos: tt.fields.roverPos,
			}
			if err := p.SpawnRover(tt.args.id, tt.args.pos); (err != nil) != tt.wantErr {
				t.Errorf("Terrain.SpawnRover() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTerrain_MoveRover(t *testing.T) {
	type fields struct {
		height   int
		width    int
		roverPos map[int]Position
	}
	type args struct {
		id int
		c  Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Move rover that does not exist", fields{1, 1, make(map[int]Position)}, args{0, Command{}}, true},
		{"Move rover", fields{2, 2, map[int]Position{0: {0, 0, "north"}}}, args{0, Command{true, ""}}, false},
		{"Turn rover", fields{1, 1, map[int]Position{0: {0, 0, "north"}}}, args{0, Command{false, TurnLeft}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Terrain{
				height:   tt.fields.height,
				width:    tt.fields.width,
				roverPos: tt.fields.roverPos,
			}
			if err := p.MoveRover(tt.args.id, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Terrain.MoveRover() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTerrain_ValidatePos(t *testing.T) {
	type fields struct {
		height   int
		width    int
		roverPos map[int]Position
	}
	type args struct {
		pos Position
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Validate position on plateau", fields{2, 2, make(map[int]Position)}, args{Position{0, 0, North}}, false},
		{"Unvalidate position on plateau", fields{2, 2, make(map[int]Position)}, args{Position{0, 3, North}}, true},
		{"Unvalidate position on plateau", fields{2, 2, make(map[int]Position)}, args{Position{3, 0, North}}, true},
		{"Unvalidate position on plateau", fields{2, 2, make(map[int]Position)}, args{Position{0, 0, ""}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plat := &Terrain{
				height:   tt.fields.height,
				width:    tt.fields.width,
				roverPos: tt.fields.roverPos,
			}
			if err := plat.ValidatePos(tt.args.pos); (err != nil) != tt.wantErr {
				t.Errorf("Terrain.ValidatePos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTerrain(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	emptyMap := make(map[int]Position)
	tests := []struct {
		name    string
		args    args
		want    *Terrain
		wantErr bool
	}{
		{
			name:    "NewTerrain with 5 and 5",
			args:    args{5, 5},
			want:    &Terrain{5, 5, emptyMap},
			wantErr: false,
		},
		{
			name:    "NewTerrain with 0 and 5",
			args:    args{0, 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "NewTerrain with 5 and 0",
			args:    args{5, 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTerrain(tt.args.height, tt.args.width)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTerrain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerrain() = %v, want %v", got, tt.want)
			}
		})
	}
}
