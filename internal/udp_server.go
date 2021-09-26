package internal

import (
	"net"
	"sync"
)

type UDPServer interface {
	Start(wg *sync.WaitGroup, errors chan error)
	Stop() error
}

type SimpleUDPServer struct {
	address *net.UDPAddr
	udpConn net.UDPConn
}

func (u *SimpleUDPServer) Start(wg *sync.WaitGroup, errors chan error) {
	address, err := net.ResolveUDPAddr("udp", ":60789")
	if err != nil {
		wg.Done()
		errors <- err
		return
	}
	u.address = address
	LoggerS.Infof("Starting UDP server on %s", address)
	udp_conn, err := net.ListenUDP("udp", u.address)
	if err != nil {
		wg.Done()
		errors <- err
		return
	}
	u.udpConn = *udp_conn
	buffer := make([]byte, 1024)
	for {
		read, remoteAddress, err := u.udpConn.ReadFromUDP(buffer)
		if err != nil {
			LoggerS.Warn("Error on read UDP", err)
			break
		}
		LoggerS.Debugw("RECEVEID UDP",
			"remoteAddress", remoteAddress,
			"dataSize", read)
	}
	wg.Done()
}

func (u *SimpleUDPServer) Stop() error {
	return u.udpConn.Close()
}

func NewUDPServer() UDPServer {
	return &SimpleUDPServer{}
}
