package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "2 канала с данными",
			input: [][]int{{1, 2}, {3, 4}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Один из каналов пустой",
			input: [][]int{{1, 2}, {}, {3}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Все каналы пустые",
			input: [][]int{{}, {}, {}},
			want:  []int{},
		},
		{
			name:  "Пустой вызов",
			input: [][]int{},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chans []<-chan int

			for _, row := range tt.input {
				ch := make(chan int, len(row))

				for _, num := range row {
					ch <- num
				}

				close(ch)
				chans = append(chans, ch)
			}

			gotChan := Merge(chans...)
			gotResult := make([]int, 0)

		Loop:
			for {
				var num int
				var open bool

				select {

				case num, open = <-gotChan:
					if !open {
						break Loop
					}
					gotResult = append(gotResult, num)
				case <-time.After(100 * time.Millisecond):
					t.Fatalf("Merge() завис и не закрыл канал")
				}
			}

			sort.Ints(gotResult)
			if !reflect.DeepEqual(tt.want, gotResult) {
				t.Errorf("Ожидали: %v, получили: %v", tt.want, gotResult)
			}
		})
	}
}
