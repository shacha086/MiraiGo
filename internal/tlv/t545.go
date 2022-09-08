package tlv

import "github.com/shacha086/MiraiGo/binary"

func T545(imei []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x545)
		w.WriteBytesShort(imei)
	})
}
