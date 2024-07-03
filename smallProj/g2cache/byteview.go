package g2cache

// ByteView 用来表示缓存值
type ByteView struct {
	b []byte // b 将会存储真实的缓存值，选择 byte 类型是为了能够支持任意的数据类型的存储
}

// Len 返回字节切片的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回一个切片的copy，防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String 用于返回字符串，实现了 fmt.Stringer 接口
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes 返回一个切片的copy
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
