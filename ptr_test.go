package common

import (
	"testing"
	"unsafe"

	asserterror "github.com/ymz-ncnk/assert/error"
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
		asserterror.Equal(t, key1, wantKey1, "unexpected key1")
		asserterror.Equal(t, key2, wantKey2, "unexpected key2")

		key1, pst := mp.Get(ptr1)
		asserterror.Equal(t, key1, wantKey1, "unexpected key1")
		asserterror.Equal(t, pst, true, "unexpected pst")

		key2, pst = mp.Get(ptr2)
		asserterror.Equal(t, key2, wantKey2, "unexpected key2")
		asserterror.Equal(t, pst, true, "unexpected pst")

		key3, pst = mp.Get(ptr3)
		asserterror.Equal(t, key3, wantKey3, "unexpected key3")
		asserterror.Equal(t, pst, false, "unexpected pst")
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
		asserterror.Equal(t, ptr, ptr1, "unexpected ptr")
		asserterror.Equal(t, pst, true, "unexpected pst")

		ptr, pst = mp.Get(key2)
		asserterror.Equal(t, ptr, ptr2, "unexpected ptr")
		asserterror.Equal(t, pst, true, "unexpected pst")

		ptr, pst = mp.Get(key3)
		asserterror.Equal(t, ptr, nil, "unexpected ptr")
		asserterror.Equal(t, pst, false, "unexpected pst")
	})
}
