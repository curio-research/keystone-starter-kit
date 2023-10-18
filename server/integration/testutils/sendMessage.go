package testutils

import (
	pb_dict "github.com/curio-research/keystone/game/proto/output/pb.dict"
	"github.com/curio-research/keystone/server"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func SendMessage(ws *websocket.Conn, cmd pb_dict.CMD, m proto.Message) error {
	networkMsg, err := server.NewMessage(0, uint32(cmd), 0, m)
	if err != nil {
		return err
	}

	buffer := networkMsg.ParseToBuffer()
	return ws.WriteMessage(websocket.TextMessage, buffer)
}
