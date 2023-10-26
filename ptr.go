package com

import "unsafe"

// PtrFlag defines the metadata used to encode pointers.
type PtrFlag byte

// These constants determine the values of PtrFlag.
const (
	NotNil PtrFlag = iota
	Nil
	Mapping
)

// NewPtrMap creates a new PtrMap.
func NewPtrMap() *PtrMap {
	return &PtrMap{
		m: make(map[unsafe.Pointer]int),
	}
}

// PtrMap helps to implement pointer mapping.
type PtrMap struct {
	sn int
	m  map[unsafe.Pointer]int
}

func (m *PtrMap) Put(ptr unsafe.Pointer) (key int) {
	key = m.sn
	m.m[ptr] = key
	m.sn++
	return
}

func (m *PtrMap) Get(ptr unsafe.Pointer) (key int, pst bool) {
	key, pst = m.m[ptr]
	return
}

// NewReversePtrMap creates a new ReversePtrMap.
func NewReversePtrMap() ReversePtrMap {
	return ReversePtrMap{
		m: make(map[int]unsafe.Pointer),
	}
}

// ReversePtrMap helps to implement pointer mapping.
type ReversePtrMap struct {
	m map[int]unsafe.Pointer
}

func (mp ReversePtrMap) Put(key int, ptr unsafe.Pointer) {
	mp.m[key] = ptr
}

func (mp ReversePtrMap) Get(key int) (ptr unsafe.Pointer, pst bool) {
	ptr, pst = mp.m[key]
	return
}
