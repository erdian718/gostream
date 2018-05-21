package stream_test

import (
	"testing"

	"ofunc/stream"
)

func TestMake(t *testing.T) {
	s := stream.Make(10, 11, 12)
	for i := 0; i < 3; i++ {
		if s.Head() != i+10 {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestConcat(t *testing.T) {
	s0 := stream.Make(0, 1, 2)
	s1 := stream.Make(3, 4, 5)
	s2 := stream.Make(6, 7, 8)
	s := stream.Concat(s0, s1, s2)
	for i := 0; i < 9; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestRepeat(t *testing.T) {
	ok := stream.Repeat(6).Take(8).All(func(x interface{}) bool {
		return x.(int) == 6
	})
	if !ok {
		t.FailNow()
	}
}

func TestGrow(t *testing.T) {
	s := stream.Grow(-5, 2).Take(10)
	k := -5
	for i := 0; i < 10; i++ {
		if s.Head() != k {
			t.FailNow()
		}
		s = s.Tail()
		k += 2
	}
	if s != nil {
		t.FailNow()
	}
}

func TestN(t *testing.T) {
	s := stream.N().Take(10)
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestRange(t *testing.T) {
	s := stream.Range(-3, 3)
	for i := -3; i < 3; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}

	s = stream.Range(3, -3)
	for i := 3; i > -3; i-- {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}
