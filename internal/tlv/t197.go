package tlv

import "github.com/shacha086/MiraiGo/binary"

func T197() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x197)
		w.Write([]byte{0, 1, 0}) // w.WriteBytesShort([]byte{0})
	})
}
