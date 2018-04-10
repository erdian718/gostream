package stream_test

import (
	"testing"

	"ofunc/stream"
)

func BenchmarkTake(b *testing.B) {
	stream.N().Take(b.N).Force()
}

func BenchmarkTakeWhile(b *testing.B) {
	stream.N().TakeWhile(func(x interface{}) bool {
		return x.(int) < b.N
	}).Force()
}

func BenchmarkDrop(b *testing.B) {
	stream.N().Drop(b.N)
}

func BenchmarkDropWhile(b *testing.B) {
	stream.N().DropWhile(func(x interface{}) bool {
		return x.(int) < b.N
	})
}

func BenchmarkCut(b *testing.B) {
	stream.N().Take(b.N).Cut(b.N / 2).Force()
}

func BenchmarkMap(b *testing.B) {
	stream.N().Map(func(x interface{}) interface{} {
		return 2 * x.(int)
	}).Take(b.N).Force()
}

func BenchmarkFilter(b *testing.B) {
	stream.N().Filter(func(x interface{}) bool {
		return x.(int)%2 == 0
	}).Take(b.N).Force()
}

func BenchmarkWalk(b *testing.B) {
	a := 0
	stream.N().Take(b.N).Walk(func(x interface{}) {
		a += x.(int)
	})
}

func BenchmarkForce(b *testing.B) {
	stream.N().Take(b.N).Force()
}

func BenchmarkFold(b *testing.B) {
	stream.N().Take(b.N).Fold(0, func(a interface{}, x interface{}) interface{} {
		return a.(int) + x.(int)
	})
}

func BenchmarkAll(b *testing.B) {
	stream.N().Take(b.N).All(func(x interface{}) bool {
		return x.(int) >= 0
	})
}

func BenchmarkAny(b *testing.B) {
	stream.N().Take(b.N).Any(func(x interface{}) bool {
		return x.(int) < 0
	})
}

func BenchmarkCount(b *testing.B) {
	n := stream.N().Take(b.N).Count()
	if n != b.N {
		b.FailNow()
	}
}

func BenchmarkConcat(b *testing.B) {
	stream.Concat(stream.N().Take(b.N), stream.Range(-b.N, 0)).Force()
}
