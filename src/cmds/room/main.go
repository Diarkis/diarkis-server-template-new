package roomcmds

import (
	"encoding/binary"
	"github.com/Diarkis/diarkis/log"
	"github.com/Diarkis/diarkis/room"
	"github.com/Diarkis/diarkis/roomSupport"
	"strconv"
)

var logger = log.New("room")

func Expose() {
	room.ExposeCommands()
	roomSupport.ExposeCommands()
	// we broadcast user ID to room members when the client leaves unexpectedly
	// The message will be propagated by OnMemberLeave even on the client
	room.SetOnDiscardCustomMessage(onDiscardCustomMessage)
}

func onDiscardCustomMessage(roomID string, userID string) []byte {
	logger.Debug("OnDiscardCustomMessage roomID:%v userID:%v", roomID, userID)
	// UE4 sample client uses uin64 as the data type for userID
	conv, err := strconv.Atoi(userID)
	if err != nil {
		return []byte(userID)
	}
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(conv))
	sizeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBytes, uint32(len(bytes)))
	return append(sizeBytes, bytes...)
}
