package unix_socket

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"syscall"

	"code.cloudfoundry.org/garden-linux/container_daemon"
)

type Listener struct {
	listener   net.Listener
	socketFile *os.File
}

type Response struct {
	ErrMessage string
	Pid        int
}

// This function should be called from the host namespace, to open the socket file in the right file system.
func NewListenerFromPath(socketPath string) (*Listener, error) {
	l := &Listener{}
	var err error

	l.listener, err = net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("unix_socket: error creating socket: %v", err)
	}

	return l, nil
}

func NewListenerFromFile(socketFile *os.File) (*Listener, error) {
	l := &Listener{}

	var err error
	l.listener, err = net.FileListener(socketFile)
	if err != nil {
		return nil, fmt.Errorf("unix_socket: error creating listener: %v", err)
	}

	return l, nil
}

func (l *Listener) Listen(ch container_daemon.ConnectionHandler) error {
	if l.listener == nil {
		return errors.New("unix_socket: listener is not initialized")
	}

	var conn net.Conn
	var err error
	for {
		conn, err = l.listener.Accept()
		if err != nil {
			return fmt.Errorf("container_daemon: Failure while accepting: %v", err)
		}

		go func(conn *net.UnixConn, ch container_daemon.ConnectionHandler) {
			defer conn.Close() // Ignore error

			decoder := json.NewDecoder(conn)

			response, err := ch.Handle(decoder)
			if err != nil {
				response = &container_daemon.ResponseMessage{ErrMessage: err.Error()}
			}
			writeData(conn, response)
		}(conn.(*net.UnixConn), ch)
	}
}

func writeData(conn *net.UnixConn, response *container_daemon.ResponseMessage) {
	data, _ := json.Marshal(response) // Ignore error

	args := make([]int, len(response.Files))
	for i, f := range response.Files {
		args[i] = int(f.Fd())
	}
	oobData := syscall.UnixRights(args...)

	conn.WriteMsgUnix(data, oobData, nil) // Ignore error

	// Close the files whose descriptors have been sent to the host to ensure that
	// a close on the host takes effect in a timely fashion.
	for _, file := range response.Files {
		file.Close() // Ignore error
	}
}

func (l *Listener) File() (*os.File, error) {
	unixListener, ok := l.listener.(*net.UnixListener)
	if !ok {
		return nil, fmt.Errorf("unix_socket: incorrect listener type: %v", l.listener)
	}
	return unixListener.File()
}

func (l *Listener) Close() error {
	return l.listener.Close()
}
