package sol

import (
	"reflect"
	"testing"
)

func CreateGraph(graph [][]int) *Node {
	nLen := len(graph)
	if nLen == 0 {
		return nil
	}
	temp := make([]*Node, nLen)
	for idx := range temp {
		temp[idx] = &Node{Val: idx + 1}
	}
	for idx := range temp {
		temp[idx].Neighbors = make([]*Node, len(graph[idx]))
		for i, val := range graph[idx] {
			temp[idx].Neighbors[i] = temp[val-1]
		}
	}
	return temp[0]
}

func BenchmarkTest(b *testing.B) {
	node := CreateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	for idx := 0; idx < b.N; idx++ {
		cloneGraph(node)
	}
}
func Test_cloneGraph(t *testing.T) {
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
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) &&
				(got.Val != tt.want.Val && len(got.Neighbors) != len(tt.want.Neighbors)) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
