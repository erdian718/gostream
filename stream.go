package stream

// 流
type Stream struct {
	head  interface{}
	tail  func() *Stream
	cache *Stream
}

// 未知
var unknown = new(Stream)

// 创建流
func New(head interface{}, tail func() *Stream) *Stream {
	return &Stream{
		head:  head,
		tail:  tail,
		cache: unknown,
	}
}

// 头
func (self *Stream) Head() interface{} {
	return self.head
}

// 尾
func (self *Stream) Tail() *Stream {
	if self.cache == unknown {
		self.cache = self.tail()
	}
	return self.cache
}

// 取前n个元素
func (self *Stream) Take(n int) *Stream {
	if self == nil || n <= 0 {
		return nil
	}
	return New(self.Head(), func() *Stream {
		return self.Tail().Take(n - 1)
	})
}

// 取前面的元素
func (self *Stream) TakeWhile(f func(interface{}) bool) *Stream {
	if self == nil {
		return nil
	}
	if h := self.Head(); f(h) {
		return New(h, func() *Stream {
			return self.Tail().TakeWhile(f)
		})
	} else {
		return nil
	}
}

// 丢弃前n个元素
func (self *Stream) Drop(n int) *Stream {
	s := self
	for i := 0; i < n && s != nil; i++ {
		s = s.Tail()
	}
	return s
}

// 丢弃前面的元素
func (self *Stream) DropWhile(f func(interface{}) bool) *Stream {
	s := self
	for s != nil {
		if h := s.Head(); f(h) {
			s = s.Tail()
		} else {
			break
		}
	}
	return s
}

// 去掉尾部n个元素
func (self *Stream) Cut(n int) *Stream {
	if n <= 0 {
		return self
	}
	return cut(self, self.Drop(n))
}

// 去掉尾部的元素
func (self *Stream) CutWhile(f func(interface{}) bool) *Stream {
	return cutWhile(self, self, f)
}

// 映射
func (self *Stream) Map(f func(interface{}) interface{}) *Stream {
	if self == nil {
		return nil
	}
	return New(f(self.Head()), func() *Stream {
		return self.Tail().Map(f)
	})
}

// 过滤
func (self *Stream) Filter(f func(interface{}) bool) *Stream {
	if self == nil {
		return nil
	}
	if h := self.Head(); f(h) {
		return New(h, func() *Stream {
			return self.Tail().Filter(f)
		})
	} else {
		return self.Tail().Filter(f)
	}
}

// 遍历
func (self *Stream) Walk(f func(interface{})) *Stream {
	for s := self; s != nil; s = s.Tail() {
		f(s.Head())
	}
	return self
}

// 强制求值
func (self *Stream) Force() *Stream {
	for s := self; s != nil; s = s.Tail() {
	}
	return self
}

// 折叠
func (self *Stream) Fold(a interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	for s := self; s != nil; s = s.Tail() {
		a = f(a, s.Head())
	}
	return a
}

// 所有元素满足条件
func (self *Stream) All(f func(interface{}) bool) bool {
	for s := self; s != nil; s = s.Tail() {
		if !f(s.Head()) {
			return false
		}
	}
	return true
}

// 任意一个元素满足条件
func (self *Stream) Any(f func(interface{}) bool) bool {
	for s := self; s != nil; s = s.Tail() {
		if f(s.Head()) {
			return true
		}
	}
	return false
}

// 个数
func (self *Stream) Count() int {
	n := 0
	for s := self; s != nil; s = s.Tail() {
		n += 1
	}
	return n
}

func cut(xs, ys *Stream) *Stream {
	if ys == nil {
		return nil
	}
	return New(xs.Head(), func() *Stream {
		return cut(xs.Tail(), ys.Tail())
	})
}

func cutWhile(xs, ys *Stream, f func(interface{}) bool) *Stream {
	if xs == ys {
		ys = ys.DropWhile(f)
	}
	if ys == nil {
		return nil
	}
	return New(xs.Head(), func() *Stream {
		if ys == xs {
			ys = ys.Tail()
		}
		return cutWhile(xs.Tail(), ys, f)
	})
}
