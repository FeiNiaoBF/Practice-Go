package g2cache

import "testing"

func TestLen(t *testing.T) {
	bv := ByteView{[]byte("hello")}
	if bv.Len() != 5 {
		t.Fatal("Len error")
	}
}

func TestByteSlice(t *testing.T) {
	bv := ByteView{[]byte("hello")}
	if string(bv.ByteSlice()) != "hello" {
		t.Fatal("ByteSlice error")
	}
}

func TestString(t *testing.T) {
	bv := ByteView{[]byte("hello")}
	if bv.String() != "hello" {
		t.Fatal("String error")
	}
}
