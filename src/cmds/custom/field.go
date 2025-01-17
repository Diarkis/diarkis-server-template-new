// © 2019-2024 Diarkis Inc. All rights reserved.

package customcmds

import (
	custom "github.com/Diarkis/diarkis-server-template/puffer/go/custom"
	"github.com/Diarkis/diarkis/field"
	"github.com/Diarkis/diarkis/server"
	"github.com/Diarkis/diarkis/user"
)

// add to matching and create a room
func getFieldInfo(ver uint8, cmd uint16, payload []byte, userData *user.User, next func(error)) {
	logger.Sys("Get Field Info Received from user: {}", userData.ID)
	fieldInfo := custom.NewGetFieldInfo()

	fieldInfo.NodeCount = int32(field.GetNodeNum())
	fieldInfo.FieldSize = field.GetFieldSize()
	fieldInfo.FieldOfVisionSize = field.GetFieldOfVisionSize()
	userData.ServerRespond(fieldInfo.Pack(), ver, cmd, server.Ok, true)
	next(nil)
}
