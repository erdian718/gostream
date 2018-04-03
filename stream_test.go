package stream_test

import (
	"testing"

	"ofunc/stream"
)

func TestTake(t *testing.T) {
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

func TestTakeWhile(t *testing.T) {
	s := stream.N().TakeWhile(func(x interface{}) bool {
		return x.(int) < 10
	})
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
	s = s.TakeWhile(func(x interface{}) bool {
		return x == 0
	})
	if s != nil {
		t.Fail()
	}
}

func TestDrop(t *testing.T) {
	x := stream.N().Drop(10).Head()
	if x != 10 {
		t.Fail()
	}
}

func TestDropWhile(t *testing.T) {
	x := stream.N().DropWhile(func(x interface{}) bool {
		return x.(int) < 10
	}).Head()
	if x != 10 {
		t.Fail()
	}
}

func TestCut(t *testing.T) {
	s := stream.N().Cut(0).Take(10).Cut(5)
	for i := 0; i < 5; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
	if s != nil {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	s := stream.N().Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	})
	for i := 0; i < 10; i++ {
		if s.Head() != 2*i {
			t.Fail()
		}
		s = s.Tail()
	}

	s = stream.Make().Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	})
	if s != nil {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	s := stream.N().Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	for i := 0; i < 10; i++ {
		if s.Head() != 2*i {
			t.Fail()
		}
		s = s.Tail()
	}

	s = stream.Make().Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	if s != nil {
		t.Fail()
	}
}

func TestWalk(t *testing.T) {
	a := 0
	stream.N().Take(10).Walk(func(x interface{}) {
		a += x.(int)
	})
	if a != 45 {
		t.Fail()
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
		t.Fail()
	}
	x = 100
	for i := 0; i < 10; i++ {
		if s.Head() != i {
			t.Fail()
		}
		s = s.Tail()
	}
}

func TestFold(t *testing.T) {
	x := stream.N().Take(10).Fold(0, func(a interface{}, x interface{}) interface{} {
		return a.(int) + x.(int)
	})
	if x != 45 {
		t.Fail()
	}
}

func TestAll(t *testing.T) {
	ok := stream.N().Take(10).All(func(x interface{}) bool {
		return x.(int) >= 0
	})
	if !ok {
		t.Fail()
	}

	ok = stream.N().Take(10).All(func(x interface{}) bool {
		return x.(int)%2 == 0
	})
	if ok {
		t.Fail()
	}
}

func TestAny(t *testing.T) {
	ok := stream.N().Take(10).Any(func(x interface{}) bool {
		return x == 5
	})
	if !ok {
		t.Fail()
	}

	ok = stream.N().Take(10).Any(func(x interface{}) bool {
		return x == 100
	})
	if ok {
		t.Fail()
	}
}
