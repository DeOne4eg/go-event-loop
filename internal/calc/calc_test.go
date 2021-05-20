package calc

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want float32
	}{
		{
			args: args{
				a: 10,
				b: 10,
			},
			want: 20,
		},
		{
			args: args{
				a: 25,
				b: 32,
			},
			want: 57,
		},
		{
			args: args{
				a: 124,
				b: 26,
			},
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run("sum", func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want float32
	}{
		{
			args: args{
				a: 10,
				b: 10,
			},
			want: 1,
		},
		{
			args: args{
				a: 25,
				b: 5,
			},
			want: 5,
		},
		{
			args: args{
				a: 1,
				b: 2,
			},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run("divide", func(t *testing.T) {
			if got := Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMut(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want float32
	}{
		{
			args: args{
				a: 5,
				b: 2,
			},
			want: 10,
		},
		{
			args: args{
				a: 43,
				b: 3,
			},
			want: 129,
		},
		{
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run("mut", func(t *testing.T) {
			if got := Mut(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want float32
	}{
		{
			args: args{
				a: 100,
				b: 10,
			},
			want: 90,
		},
		{
			args: args{
				a: 111,
				b: 11,
			},
			want: 100,
		},
		{
			args: args{
				a: 1,
				b: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run("sub", func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
