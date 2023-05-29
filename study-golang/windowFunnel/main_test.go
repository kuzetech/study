package windowFunnel

import (
	"reflect"
	"testing"
)

func Test_findFunnelClickHouse(t *testing.T) {
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
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 3},
				},
				len:    3,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
				{time: 2, sourceIndex: 2},
				{time: 4, sourceIndex: 4},
			},
		},
		{
			name: "2",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 3},
					{time: 4, step: 2},
					{time: 5, step: 4},
				},
				len:    4,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
				{time: 2, sourceIndex: 2},
				{time: 3, sourceIndex: 3},
				{time: 5, sourceIndex: 5},
			},
		},
		{
			name: "3",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 3},
					{time: 6, step: 1},
					{time: 7, step: 2},
					{time: 8, step: 1},
					{time: 9, step: 2},
					{time: 10, step: 3},
				},
				len:    3,
				window: 10,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
				{time: 4, sourceIndex: 4},
				{time: 5, sourceIndex: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFunnelClickHouse(tt.args.data, tt.args.len, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFunnelClickHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFunnel(t *testing.T) {
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
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 2},
					{time: 6, step: 3},
					{time: 7, step: 1},
					{time: 8, step: 2},
					{time: 9, step: 1},
					{time: 10, step: 2},
					{time: 11, step: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, sourceIndex: 3},
				{time: 5, sourceIndex: 5},
				{time: 6, sourceIndex: 6},
			},
		},
		{
			// ABABCABABC 找 ABC 结果 345
			name: "2",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 3},
					{time: 6, step: 1},
					{time: 7, step: 2},
					{time: 8, step: 1},
					{time: 9, step: 2},
					{time: 10, step: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, sourceIndex: 3},
				{time: 4, sourceIndex: 4},
				{time: 5, sourceIndex: 5},
			},
		},
		{
			// ABAC 找 ABC 结果 124
			name: "3",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
				{time: 2, sourceIndex: 2},
				{time: 4, sourceIndex: 4},
			},
		},
		{
			// ABCBD 找 ABCD 结果 1235
			name: "4",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 3},
					{time: 4, step: 2},
					{time: 5, step: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
				{time: 2, sourceIndex: 2},
				{time: 3, sourceIndex: 3},
				{time: 5, sourceIndex: 5},
			},
		},
		{
			// ABABCABABC 找 ABCD 结果 345
			name: "5",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 3},
					{time: 6, step: 1},
					{time: 7, step: 2},
					{time: 8, step: 1},
					{time: 9, step: 2},
					{time: 10, step: 3},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, sourceIndex: 3},
				{time: 4, sourceIndex: 4},
				{time: 5, sourceIndex: 5},
			},
		},
		{
			// AAABBBCCCDDD 找 ABCD 结果 369 10
			name: "6",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 1},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 2},
					{time: 6, step: 2},
					{time: 7, step: 3},
					{time: 8, step: 3},
					{time: 9, step: 3},
					{time: 10, step: 4},
					{time: 11, step: 4},
					{time: 12, step: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 3, sourceIndex: 3},
				{time: 6, sourceIndex: 6},
				{time: 9, sourceIndex: 9},
				{time: 10, sourceIndex: 10},
			},
		},
		{
			// ABABCABCD 找 ABCD 结果 6789
			name: "7",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 2},
					{time: 3, step: 1},
					{time: 4, step: 2},
					{time: 5, step: 3},
					{time: 6, step: 1},
					{time: 7, step: 2},
					{time: 8, step: 3},
					{time: 9, step: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 6, sourceIndex: 6},
				{time: 7, sourceIndex: 7},
				{time: 8, sourceIndex: 8},
				{time: 9, sourceIndex: 9},
			},
		},
		{
			// AABBCCAB 找 ABCD 结果 246
			name: "8",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 1},
					{time: 3, step: 2},
					{time: 4, step: 2},
					{time: 5, step: 3},
					{time: 6, step: 3},
					{time: 7, step: 1},
					{time: 8, step: 2},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 2, sourceIndex: 2},
				{time: 4, sourceIndex: 4},
				{time: 5, sourceIndex: 5},
			},
		},
		{
			// AABACBBD 找 ABCD 结果 2358
			name: "9",
			args: args{
				data: []Event{
					{time: 1, step: 1},
					{time: 2, step: 1},
					{time: 3, step: 2},
					{time: 4, step: 1},
					{time: 5, step: 3},
					{time: 6, step: 2},
					{time: 7, step: 2},
					{time: 8, step: 4},
				},
				length: 4,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 2, sourceIndex: 2},
				{time: 3, sourceIndex: 3},
				{time: 5, sourceIndex: 5},
				{time: 8, sourceIndex: 8},
			},
		},
		{
			// DDDDDDDD 找 ABC 结果
			name: "10",
			args: args{
				data: []Event{
					{time: 1, step: 4},
					{time: 2, step: 4},
					{time: 3, step: 4},
					{time: 4, step: 4},
					{time: 5, step: 4},
					{time: 6, step: 4},
					{time: 7, step: 4},
					{time: 8, step: 4},
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
					{time: 1, step: 3},
					{time: 2, step: 3},
					{time: 3, step: 4},
					{time: 4, step: 4},
					{time: 5, step: 4},
					{time: 6, step: 4},
					{time: 7, step: 4},
					{time: 8, step: 4},
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
					{time: 1, step: 2},
					{time: 2, step: 3},
					{time: 3, step: 4},
					{time: 4, step: 4},
					{time: 5, step: 4},
					{time: 6, step: 4},
					{time: 7, step: 4},
					{time: 8, step: 4},
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
					{time: 1, step: 1},
					{time: 2, step: 1},
					{time: 3, step: 4},
					{time: 4, step: 4},
					{time: 5, step: 4},
					{time: 6, step: 4},
					{time: 7, step: 4},
					{time: 8, step: 4},
				},
				length: 3,
				window: 100,
			},
			want: []EventsTimestamp{
				{time: 1, sourceIndex: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFunnel(tt.args.data, tt.args.length, tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}
