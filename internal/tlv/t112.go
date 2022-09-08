package tlv

import (
	"strconv"

	"github.com/shacha086/MiraiGo/binary"
)

func T112(uin int64) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x112)
		w.WriteStringShort(strconv.FormatInt(uin, 10))
	})
}
