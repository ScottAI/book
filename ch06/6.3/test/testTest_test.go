package testTest

import "testing"

func TestFb1(t *testing.T)  {
	if fb1(0) != 0 {
		t.Error(`fb1(0)!=0`)
	}
	if fb1(1) != 1 {
		t.Error(`fb1(1)!=1`)
	}
	if fb1(2) != 1 {
		t.Error(`fb1(2)!=1`)
	}
	if fb1(10) != 55 {
		t.Error(`fb1(10)!=55`)
	}
}

func TestFb2(t *testing.T)  {
	if fb2(0) != 0 {
		t.Error(`fb2(0)!=0`)
	}
	if fb2(1) != 1 {
		t.Error(`fb2(1)!=1`)
	}
	if fb2(2) != 1 {
		t.Error(`fb2(2)!=1`)
	}
	if fb2(10) != 55 {
		t.Error(`fb2(10)!=55`)
	}
}
