package bm

import (
	"errors"
	"sync"
)

var (
	ErrOutOfBounds = errors.New("out of bounds")
)

type Bitmap struct {
	bits []uint8
	size int64
	mu   sync.Mutex
}

// New creates a new bitmap of the given size (number of bits to manage)
func New(size int64) *Bitmap {
	b := (size + 7) / 8
	return &Bitmap{
		bits: make([]uint8, b),
		size: size,
	}
}

// Set sets the given bit to 1 in the bitmap
// ErrOutOfBounds is returned if bit is -ve or larger than the size of the bitmap
func (b *Bitmap) Set(bit int64) error {
	if bit < 0 || bit >= b.size {
		return ErrOutOfBounds
	}
	idx := bit / 8
	offset := bit % 8
	b.mu.Lock()
	defer b.mu.Unlock()
	b.bits[idx] = b.bits[idx] | (1 << uint8(offset))
	return nil
}

// Clear unsets the given bit (sets it to 0)
// ErrOutOfBounds is returned if bit is -ve or larger than the size of the bitmap
func (b *Bitmap) Clear(bit int64) error {
	if bit < 0 || bit >= b.size {
		return ErrOutOfBounds
	}
	idx := bit / 8
	offset := bit % 8
	b.mu.Lock()
	defer b.mu.Unlock()
	b.bits[idx] &= ^(1 << uint8(offset))
	return nil
}

// IsSet returns true if a bit is set, otherwise returns false
// If the given bit is out of bounds, it returns false
func (b *Bitmap) IsSet(bit int64) bool {
	if bit < 0 || bit >= b.size {
		return false
	}
	idx := bit / 8
	offset := bit % 8
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.bits[idx]&(1<<uint8(offset)) != 0
}

func (b *Bitmap) Size() int64 {
	return b.size
}
