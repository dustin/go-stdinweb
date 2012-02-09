package stdinweb

import (
	"errors"
	"net"
	"net/http"
	"os"
	"time"
)

type stdioListener bool
type stdioConn bool

var stategen = make(chan bool)

func init() {
	// Must send this async or it deadlocks
	go func() { stategen <- true }()
}

// Listener implementation

func (sl *stdioListener) Accept() (net.Conn, error) {
	// I can only accept once, but this is called in a relatively
	// tight loop.  I'm driving the two states with a simple
	// boolean channel.  The first accept is a real connection.
	// The next one blocks until that connection is done at which
	// point it returns an error and the server exits.
	if <-stategen {
		return new(stdioConn), nil
	}
	return nil, errors.New("Done")
}

func (sl *stdioListener) Close() error {
	return nil
}

func (sl *stdioListener) Addr() net.Addr {
	return &net.IPAddr{net.IPv4(0, 0, 0, 0)}
}

// Conn impl

func (sl *stdioConn) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (sl *stdioConn) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func (sl *stdioConn) LocalAddr() net.Addr {
	return &net.IPAddr{net.IPv4(0, 0, 0, 0)}
}

func (sl *stdioConn) RemoteAddr() net.Addr {
	return sl.LocalAddr()
}

func (sl *stdioConn) SetDeadline(t time.Time) error {
	return nil
}

func (sl *stdioConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (sl *stdioConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func (sl *stdioConn) Close() error {
	stategen <- false
	return nil
}

func ServeStdin(s http.Server) error {
	var sl stdioListener
	return s.Serve(&sl)
}
