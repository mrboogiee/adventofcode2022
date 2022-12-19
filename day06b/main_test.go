package main

import "testing"

func Test_findMarker(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "test1",
			args: args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			want: 19,
		},
		{
			name: "test2",
			args: args{"bvwbjplbgvbhsrlpgdmjqwftvncz"},
			want: 23,
		},
		{
			name: "test3",
			args: args{"nppdvjthqldpwncqszvftbrmjlhg"},
			want: 23,
		},
		{
			name: "test4",
			args: args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			want: 29,
		},
		{
			name: "test5",
			args: args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			want: 26,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.line); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
