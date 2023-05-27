package networkPack

type Buffer struct {
	buf []int
	cap int
	len int
}

func (b *Buffer) CanAppend() bool {
	return b.len < b.cap
}

func (b *Buffer) Push(num int) {
	if b.CanAppend() {
		b.buf = append(b.buf, num)
		b.len++
	}
}

func (b *Buffer) Poll(num int) {
	if b.len == 0 {
		return
	}

	for {
		if value, ok := b.Top(); ok && value <= num {
			b.buf = b.buf[1:]
			b.len--
		} else {
			break
		}
	}
}

func (b *Buffer) Top() (int, bool) {
	if b.len == 0 {
		return 0, false
	}

	return b.buf[0], true
}

func NetworkPackProc(arrival []int, duration []int, bufSize int) []int {
	res := make([]int, 0)
	buffer := Buffer{buf: []int{}, cap: bufSize, len: 0}
	currentTime := 0
	for i := 0; i < len(arrival); i++ {
		buffer.Poll(arrival[i])

		if !buffer.CanAppend() {
			res = append(res, -1)
			continue
		}

		if currentTime < arrival[i] {
			currentTime = arrival[i]
		}

		res = append(res, currentTime)
		currentTime += duration[i]
		buffer.Push(currentTime)
	}

	return res
}
