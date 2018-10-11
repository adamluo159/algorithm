package sort

import (
	"reflect"
	"testing"
)

func TestInsertTionSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"insert_sort",
			args{[]int{4, 0, 3, 1, 4, 2, 7, 20}},
			[]int{0, 1, 2, 3, 4, 4, 7, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertTionSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertTionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
