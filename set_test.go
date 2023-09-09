package collections

import (
	"reflect"
	"testing"
)

func TestFromArray(t *testing.T) {
	type args[T comparable] struct {
		a []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want *Set[T]
	}
	intTests := []testCase[int]{
		{name: "ints1", args: args[int]{a: []int{1, 2, 3, 4, 1, 2, 3, 4}}, want: &Set[int]{elems: []int{1, 2, 3, 4}}},
		{name: "ints2", args: args[int]{a: []int{3, 2, 1, 2, 1}}, want: &Set[int]{elems: []int{3, 2, 1}}},
		{name: "ints3", args: args[int]{a: []int{}}, want: &Set[int]{make([]int, 0)}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromArray(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSet(t *testing.T) {
	type args struct {
		initialSize int
	}
	type testCase[T comparable] struct {
		name string
		args args
		want *Set[T]
	}
	tests := []testCase[string]{
		{name: "size10", args: args{initialSize: 10}, want: &Set[string]{elems: make([]string, 0, 10)}},
		{name: "size0", args: args{initialSize: 0}, want: &Set[string]{elems: make([]string, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet[string](tt.args.initialSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	strSet := &Set[string]{[]string{"four"}}
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		elem T
		want *Set[T]
	}
	tests := []testCase[string]{
		{"new elem", strSet, "two", &Set[string]{[]string{"four", "two"}}},
		{"existing elem", strSet, "four", &Set[string]{[]string{"four", "two"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.elem)
			if !reflect.DeepEqual(strSet, tt.want) {
				t.Errorf("Add(elem T) = %v, want %v", strSet, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	strSet := &Set[string]{[]string{"four", "two", "zero"}}
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		arg  T
		want bool
	}
	tests := []testCase[string]{
		{"existing elem", strSet, "two", true},
		{"non existing elem", strSet, "three", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.arg); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {

	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		s2   *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{"same sets", &Set[int]{[]int{1, 2, 3}}, &Set[int]{[]int{1, 2, 3}}, &Set[int]{make([]int, 0)}},
		{"disjoint sets", &Set[int]{[]int{1, 2, 3}}, &Set[int]{[]int{4, 5, 6}}, &Set[int]{[]int{1, 2, 3}}},
		{"similar sets", &Set[int]{[]int{1, 2, 3, 4, 5}}, &Set[int]{[]int{3, 4, 1}}, &Set[int]{[]int{2, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		s2   *Set[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{"same sets", &Set[int]{[]int{1, 2, 3}}, &Set[int]{[]int{1, 2, 3}}, &Set[int]{[]int{1, 2, 3}}},
		{"disjoint sets", &Set[int]{[]int{1, 2, 3}}, &Set[int]{[]int{4, 5, 6}}, &Set[int]{make([]int, 0)}},
		{"overlapping sets", &Set[int]{[]int{1, 2, 3, 4, 5}}, &Set[int]{[]int{3, 4, 1}}, &Set[int]{[]int{3, 4, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubsetOf(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		s2   *Set[T]
		want bool
	}
	tests := []testCase[string]{
		{"same sets", &Set[string]{[]string{"two", "four"}}, &Set[string]{[]string{"four", "two"}}, true},
		{"disjoint sets", &Set[string]{[]string{"two"}}, &Set[string]{[]string{"four"}}, false},
		{"similar sets", &Set[string]{[]string{"four"}}, &Set[string]{[]string{"four", "two"}}, true},
		{"empty set", &Set[string]{make([]string, 0)}, &Set[string]{[]string{"four", "two"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSubsetOf(tt.s2); got != tt.want {
				t.Errorf("IsSubsetOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Items(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want []T
	}
	tests := []testCase[int]{
		{"test1", Set[int]{[]int{1, 2, 3}}, []int{1, 2, 3}},
		{"test2", Set[int]{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want int
	}
	tests := []testCase[int]{
		{"size2", Set[int]{[]int{1, 2}}, 2},
		{"size0", Set[int]{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {

	type testCase[T comparable] struct {
		name string
		s    *Set[T]
		s2   *Set[T]
		want *Set[T]
	}
	tests := []testCase[string]{
		{"same sets", &Set[string]{[]string{"two", "four"}}, &Set[string]{[]string{"four", "two"}}, &Set[string]{[]string{"two", "four"}}},
		{"disjoint sets", &Set[string]{[]string{"two"}}, &Set[string]{[]string{"four"}}, &Set[string]{[]string{"two", "four"}}},
		{"similar sets", &Set[string]{[]string{"four"}}, &Set[string]{[]string{"four", "two"}}, &Set[string]{[]string{"two", "four"}}},
		{"empty set", &Set[string]{make([]string, 0)}, &Set[string]{[]string{"four", "two"}}, &Set[string]{[]string{"two", "four"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Union(tt.s2)
		})
	}
}
