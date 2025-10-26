package common

import (
	"testing"
	"unsafe"
)

func TestPtrMap(t *testing.T) {
	t.Run("Put and Get should work correctly", func(t *testing.T) {
		var (
			wantKey1 = 0
			wantKey2 = 1
			wantKey3 = 0
			n1       = 1
			n2       = 2
			n3       = 3
			ptr1     = unsafe.Pointer(&n1)
			ptr2     = unsafe.Pointer(&n2)
			ptr3     = unsafe.Pointer(&n3)
			mp       = NewPtrMap()
			key1     = mp.Put(ptr1)
			key2     = mp.Put(ptr2)
			key3     = 0
		)
		if key1 != wantKey1 {
			t.Errorf("unexpected key1, want %v actual %v", wantKey1, key1)
		}
		if key2 != wantKey2 {
			t.Errorf("unexpected key1, want %v actual %v", wantKey2, key2)
		}
		key1, pst := mp.Get(ptr1)
		if key1 != wantKey1 {
			t.Errorf("unexpected key1, want %v actual %v", wantKey1, key1)
		}
		if pst != true {
			t.Errorf("unexpected pst, want %v actual %v", true, pst)
		}
		key2, pst = mp.Get(ptr2)
		if key2 != wantKey2 {
			t.Errorf("unexpected key1, want %v actual %v", wantKey2, key2)
		}
		if pst != true {
			t.Errorf("unexpected pst, want %v actual %v", true, pst)
		}
		key3, pst = mp.Get(ptr3)
		if key3 != wantKey3 {
			t.Errorf("unexpected key3, want %v actual %v", wantKey3, key3)
		}
		if pst != false {
			t.Errorf("unexpected pst, want %v actual %v", false, pst)
		}
	})
}

func TestReversePtrMap(t *testing.T) {
	t.Run("Put and Get should work correctly", func(t *testing.T) {
		var (
			n1   = 1
			n2   = 2
			ptr1 = unsafe.Pointer(&n1)
			ptr2 = unsafe.Pointer(&n2)
			key1 = 0
			key2 = 1
			key3 = 3
			mp   = NewReversePtrMap()
		)
		mp.Put(key1, ptr1)
		mp.Put(key2, ptr2)
		ptr, pst := mp.Get(key1)
		if ptr != ptr1 {
			t.Errorf("unexpected ptr, want %v actual %v", ptr1, ptr)
		}
		if pst != true {
			t.Errorf("unexpected pst, want %v actual %v", true, pst)
		}
		ptr, pst = mp.Get(key2)
		if ptr != ptr2 {
			t.Errorf("unexpected ptr, want %v actual %v", ptr2, ptr)
		}
		if pst != true {
			t.Errorf("unexpected pst, want %v actual %v", true, pst)
		}
		ptr, pst = mp.Get(key3)
		if ptr != nil {
			t.Errorf("unexpected ptr, want %v actual %v", nil, ptr)
		}
		if pst != false {
			t.Errorf("unexpected pst, want %v actual %v", false, pst)
		}
	})
}
