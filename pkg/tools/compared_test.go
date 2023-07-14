package tools

import (
	"reflect"
	"testing"
)

func TestMaxVal4Slice(t *testing.T) {
	type args[T Ordered] struct {
		s []T
	}
	type testCase[T Ordered] struct {
		name    string
		args    args[T]
		wantMax T
		wantErr bool
	}
	tests := []testCase[float64]{
		// TODO: Add test cases.
		{name: "t1", args: args[float64]{s: []float64{2.1}}, wantMax: 2.1, wantErr: false},
		{name: "t2", args: args[float64]{s: []float64{2.6, 5.0}}, wantMax: 5.0, wantErr: false},
		{name: "t3", args: args[float64]{s: []float64{0, 3}}, wantMax: 3.0, wantErr: false},
		{name: "t4", args: args[float64]{s: []float64{10, 3, 7, 9}}, wantMax: 10, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, err := MaxVal4Slice(tt.args.s...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxVal4Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MaxVal4Slice() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestMinVal4Slice(t *testing.T) {
	type args[T Ordered] struct {
		s []T
	}
	type testCase[T Ordered] struct {
		name    string
		args    args[T]
		wantMin T
		wantErr bool
	}
	tests := []testCase[float64]{
		// TODO: Add test cases.
		{name: "t1", args: args[float64]{s: []float64{2.1}}, wantMin: 2.1, wantErr: false},
		{name: "t2", args: args[float64]{s: []float64{2.6, 5.0}}, wantMin: 2.6, wantErr: false},
		{name: "t3", args: args[float64]{s: []float64{0, 3}}, wantMin: 0, wantErr: false},
		{name: "t4", args: args[float64]{s: []float64{10, 3, 7, 9}}, wantMin: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, err := MinVal4Slice(tt.args.s...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinVal4Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("MinVal4Slice() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}

func TestMaxMin4Slice(t *testing.T) {
	type args[T Ordered] struct {
		s []T
	}
	type testCase[T Ordered] struct {
		name    string
		args    args[T]
		wantMax T
		wantMin T
		wantErr bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{name: "t1", args: args[int]{s: []int{1, 2, 3, 4, 10}}, wantMin: 1, wantMax: 10, wantErr: false},
		{name: "t2", args: args[int]{s: []int{1, -2, 3, 4, 10}}, wantMin: -2, wantMax: 10, wantErr: false},
		{name: "t3", args: args[int]{s: []int{1, -2, 3, -4, 2}}, wantMin: -4, wantMax: 3, wantErr: false},
		{name: "t4", args: args[int]{s: []int{}}, wantMin: 0, wantMax: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin, err := MaxMin4Slice(tt.args.s...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxMin4Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MaxMin4Slice() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("MaxMin4Slice() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
