package stream_test

import (
	"testing"

	"ofunc/stream"
)

func TestTake(t *testing.T) {
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

func TestTakeWhile(t *testing.T) {
	s := stream.N().TakeWhile(func(x interface{}) bool {
		return x.(int) < 10
	})
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
	s = s.TakeWhile(func(x interface{}) bool {
		return x == 0
	})
	if s != nil {
		t.FailNow()
	}
}

func TestDrop(t *testing.T) {
	x := stream.N().Drop(10).Head()
	if x != 10 {
		t.FailNow()
	}

	x0 := stream.N()
	x1 := x0.Drop(10)
	x2 := x1.Drop(10)
	x3 := x0.Drop(20)
	if x3 != x2 {
		t.FailNow()
	}
}

func TestDropWhile(t *testing.T) {
	x := stream.N().DropWhile(func(x interface{}) bool {
		return x.(int) < 10
	}).Head()
	if x != 10 {
		t.FailNow()
	}

	x0 := stream.N()
	x1 := x0.DropWhile(func(x interface{}) bool {
		return x.(int) < 10
	})
	x2 := x1.DropWhile(func(x interface{}) bool {
		return x.(int) < 20
	})
	x3 := x0.DropWhile(func(x interface{}) bool {
		return x.(int) < 20
	})
	if x3 != x2 {
		t.FailNow()
	}
}

func TestCut(t *testing.T) {
	s := stream.N().Cut(0).Take(10).Cut(5)
	for i := 0; i < 5; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestCutWhile(t *testing.T) {
	s := stream.N().Take(100).CutWhile(func(x interface{}) bool {
		return x.(int) >= 50
	})
	for i := 0; i < 50; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestMap(t *testing.T) {
	s := stream.N().Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	})
	for i := 0; i < 10; i++ {
		if s.Head() != 2*i {
			t.FailNow()
		}
		s = s.Tail()
	}

	s = stream.Make().Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	})
	if s != nil {
		t.FailNow()
	}
}

func TestFilter(t *testing.T) {
	s := stream.N().Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	for i := 0; i < 10; i++ {
		if s.Head() != 2*i {
			t.FailNow()
		}
		s = s.Tail()
	}

	s = stream.Make().Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	if s != nil {
		t.FailNow()
	}
}

func TestWalk(t *testing.T) {
	a := 0
	stream.N().Take(10).Walk(func(x interface{}) {
		a += x.(int)
	})
	if a != 45 {
		t.FailNow()
	}
}

func TestForce(t *testing.T) {
	var tail func() *stream.Stream
	x := 0
	tail = func() *stream.Stream {
		x += 1
		return stream.New(x, tail)
	}
	s := stream.New(0, tail).Take(10).Force()
	if x != 10 {
		t.FailNow()
	}
	x = 100
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.FailNow()
		}
		s = s.Tail()
	}
}

func TestFold(t *testing.T) {
	x := stream.N().Take(10).Fold(0, func(a interface{}, x interface{}) interface{} {
		return a.(int) + x.(int)
	})
	if x != 45 {
		t.FailNow()
	}
}

func TestAll(t *testing.T) {
	ok := stream.N().Take(10).All(func(x interface{}) bool {
		return x.(int) >= 0
	})
	if !ok {
		t.FailNow()
	}

	ok = stream.N().Take(10).All(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	if ok {
		t.FailNow()
	}
}

func TestAny(t *testing.T) {
	ok := stream.N().Take(10).Any(func(x interface{}) bool {
		return x == 5
	})
	if !ok {
		t.FailNow()
	}

	ok = stream.N().Take(10).Any(func(x interface{}) bool {
		return x == 100
	})
	if ok {
		t.FailNow()
	}
}

func TestCount(t *testing.T) {
	n := stream.N().Take(10).Count()
	if n != 10 {
		t.FailNow()
	}
}
