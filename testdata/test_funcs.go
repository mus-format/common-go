package testdata

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/mok"
)

func TestUnmarshalResults[T any](wantV, v T, wantN, n int, wantErr, err error,
	mocks []*mok.Mock, t *testing.T) {
	if !reflect.DeepEqual(v, wantV) {
		t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
	}
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
	if info := mok.CheckCalls(mocks); len(info) > 0 {
		t.Error(info)
	}
}

func TestSkipResults(wantN, n int, wantErr, err error, t *testing.T) {
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
}

func ComparePtrs(t, v any) bool {
	p1 := reflect.ValueOf(t).Pointer()
	p2 := reflect.ValueOf(v).Pointer()
	return p1 == p2
}
