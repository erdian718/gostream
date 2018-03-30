package stream

// 流
type Stream struct {
	head  interface{}
	tail  func() *Stream
	cache *Stream
	count int
}

// 未知
var unknown = new(Stream)

// 创建流
func New(head interface{}, tail func() *Stream) *Stream {
	return &Stream{
		head:  head,
		tail:  tail,
		cache: unknown,
		count: -1,
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

// 将元素加到流的头部
func (self *Stream) Cons(x interface{}) *Stream {
	return New(x, func() *Stream {
		return self
	})
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
	if self == nil || n <= 0 {
		return self
	}
	return self.Tail().Drop(n - 1)
}

// 丢弃前面的元素
func (self *Stream) DropWhile(f func(interface{}) bool) *Stream {
	if self == nil {
		return nil
	}
	if h := self.Head(); f(h) {
		return self.Tail().DropWhile(f)
	} else {
		return self
	}
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
	if self == nil {
		return nil
	}
	f(self.Head())
	self.Tail().Walk(f)
	return self
}

// 折叠
func (self *Stream) Fold(x interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	if self == nil {
		return x
	}
	return self.Tail().Fold(f(x, self.Head()), f)
}

// 所有元素满足条件
func (self *Stream) All(f func(interface{}) bool) bool {
	if self == nil {
		return true
	}
	if h := self.Head(); f(h) {
		return self.Tail().All(f)
	} else {
		return false
	}
}

// 任意一个元素满足条件
func (self *Stream) Any(f func(interface{}) bool) bool {
	if self == nil {
		return false
	}
	if h := self.Head(); f(h) {
		return true
	} else {
		return self.Tail().Any(f)
	}
}

// 元素个数
func (self *Stream) Count() int {
	if self == nil {
		return 0
	}
	if self.count < 0 {
		self.count = 1 + self.Tail().Count()
	}
	return self.count
}
