// Code generated by Diarkis Puffer module: DO NOT EDIT.
//
// Auto-generated by Diarkis Version 1.0.0
//
// - Maximum length of a string is 65535 bytes
// - Maximum length of a byte array is 65535 bytes
// - Maximum length of any array is 65535 elements
package custom

import "encoding/binary"
import "errors"
import "fmt"
import "strings"
import "server_bina/custom"

// DiarkisCharacterSyncPayloadVer represents the ver of the protocol's command.
//
//	[NOTE] The value is optional and if ver is not given in the definition JSON, it will be 0.
const DiarkisCharacterSyncPayloadVer uint8 = 0

// DiarkisCharacterSyncPayloadCmd represents the command ID of the protocol's command ID.
//
//	[NOTE] The value is optional and if cmd is not given in the definition JSON, it will be 0.
const DiarkisCharacterSyncPayloadCmd uint16 = 0

// DiarkisCharacterSyncPayload represents the command protocol data structure.
type DiarkisCharacterSyncPayload struct {
	// Command version of the protocol
	Ver uint8
	// Command ID of the protocol
	Cmd uint16
	Frames []*custom.DiarkisCharacterFrameData
	PacketGUID string
	UID string
}

// NewDiarkisCharacterSyncPayload creates a new instance of DiarkisCharacterSyncPayload struct.
func NewDiarkisCharacterSyncPayload() *DiarkisCharacterSyncPayload {
	return &DiarkisCharacterSyncPayload{ Ver: 0, Cmd: 0, Frames: make([]*diarkischaracterframedata.DiarkisCharacterFrameData, 0), PacketGUID: "", UID: "" }
}

// Pack encodes DiarkisCharacterSyncPayload struct to a byte array to be delivered over the command.
func (proto *DiarkisCharacterSyncPayload) Pack() []byte {
	bytes := make([]byte, 0)

	/* []custom.DiarkisCharacterFrameData */
	framesLengthBytes := make([]byte, 2)
	framesLength := len(proto.Frames)
	binary.BigEndian.PutUint16(framesLengthBytes, uint16(framesLength))
	bytes = append(bytes, framesLengthBytes...)
	for i := 0; i < framesLength; i++ {
		framesSizeBytes := make([]byte, 2)
		framesPacked := proto.Frames[i].Pack()
		binary.BigEndian.PutUint16(framesSizeBytes, uint16(len(framesPacked)))
		bytes = append(bytes, framesSizeBytes...)
		bytes = append(bytes, framesPacked...)
	}

	/* string */
	packetGUIDSizeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(packetGUIDSizeBytes, uint16(len(proto.PacketGUID)))
	bytes = append(bytes, packetGUIDSizeBytes...)
	bytes = append(bytes, []byte(proto.PacketGUID)...)

	/* string */
	uIDSizeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(uIDSizeBytes, uint16(len(proto.UID)))
	bytes = append(bytes, uIDSizeBytes...)
	bytes = append(bytes, []byte(proto.UID)...)

	// done
	return bytes
}

// Unpack decodes the command payload byte array to DiarkisCharacterSyncPayload struct.
func (proto *DiarkisCharacterSyncPayload) Unpack(bytes []byte) error {
	if len(bytes) < 6 {
		return errors.New("DiarkisCharacterSyncPayloadUnpackError")
	}

	offset := 0

	/* []custom.DiarkisCharacterFrameData */
	framesLength := int(binary.BigEndian.Uint16((bytes[offset:offset + 2])))
	offset += 2
	proto.Frames = make([]*custom.DiarkisCharacterFrameData, framesLength)
	for i := 0; i < framesLength; i++ {
		framesSize := int(binary.BigEndian.Uint16((bytes[offset:offset + 2])))
		if framesSize + offset > len(bytes) {
			return errors.New("UnpackError")
		}
		offset += 2
		framesBytes := bytes[offset:offset + framesSize]
		item := &custom.DiarkisCharacterFrameData{ Ver: 0, Cmd: 0 }
		item.Unpack(framesBytes)
		proto.Frames[i] = item
		offset += framesSize
	}

	/* string */
	packetGUIDSize := int(binary.BigEndian.Uint16(bytes[offset:offset + 2]))
	if packetGUIDSize + offset > len(bytes) {
		return errors.New("UnpackError")
	}
	offset += 2
	proto.PacketGUID = string(bytes[offset:offset + packetGUIDSize])
	offset += packetGUIDSize

	/* string */
	uIDSize := int(binary.BigEndian.Uint16(bytes[offset:offset + 2]))
	if uIDSize + offset > len(bytes) {
		return errors.New("UnpackError")
	}
	offset += 2
	proto.UID = string(bytes[offset:offset + uIDSize])
	offset += uIDSize


	return nil
}

func (proto *DiarkisCharacterSyncPayload) String() string {
	list := make([]string, 0)
	for i, item := range proto.Frames { list = append(list, fmt.Sprint("Frames[", i, "] = ", "[", item.String(), "]")) }
	list = append(list, fmt.Sprint("PacketGUID = ", proto.PacketGUID))
	list = append(list, fmt.Sprint("UID = ", proto.UID))
	return strings.Join(list, " | ")
}

func (proto *DiarkisCharacterSyncPayload) GetVer() uint8 {
	return proto.Ver
}
func (proto *DiarkisCharacterSyncPayload) GetCmd() uint16 {
	return proto.Cmd
}
