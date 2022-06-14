package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTestDFS(b *testing.B) {
	node := CreateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	for idx := 0; idx < b.N; idx++ {
		cloneGraphDFS(node)
	}
}
func Test_cloneGraphDFS(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Example1",
			args: args{node: CreateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
			want: CreateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			name: "Example2",
			args: args{node: CreateGraph([][]int{{}})},
			want: CreateGraph([][]int{{}}),
		},
		{
			name: "Example3",
			args: args{node: CreateGraph([][]int{})},
			want: CreateGraph([][]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraphDFS(tt.args.node); !reflect.DeepEqual(got, tt.want) &&
				(got.Val != tt.want.Val && len(got.Neighbors) != len(tt.want.Neighbors)) {
				t.Errorf("cloneGraphDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
