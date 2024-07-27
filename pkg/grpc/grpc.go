package grpc

import "net"

func NewListenerConn(port string) net.Conn {
	l, _ := net.Listen("tcp", "127.0.0.1:"+port)
	c, _ := l.Accept()

	defer l.Close()
	defer c.Close()

	return c
}
