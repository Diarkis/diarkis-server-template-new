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

// GetFieldInfoVer represents the ver of the protocol's command.
//
//	[NOTE] The value is optional and if ver is not given in the definition JSON, it will be 0.
const GetFieldInfoVer uint8 = 2

// GetFieldInfoCmd represents the command ID of the protocol's command ID.
//
//	[NOTE] The value is optional and if cmd is not given in the definition JSON, it will be 0.
const GetFieldInfoCmd uint16 = 4001

// GetFieldInfo represents the command protocol data structure.
type GetFieldInfo struct {
	// Command version of the protocol
	Ver uint8
	// Command ID of the protocol
	Cmd uint16
	FieldOfVisionSize int64
	FieldSize int64
	NodeCount int32
}

// NewGetFieldInfo creates a new instance of GetFieldInfo struct.
func NewGetFieldInfo() *GetFieldInfo {
	return &GetFieldInfo{ Ver: 2, Cmd: 4001, FieldSize: 0, NodeCount: 0, FieldOfVisionSize: 0 }
}

// Pack encodes GetFieldInfo struct to a byte array to be delivered over the command.
func (proto *GetFieldInfo) Pack() []byte {
	bytes := make([]byte, 0)

	/* int64 */
	fieldOfVisionSizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(fieldOfVisionSizeBytes, uint64(proto.FieldOfVisionSize))
	bytes = append(bytes, fieldOfVisionSizeBytes...)

	/* int64 */
	fieldSizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(fieldSizeBytes, uint64(proto.FieldSize))
	bytes = append(bytes, fieldSizeBytes...)

	/* int32 */
	nodeCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nodeCountBytes, uint32(proto.NodeCount))
	bytes = append(bytes, nodeCountBytes...)

	// done
	return bytes
}

// Unpack decodes the command payload byte array to GetFieldInfo struct.
func (proto *GetFieldInfo) Unpack(bytes []byte) error {
	if len(bytes) < 20 {
		return errors.New("GetFieldInfoUnpackError")
	}

	offset := 0

	/* int64 */
	proto.FieldOfVisionSize = int64(binary.BigEndian.Uint64(bytes[offset:offset + 8]))
	offset += 8

	/* int64 */
	proto.FieldSize = int64(binary.BigEndian.Uint64(bytes[offset:offset + 8]))
	offset += 8

	/* int32 */
	proto.NodeCount = int32(binary.BigEndian.Uint32(bytes[offset:offset + 4]))
	offset += 4


	return nil
}

func (proto *GetFieldInfo) String() string {
	list := make([]string, 0)
	list = append(list, fmt.Sprint("FieldOfVisionSize = ", proto.FieldOfVisionSize))
	list = append(list, fmt.Sprint("FieldSize = ", proto.FieldSize))
	list = append(list, fmt.Sprint("NodeCount = ", proto.NodeCount))
	return strings.Join(list, " | ")
}

func (proto *GetFieldInfo) GetVer() uint8 {
	return proto.Ver
}
func (proto *GetFieldInfo) GetCmd() uint16 {
	return proto.Cmd
}
