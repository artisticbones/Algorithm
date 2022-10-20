package bitmap

type BitMap struct {
	bits []byte
	vmax uint
}

func NewBitMap(maxVal ...uint) *BitMap {
	var max uint = 8192
	if len(maxVal) > 0 && maxVal[0] > 0 {
		max = maxVal[0]
	}
	bm := &BitMap{}
	bm.vmax = max
	sz := (max + 7) / 8
	bm.bits = make([]byte, sz)
	return bm
}

func (bm *BitMap) Set(num uint) {
	if num > bm.vmax {
		bm.vmax += 1024
		if bm.vmax < num {
			bm.vmax = num
		}
		dd := int(num+7)/8 - len(bm.bits)
		if dd > 0 {
			tmpArr := make([]byte, 0, dd)
			bm.bits = append(bm.bits, tmpArr...)
		}
	}
	bm.bits[num/8] |= 1 << (num % 8)
}

func (bm *BitMap) Unset(num uint) {
	if num > bm.vmax {
		return
	}
	bm.bits[num/8] &^= 1 << (num % 8)
}

func (bm *BitMap) Check(num uint) bool {
	if num > bm.vmax {
		return false
	}
	return bm.bits[num/8]&(1<<(num%8)) != 0
}
