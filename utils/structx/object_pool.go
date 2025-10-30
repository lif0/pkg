package structx

type ObjectPool[T any] struct {
	s    uint32
	obj  []*T
	make func(uint32) []*T
}

func NewObjectPool[T any](size ...uint32) *ObjectPool[T] {
	var cap uint32
	if len(size) > 0 && size[0] > 0 {
		cap = size[0]
	}

	return &ObjectPool[T]{
		s:    cap,
		obj:  fmakeDefault[T](cap),
		make: fmakeDefault[T],
	}
}

func fmakeDefault[T any](n uint32) []*T {
	objs := make([]T, n)
	out := make([]*T, n)

	for i := range objs {
		out[i] = &objs[i]
	}
	return out
}

func (p *ObjectPool[T]) Get() *T {
	if len(p.obj) == 0 {
		p.allocObjects()
	}

	n := len(p.obj) - 1
	x := p.obj[n]
	p.obj = p.obj[:n]
	return x
}

// func (p *ObjectPool[T]) Put(x *T) {
// 	p.obj = append(p.obj, x)
// }

func (p *ObjectPool[T]) allocObjects() {
	const threshold = 512

	/* Можно сделать ускорение
	00:       0 -> 1       | +inf%
	01:       1 -> 2       | +100%
	02:       2 -> 4       | +100%
	03:       4 -> 8       | +100%
	04:       8 -> 16      | +100%
	05:      16 -> 32      | +100%
	06:      32 -> 64      | +100%
	07:      64 -> 128     | +100%
	08:     128 -> 256     | +100%
	09:     256 -> 512     | +100%
	10:     512 -> 848     | +66%
	11:     848 -> 1280    | +51%
	12:    1280 -> 1792    | +40%
	13:    1792 -> 2560    | +43%
	14:    2560 -> 3408    | +33%
	15:    3408 -> 5120    | +50%
	16:    5120 -> 7168    | +40%
	17:    7168 -> 9216    | +29%
	18:    9216 -> 12288   | +33%
	19:   12288 -> 16384   | +33%
	20:   16384 -> 21504   | +31%
	21:   21504 -> 27648   | +29%
	22:   27648 -> 34816   | +26%
	23:   34816 -> 44032   | +26%
	24:   44032 -> 55296   | +26%
	25:   55296 -> 69632   | +26%
	26:   69632 -> 88064   | +26%
	27:   88064 -> 110592  | +26%
	28:  110592 -> 139264  | +26%
	29:  139264 -> 175104  | +26%
	30:  175104 -> 219136  | +25%
	31:  219136 -> 274432  | +25%
	32:  274432 -> 344064  | +25%
	33:  344064 -> 431104  | +25%
	34:  431104 -> 539648  | +25%
	35:  539648 -> 674816  | +25%
	36:  674816 -> 843776  | +25%
	37:  843776 -> 1055744 | +25%
	*/
	cur := p.s

	var line uint32
	if cur < threshold {
		line = cur * 2
	} else {
		line = cur + (cur >> 2)
	}

	if line == 0 {
		line = 1
	}

	if uint32(cap(p.obj)) > line {
		line = uint32(cap(p.obj))
	}

	free := p.make(line)
	p.obj = append(p.obj, free...)

	// diff := uint32(cap(p.obj) - len(p.obj))
	// if diff > 1 {
	// 	free = p.make(diff)
	// 	p.obj = append(p.obj, free...)
	// 	line += diff
	// }

	p.s = line
}
