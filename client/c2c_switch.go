// Code generated by internal/generator/c2c_switcher DO NOT EDIT.

package client

import (
	"github.com/shacha086/MiraiGo/client/internal/network"
	"github.com/shacha086/MiraiGo/client/pb/msg"
)

const (
	UnknownDecoder = iota
	nonSvcNotifyTroopSystemMsgDecoders
	otherDecoders
	privateMsgDecoders
	sysMsgDecoders
	troopSystemMsgDecoders
)

func peekC2CDecoder(msgType int32) (decoder func(*QQClient, *msg.Message, *network.IncomingPacketInfo), decoderType uint8) {
	switch msgType {
	case 9:
		return privateMessageDecoder, privateMsgDecoders
	case 10:
		return privateMessageDecoder, privateMsgDecoders
	case 31:
		return privateMessageDecoder, privateMsgDecoders
	case 33:
		return troopAddMemberBroadcastDecoder, otherDecoders
	case 35:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 36:
		return troopSystemMessageDecoder, nonSvcNotifyTroopSystemMsgDecoders
	case 37:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 45:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 46:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 79:
		return privateMessageDecoder, privateMsgDecoders
	case 84:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 85:
		return troopSystemMessageDecoder, nonSvcNotifyTroopSystemMsgDecoders
	case 86:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 87:
		return troopSystemMessageDecoder, troopSystemMsgDecoders
	case 97:
		return privateMessageDecoder, privateMsgDecoders
	case 120:
		return privateMessageDecoder, privateMsgDecoders
	case 132:
		return privateMessageDecoder, privateMsgDecoders
	case 133:
		return privateMessageDecoder, privateMsgDecoders
	case 140:
		return tempSessionDecoder, privateMsgDecoders
	case 141:
		return tempSessionDecoder, privateMsgDecoders
	case 166:
		return privateMessageDecoder, privateMsgDecoders
	case 167:
		return privateMessageDecoder, privateMsgDecoders
	case 187:
		return systemMessageDecoder, sysMsgDecoders
	case 188:
		return systemMessageDecoder, sysMsgDecoders
	case 189:
		return systemMessageDecoder, sysMsgDecoders
	case 190:
		return systemMessageDecoder, sysMsgDecoders
	case 191:
		return systemMessageDecoder, sysMsgDecoders
	case 208:
		return privatePttDecoder, privateMsgDecoders
	case 529:
		return msgType0x211Decoder, otherDecoders
	default:
		return nil, UnknownDecoder
	}
}
