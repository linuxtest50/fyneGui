package client_handler

import (
	"fyneGui/client_proto"
	"fyneGui/cookie"
	"fyneGui/packet"
)

func Login(name string) {
	cookie.TellServer([][]byte{packet.Pack(Code["user_login_req"], client_proto.S_entity_id{name}, nil)})
}
