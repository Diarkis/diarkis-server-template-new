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
import "math"
import "strings"
import "github.com/Diarkis/diarkis/util"

// DiarkisQuaternionVer represents the ver of the protocol's command.
//
//	[NOTE] The value is optional and if ver is not given in the definition JSON, it will be 0.
const DiarkisQuaternionVer uint8 = 0

// DiarkisQuaternionCmd represents the command ID of the protocol's command ID.
//
//	[NOTE] The value is optional and if cmd is not given in the definition JSON, it will be 0.
const DiarkisQuaternionCmd uint16 = 0

// DiarkisQuaternion represents the command protocol data structure.
type DiarkisQuaternion struct {
	// Command version of the protocol
	Ver uint8
	// Command ID of the protocol
	Cmd uint16
	W float32
	X float32
	Y float32
	Z float32
}

// NewDiarkisQuaternion creates a new instance of DiarkisQuaternion struct.
func NewDiarkisQuaternion() *DiarkisQuaternion {
	return &DiarkisQuaternion{ Ver: 0, Cmd: 0, Y: 0, Z: 0, W: 0, X: 0 }
}

// Pack encodes DiarkisQuaternion struct to a byte array to be delivered over the command.
func (proto *DiarkisQuaternion) Pack() []byte {
	bytes := make([]byte, 0)

	/* float32 */
	wBytes := make([]byte, 4)
	wBits := math.Float32bits(proto.W)
	wBytes[0] = byte(wBits >> 24)
	wBytes[1] = byte(wBits >> 16)
	wBytes[2] = byte(wBits >> 8)
	wBytes[3] = byte(wBits)
	wBytes = util.ReverseBytes(wBytes)
	bytes = append(bytes, wBytes...)

	/* float32 */
	xBytes := make([]byte, 4)
	xBits := math.Float32bits(proto.X)
	xBytes[0] = byte(xBits >> 24)
	xBytes[1] = byte(xBits >> 16)
	xBytes[2] = byte(xBits >> 8)
	xBytes[3] = byte(xBits)
	xBytes = util.ReverseBytes(xBytes)
	bytes = append(bytes, xBytes...)

	/* float32 */
	yBytes := make([]byte, 4)
	yBits := math.Float32bits(proto.Y)
	yBytes[0] = byte(yBits >> 24)
	yBytes[1] = byte(yBits >> 16)
	yBytes[2] = byte(yBits >> 8)
	yBytes[3] = byte(yBits)
	yBytes = util.ReverseBytes(yBytes)
	bytes = append(bytes, yBytes...)

	/* float32 */
	zBytes := make([]byte, 4)
	zBits := math.Float32bits(proto.Z)
	zBytes[0] = byte(zBits >> 24)
	zBytes[1] = byte(zBits >> 16)
	zBytes[2] = byte(zBits >> 8)
	zBytes[3] = byte(zBits)
	zBytes = util.ReverseBytes(zBytes)
	bytes = append(bytes, zBytes...)

	// done
	return bytes
}

// Unpack decodes the command payload byte array to DiarkisQuaternion struct.
func (proto *DiarkisQuaternion) Unpack(bytes []byte) error {
	if len(bytes) < 16 {
		return errors.New("DiarkisQuaternionUnpackError")
	}

	offset := 0

	/* float32 */
	wBytes := util.ReverseBytes(bytes[offset:offset + 4])
	wBits := binary.BigEndian.Uint32(wBytes)
	offset += 4
	proto.W = math.Float32frombits(wBits)

	/* float32 */
	xBytes := util.ReverseBytes(bytes[offset:offset + 4])
	xBits := binary.BigEndian.Uint32(xBytes)
	offset += 4
	proto.X = math.Float32frombits(xBits)

	/* float32 */
	yBytes := util.ReverseBytes(bytes[offset:offset + 4])
	yBits := binary.BigEndian.Uint32(yBytes)
	offset += 4
	proto.Y = math.Float32frombits(yBits)

	/* float32 */
	zBytes := util.ReverseBytes(bytes[offset:offset + 4])
	zBits := binary.BigEndian.Uint32(zBytes)
	offset += 4
	proto.Z = math.Float32frombits(zBits)


	return nil
}

func (proto *DiarkisQuaternion) String() string {
	list := make([]string, 0)
	list = append(list, fmt.Sprint("W = ", proto.W))
	list = append(list, fmt.Sprint("X = ", proto.X))
	list = append(list, fmt.Sprint("Y = ", proto.Y))
	list = append(list, fmt.Sprint("Z = ", proto.Z))
	return strings.Join(list, " | ")
}

func (proto *DiarkisQuaternion) GetVer() uint8 {
	return proto.Ver
}
func (proto *DiarkisQuaternion) GetCmd() uint16 {
	return proto.Cmd
}
