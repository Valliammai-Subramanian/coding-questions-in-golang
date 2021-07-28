package kMeans

import (
	"reflect"
	"testing"
	"sort"
)

type byPoints []Point
func (p byPoints) Len() int {
    return len(p)
}
func (p byPoints) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p byPoints) Less(i, j int) bool {
    return p[i][0] < p[j][0] || ((p[i][0] == p[j][0]) && p[i][1] < p[j][1])
}

func Test_kmeans(t *testing.T) {
	type args struct {
		points    []Point
		centroids []Point
		maxIter   int
		eta       float64
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "First Test",
			args: args{
				points: []Point{
					Point{3, 3},
					Point{2, 3},
					Point{2, 2},
					Point{1, 1},
					Point{0, 0},
					Point{1, -1},
					Point{-1, -1},
					Point{-3, -2},
					Point{-2, 0},
					Point{-2, -2},
				},
				centroids: []Point{
					Point{2, 2},
					Point{1, 1},
					Point{-3, -2},
					Point{-2, -2},
				},
				maxIter: 10,
				eta:     0.4,
			},
			want: []Point{
				Point{2.3333333333333335, 2.6666666666666665},
				Point{0.6666666666666666, 0},
				Point{-2.5, -2},
				Point{-1.5, -0.5},
			},
		},
		{
			name: "Second Test",
			args: args{
				points: []Point{
					Point{4, 4},
					Point{3, 3},
					Point{2, 2},
					Point{1, 1},
					Point{0, 0},
					Point{-1, -1},
					Point{-2, -2},
					Point{-3, -3},
					Point{-4, -4},
				},
				centroids: []Point{
					Point{3, 2},
					Point{0, 0},
					Point{-3, -2},
				},
				maxIter: 9,
				eta:     0.4,
			},
			want: []Point{
				Point{3, 3},
				Point{0, 0},
				Point{-3, -3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmeans(tt.args.points, tt.args.centroids, tt.args.maxIter, tt.args.eta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmeans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classify(t *testing.T) {
	type args struct {
		points    []Point
		centroids []Point
	}
	tests := []struct {
		name string
		args args
		want map[int][]Point
	}{
		{
			name: "First Test",
			args: args{
				points: []Point{
					Point{3, 3},
					Point{2, 3},
					Point{2, 2},
					Point{1, 1},
					Point{0, 0},
					Point{1, -1},
					Point{-1, -1},
					Point{-3, -2},
					Point{-2, 0},
					Point{-2, -2},
				},
				centroids: []Point{
					Point{2, 2},
					Point{1, 1},
					Point{-3, -2},
					Point{-2, -2},
				},
			},
			want: map[int][]Point{
				0: []Point{
					Point{2, 2},
					Point{2, 3},
					Point{3, 3},
				},
				1: []Point{
					Point{0, 0},
					Point{1, -1},
					Point{1, 1},
				},
				2: []Point{
					Point{-3, -2},
				},
				3: []Point{
					Point{-2, -2},
					Point{-2, 0},
					Point{-1, -1},
				},
			},
		},
		{
			name: "Second Test",
			args: args{
				points: []Point{
					Point{4, 4},
					Point{3, 3},
					Point{2, 2},
					Point{1, 1},
					Point{0, 0},
					Point{-1, -1},
					Point{-2, -2},
					Point{-3, -3},
					Point{-4, -4},
				},
				centroids: []Point{
					Point{3, 2},
					Point{0, 0},
					Point{-3, -2},
				},
			},
			want: map[int][]Point{
				0: []Point{
					Point{2, 2},
					Point{3, 3},
					Point{4, 4},
				},
				1: []Point{
					Point{-1, -1},
					Point{0, 0},
					Point{1, 1},
				},
				2: []Point{
					Point{-4, -4},
					Point{-3, -3},
					Point{-2, -2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classify(tt.args.points, tt.args.centroids)
			for index, points := range got {
				sort.Sort(byPoints(points))
				if !reflect.DeepEqual(tt.want[index], points) {
					t.Errorf("classify() = %v, want %v", got, tt.want)				
				}
			}
		})
	}
}