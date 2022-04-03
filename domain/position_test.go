package domain

import (
	"reflect"
	"testing"
)

func TestDirection_TurnTo(t *testing.T) {
	type args struct {
		r Rotation
	}
	tests := []struct {
		name string
		d    Direction
		args args
		want Direction
	}{
		{"Turn Left as North", North, args{TurnLeft}, West},
		{"Turn Left as West", West, args{TurnLeft}, South},
		{"Turn Left as South", South, args{TurnLeft}, East},
		{"Turn Left as East", East, args{TurnLeft}, North},
		{"Turn Right as North", North, args{TurnRight}, East},
		{"Turn Right as West", West, args{TurnRight}, North},
		{"Turn Right as South", South, args{TurnRight}, West},
		{"Turn Right as East", East, args{TurnRight}, South},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.TurnTo(tt.args.r); got != tt.want {
				t.Errorf("Direction.TurnTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_ApplyDelta(t *testing.T) {
	type fields struct {
		x int
		y int
		d Direction
	}
	type args struct {
		d Delta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		{"Applying delta on position", fields{0, 0, North}, args{Delta{1, 0}}, Position{1, 0, North}},
		{"Applying delta on position", fields{0, 0, North}, args{Delta{0, 1}}, Position{0, 1, North}},
		{"Applying delta on position", fields{0, 0, North}, args{Delta{0, -1}}, Position{0, -1, North}},
		{"Applying delta on position", fields{0, 0, North}, args{Delta{-1, 0}}, Position{-1, 0, North}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				x: tt.fields.x,
				y: tt.fields.y,
				d: tt.fields.d,
			}
			if got := p.ApplyDelta(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Position.ApplyDelta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_MoveForward(t *testing.T) {
	type fields struct {
		x int
		y int
		d Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   Position
	}{
		{"Move forward", fields{0, 0, North}, Position{0, 1, North}},
		{"Move forward", fields{0, 0, East}, Position{1, 0, East}},
		{"Move forward", fields{0, 0, South}, Position{0, -1, South}},
		{"Move forward", fields{0, 0, West}, Position{-1, 0, West}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				x: tt.fields.x,
				y: tt.fields.y,
				d: tt.fields.d,
			}
			if got := p.MoveForward(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Position.MoveForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_ApplyCommand(t *testing.T) {
	type fields struct {
		x int
		y int
		d Direction
	}
	type args struct {
		c Command
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				x: tt.fields.x,
				y: tt.fields.y,
				d: tt.fields.d,
			}
			if got := p.ApplyCommand(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Position.ApplyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
