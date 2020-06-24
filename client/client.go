package client

import (
	"bufio"
	"fyneGui/client_handler"
	"fyneGui/enmu"
	"fyneGui/helper/stack"
	"fyneGui/packet"
	"fyneGui/cookie"
	"github.com/golang/glog"
	"log"
	"net"
)

func ClientRun() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", enmu.ServerHost+":"+enmu.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	glog.Infof("Start Client Run" + enmu.ServerHost + ":" + enmu.ServerPort)
	handleRequest(conn)
}

func handleRequest(conn net.Conn) {
	in := make(chan []byte, 16)
	sess := cookie.NewCookie(in)
	defer func() {
		glog.Info("disconnect:" + conn.RemoteAddr().String())
		//sess.OffLine(sess.User.Id)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	go func() {
		for msg := range in {
			writer.Write(msg)
			writer.Write([]byte("\n"))
			writer.Flush()
		}
	}()
	for {
		msg, _, err := reader.ReadLine()
		if err != nil {
			glog.Info(err)
			return
		}
		reader := packet.Reader(msg)
		c, err := reader.ReadS16()
		if err != nil {
			glog.Info("err=", err)
			return
		}
		bytes := executeHandler(c, sess, reader)
		for _, byt := range bytes {
			in <- byt
		}
	}
}

//执行方法
func executeHandler(code int16, sess *cookie.Cookie, reader *packet.Packet) [][]byte {
	defer stack.PrintRecoverFromPanic()
	handle := client_handler.Handlers[code]
	if handle == nil {
		return nil
	}
	retByte := handle(sess, reader)
	return retByte
	return nil
}
