// © 2019-2024 Diarkis Inc. All rights reserved.

// Code generated by Diarkis Puffer module: DO NOT EDIT.
//
// # Auto-generated by Diarkis Version 1.0.0
//
// - Maximum length of a string is 65535 bytes
// - Maximum length of a byte array is 65535 bytes
// - Maximum length of any array is 65535 elements
package onlinestatus

import "encoding/binary"
import "errors"
import "fmt"
import "strings"

// OnlineStatusResponseVer represents the ver of the protocol's command.
//
//	[NOTE] The value is optional and if ver is not given in the definition JSON, it will be 0.
const OnlineStatusResponseVer uint8 = 2

// OnlineStatusResponseCmd represents the command ID of the protocol's command ID.
//
//	[NOTE] The value is optional and if cmd is not given in the definition JSON, it will be 0.
const OnlineStatusResponseCmd uint16 = 500

// OnlineStatusResponse represents the command protocol data structure.
type OnlineStatusResponse struct {
	// Command version of the protocol
	Ver uint8
	// Command ID of the protocol
	Cmd            uint16
	Status         uint16
	UserStatusList []*UserStatus
}

// NewOnlineStatusResponse creates a new instance of OnlineStatusResponse struct.
func NewOnlineStatusResponse() *OnlineStatusResponse {
	return &OnlineStatusResponse{Ver: 2, Cmd: 500, Status: 0, UserStatusList: make([]*UserStatus, 0)}
}

// Pack encodes OnlineStatusResponse struct to a byte array to be delivered over the command.
func (proto *OnlineStatusResponse) Pack() []byte {
	bytes := make([]byte, 0)

	/* uint16 */
	statusBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(statusBytes, uint16(proto.Status))
	bytes = append(bytes, statusBytes...)

	/* []UserStatus */
	userStatusListLengthBytes := make([]byte, 2)
	userStatusListLength := len(proto.UserStatusList)
	binary.BigEndian.PutUint16(userStatusListLengthBytes, uint16(userStatusListLength))
	bytes = append(bytes, userStatusListLengthBytes...)
	for i := 0; i < userStatusListLength; i++ {
		userStatusListSizeBytes := make([]byte, 2)
		userStatusListPacked := proto.UserStatusList[i].Pack()
		binary.BigEndian.PutUint16(userStatusListSizeBytes, uint16(len(userStatusListPacked)))
		bytes = append(bytes, userStatusListSizeBytes...)
		bytes = append(bytes, userStatusListPacked...)
	}

	// done
	return bytes
}

// Unpack decodes the command payload byte array to OnlineStatusResponse struct.
func (proto *OnlineStatusResponse) Unpack(bytes []byte) error {
	if len(bytes) < 4 {
		return errors.New("OnlineStatusResponseUnpackError")
	}

	offset := 0

	/* uint16 */
	proto.Status = binary.BigEndian.Uint16(bytes[offset : offset+2])
	offset += 2

	/* []UserStatus */
	userStatusListLength := int(binary.BigEndian.Uint16((bytes[offset : offset+2])))
	offset += 2
	proto.UserStatusList = make([]*UserStatus, userStatusListLength)
	for i := 0; i < userStatusListLength; i++ {
		userStatusListSize := int(binary.BigEndian.Uint16((bytes[offset : offset+2])))
		if userStatusListSize+offset > len(bytes) {
			return errors.New("UnpackError")
		}
		offset += 2
		userStatusListBytes := bytes[offset : offset+userStatusListSize]
		item := &UserStatus{Ver: 2, Cmd: 500}
		item.Unpack(userStatusListBytes)
		proto.UserStatusList[i] = item
		offset += userStatusListSize
	}

	return nil
}

func (proto *OnlineStatusResponse) String() string {
	list := make([]string, 0)
	list = append(list, fmt.Sprint("Status = ", proto.Status))
	for i, item := range proto.UserStatusList {
		list = append(list, fmt.Sprint("UserStatusList[", i, "] = ", "[", item.String(), "]"))
	}
	return strings.Join(list, " | ")
}

func (proto *OnlineStatusResponse) GetVer() uint8 {
	return proto.Ver
}
func (proto *OnlineStatusResponse) GetCmd() uint16 {
	return proto.Cmd
}
