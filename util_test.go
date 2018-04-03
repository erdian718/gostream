package stream_test

import (
	"math"
	"testing"

	"ofunc/stream"
)

func TestMake(t *testing.T) {
	s := stream.Make(10, 11, 12)
	for i := 0; i < 3; i++ {
		if s.Head() != i+10 {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
}

func TestConcat(t *testing.T) {
	s0 := stream.Make(0, 1, 2)
	s1 := stream.Make(3, 4, 5)
	s2 := stream.Make(6, 7, 8)
	s := stream.Concat(s0, s1, s2)
	for i := 0; i < 9; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
}

func TestRepeat(t *testing.T) {
	ok := stream.Repeat(6).Take(8).All(func(x interface{}) bool {
		return x.(int) == 6
	})
	if !ok {
		t.Fail()
	}
}

func TestGrow(t *testing.T) {
	s := stream.Grow(-5, 2).Take(10)
	k := -5
	for i := 0; i < 10; i++ {
		if s.Head() != k {
			t.Fail()
		}
		s = s.Tail()
		k += 2
	}
	if s != nil {
		t.Fail()
	}
}

func TestN(t *testing.T) {
	s := stream.N().Take(10)
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
}

func TestRange(t *testing.T) {
	s := stream.Range(-3, 3)
	for i := -3; i < 3; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}

	s = stream.Range(3, -3)
	for i := 3; i > -3; i-- {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
}

func TestRand(t *testing.T) {
	xs := stream.Rand()
	ys := xs.Drop(10)
	zs := ys.Drop(10)

	s := xs.Take(30).Fold(0.0, func(a interface{}, x interface{}) interface{} {
		return a.(float64) + x.(float64)
	}).(float64)
	sx := xs.Take(10).Fold(0.0, func(a interface{}, x interface{}) interface{} {
		return a.(float64) + x.(float64)
	}).(float64)
	sy := ys.Take(10).Fold(0.0, func(a interface{}, x interface{}) interface{} {
		return a.(float64) + x.(float64)
	}).(float64)
	sz := zs.Take(10).Fold(0.0, func(a interface{}, x interface{}) interface{} {
		return a.(float64) + x.(float64)
	}).(float64)
	if math.Abs(sx+sy+sz-s) > 1e-8 {
		t.Fail()
	}
}
