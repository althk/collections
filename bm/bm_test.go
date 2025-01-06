package bm

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitmap_Size(t *testing.T) {

	tests := []struct {
		name string
		size int64
		want int64
	}{
		{"size_3", 3, 3},
		{"size_16", 16, 16},
		{"size_32", 32, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(tt.size)
			if got := b.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_Set(t *testing.T) {
	b := New(16)

	tests := []struct {
		name string
		bit  int64
		err  error
	}{
		{"set_2", 2, nil},
		{"set_3", 3, nil},
		{"set_16", 16, ErrOutOfBounds},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := b.Set(tt.bit); !errors.Is(err, tt.err) {
				t.Errorf("Set() error = %v, want err %v", err, tt.err)
			}
		})
	}
}

func TestBitmap_IsSet(t *testing.T) {
	b := New(16)
	_ = b.Set(2)
	_ = b.Set(3)
	_ = b.Set(16)
	tests := []struct {
		name string
		bit  int64
		want bool
	}{
		{"is_set_2_true", 2, true},
		{"is_set_3_true", 3, true},
		{"is_set_1_false", 1, false},
		{"is_set_16_false", 16, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, b.IsSet(tt.bit), tt.want)
		})
	}
}

func TestBitmap_Clear(t *testing.T) {
	b := New(16)
	_ = b.Set(2)
	_ = b.Set(3)
	_ = b.Set(1)
	tests := []struct {
		name string
		bit  int64
		want bool
		err  error
	}{
		{"clear_2", 2, false, nil},
		{"clear_3", 3, false, nil},
		{"clear_16_err", 16, false, ErrOutOfBounds},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := b.Clear(tt.bit); !errors.Is(err, tt.err) {
				t.Errorf("Clear() error = %v, want err %v", err, tt.err)
			} else {
				require.Equal(t, b.IsSet(tt.bit), tt.want)
			}
		})
	}
}

func BenchmarkBitmap_Set(b *testing.B) {
	sizes := []int64{10000, 100000, 1000000}
	for _, size := range sizes {
		bmap := New(size)
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bmap.Set(int64(i) % size)
			}
		})
	}
}
