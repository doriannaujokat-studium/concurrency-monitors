package buff

type Buffer interface {
	Ins(byte)
	Num() uint
	Get() byte
	Full() bool
	Empty() bool
}

type buffer struct {
	data     []byte
	max_size uint
	index    uint
	size     uint
}

func new_(size uint) Buffer {
	x := new(buffer)
	x.data = make([]byte, size)
	x.max_size = size
	x.size = 0
	x.index = 0
	return x
}

func New(size uint) Buffer { return new_(size) }

func (x *buffer) Ins(b byte) {
	if x.size >= x.max_size {
		panic("buffer full")
	}
	x.data[(x.index+x.size)%(x.max_size)] = b
	x.size++
}

func (x *buffer) Num() uint {
	return x.size
}

func (x *buffer) Get() byte {
	if x.size <= 0 {
		panic("buffer empty")
	}
	var cb = x.data[x.index]
	x.size--
	x.index = (x.index + 1) % x.max_size
	return cb
}

func (x *buffer) Full() bool {
	return x.size >= x.max_size
}

func (x *buffer) Empty() bool {
	return x.size <= 0
}
