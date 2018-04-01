package stream

import (
	"math/rand"
	"time"
)

// 创建流
func Make(xs ...interface{}) *Stream {
	if len(xs) == 0 {
		return nil
	}
	return New(xs[0], func() *Stream {
		return Make(xs[1:]...)
	})
}

// 连接流
func Concat(s *Stream, ss ...*Stream) *Stream {
	if len(ss) == 0 {
		return s
	}
	if s == nil {
		return Concat(ss[0], ss[1:]...)
	}
	return New(s.Head(), func() *Stream {
		return Concat(s.Tail(), ss...)
	})
}

// 重复
func Repeat(x interface{}) *Stream {
	return New(x, func() *Stream {
		return Repeat(x)
	})
}

// 增长
func Grow(x int, s int) *Stream {
	return New(x, func() *Stream {
		return Grow(x+s, s)
	})
}

// 自然数
func N() *Stream {
	return Grow(0, 1)
}

// 范围
func Range(x int, y int) *Stream {
	if x < y {
		return Grow(x, 1).Take(y - x)
	} else {
		return Grow(x, -1).Take(x - y)
	}
}

// 随机
func Rand() *Stream {
	var tail func() *Stream
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	tail = func() *Stream {
		return New(rnd.Float64(), tail)
	}
	return New(rnd.Float64(), tail)
}
