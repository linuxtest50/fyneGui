package cookie

var cook *Cookie

type Cookie struct {
	send2service chan []byte
}

func NewCookie(ch chan []byte) *Cookie {
	cook = &Cookie{send2service: ch}
	return cook
}

func (c *Cookie) tellServer(msg []byte) {
	c.send2service <- msg
}

func TellServer(msgs [][]byte) {
	for _, msg := range msgs {
		cook.tellServer(msg)
	}
}
