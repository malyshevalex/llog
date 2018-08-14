package llog

import "time"

func itoa(b *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var tmp [20]byte
	bp := len(tmp) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		tmp[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	tmp[bp] = byte('0' + i)
	*b = append(*b, tmp[bp:]...)
}

func formatHeader(b *[]byte, l Level, t time.Time) {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	itoa(b, year, 4)
	*b = append(*b, '/')
	itoa(b, int(month), 2)
	*b = append(*b, '/')
	itoa(b, day, 2)
	*b = append(*b, ' ')

	itoa(b, hour, 2)
	*b = append(*b, ':')
	itoa(b, min, 2)
	*b = append(*b, ':')
	itoa(b, sec, 2)

	*b = append(*b, ' ', '[')
	*b = append(*b, levelPrefix[l]...)
	*b = append(*b, ']', ' ')
}
