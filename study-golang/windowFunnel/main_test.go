package windowFunnel

import (
	"reflect"
	"testing"
)

func Test_test(t *testing.T) {
	type args struct {
		data   []Event
		len    int64
		window int64
	}
	tests := []struct {
		name string
		args args
		want []EventsTimestamp
	}{
		{
			name: "1",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 3},
				},
				len:    3,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
				{time: 2, index: 2},
				{time: 4, index: 4},
			},
		},
		{
			name: "2",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 3},
					{time: 4, index: 2},
					{time: 5, index: 4},
				},
				len:    4,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
				{time: 2, index: 2},
				{time: 3, index: 3},
				{time: 5, index: 5},
			},
		},
		{
			name: "3",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 3},
					{time: 6, index: 1},
					{time: 7, index: 2},
					{time: 8, index: 1},
					{time: 9, index: 2},
					{time: 10, index: 3},
				},
				len:    3,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
				{time: 4, index: 4},
				{time: 5, index: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(tt.args.data, tt.args.len, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test2(t *testing.T) {
	type args struct {
		data   []Event
		length int64
		window int64
	}
	tests := []struct {
		name string
		args args
		want []EventsTimestamp
	}{
		{
			// ABABBCABABC 找 ABC 结果 356
			name: "1",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 2},
					{time: 6, index: 3},
					{time: 7, index: 1},
					{time: 8, index: 2},
					{time: 9, index: 1},
					{time: 10, index: 2},
					{time: 11, index: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, index: 3},
				{time: 5, index: 5},
				{time: 6, index: 6},
			},
		},
		{
			// ABABCABABC 找 ABC 结果 356
			name: "2",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 3},
					{time: 6, index: 1},
					{time: 7, index: 2},
					{time: 8, index: 1},
					{time: 9, index: 2},
					{time: 10, index: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, index: 3},
				{time: 4, index: 4},
				{time: 5, index: 5},
			},
		},
		{
			// ABAC 找 ABC 结果 124
			name: "3",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
				{time: 2, index: 2},
				{time: 4, index: 4},
			},
		},
		{
			// ABCBD 找 ABCD 结果 1235
			name: "4",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 3},
					{time: 4, index: 2},
					{time: 5, index: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
				{time: 2, index: 2},
				{time: 3, index: 3},
				{time: 5, index: 5},
			},
		},
		{
			// ABABCABABC 找 ABCD 结果 345
			name: "5",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 3},
					{time: 6, index: 1},
					{time: 7, index: 2},
					{time: 8, index: 1},
					{time: 9, index: 2},
					{time: 10, index: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, index: 3},
				{time: 4, index: 4},
				{time: 5, index: 5},
			},
		},
		{
			// AAABBBCCCDDD 找 ABCD 结果 369 10
			name: "6",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 1},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 2},
					{time: 6, index: 2},
					{time: 7, index: 3},
					{time: 8, index: 3},
					{time: 9, index: 3},
					{time: 10, index: 4},
					{time: 11, index: 4},
					{time: 12, index: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, index: 3},
				{time: 6, index: 6},
				{time: 9, index: 9},
				{time: 10, index: 10},
			},
		},
		{
			// ABABCABCD 找 ABCD 结果 6789
			name: "7",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 2},
					{time: 3, index: 1},
					{time: 4, index: 2},
					{time: 5, index: 3},
					{time: 6, index: 1},
					{time: 7, index: 2},
					{time: 8, index: 3},
					{time: 9, index: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 6, index: 6},
				{time: 7, index: 7},
				{time: 8, index: 8},
				{time: 9, index: 9},
			},
		},
		{
			// AABBCCAB 找 ABCD 结果 246
			name: "8",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 1},
					{time: 3, index: 2},
					{time: 4, index: 2},
					{time: 5, index: 3},
					{time: 6, index: 3},
					{time: 7, index: 1},
					{time: 8, index: 2},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 2, index: 2},
				{time: 4, index: 4},
				{time: 5, index: 5},
			},
		},
		{
			// AABACBBD 找 ABCD 结果 2358
			name: "9",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 1},
					{time: 3, index: 2},
					{time: 4, index: 1},
					{time: 5, index: 3},
					{time: 6, index: 2},
					{time: 7, index: 2},
					{time: 8, index: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 2, index: 2},
				{time: 3, index: 3},
				{time: 5, index: 5},
				{time: 8, index: 8},
			},
		},
		{
			// DDDDDDDD 找 ABC 结果
			name: "10",
			args: args{
				data: []Event{
					{time: 1, index: 4},
					{time: 2, index: 4},
					{time: 3, index: 4},
					{time: 4, index: 4},
					{time: 5, index: 4},
					{time: 6, index: 4},
					{time: 7, index: 4},
					{time: 8, index: 4},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{},
		},
		{
			// CCDDDDDD 找 ABC 结果
			name: "11",
			args: args{
				data: []Event{
					{time: 1, index: 3},
					{time: 2, index: 3},
					{time: 3, index: 4},
					{time: 4, index: 4},
					{time: 5, index: 4},
					{time: 6, index: 4},
					{time: 7, index: 4},
					{time: 8, index: 4},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{},
		},
		{
			// BCDDDDDD 找 ABC 结果
			name: "12",
			args: args{
				data: []Event{
					{time: 1, index: 2},
					{time: 2, index: 3},
					{time: 3, index: 4},
					{time: 4, index: 4},
					{time: 5, index: 4},
					{time: 6, index: 4},
					{time: 7, index: 4},
					{time: 8, index: 4},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{},
		},
		{
			// AADDDDDD 找 ABC 结果
			name: "13",
			args: args{
				data: []Event{
					{time: 1, index: 1},
					{time: 2, index: 1},
					{time: 3, index: 4},
					{time: 4, index: 4},
					{time: 5, index: 4},
					{time: 6, index: 4},
					{time: 7, index: 4},
					{time: 8, index: 4},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, index: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test2(tt.args.data, tt.args.length, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("test2() = %v, want %v", got, tt.want)
			}
		})
	}
}
