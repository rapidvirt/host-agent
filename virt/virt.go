package virt

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

// Connection to libvirt. This struct maintains a pointer to the
// libvirt connection
type Connection struct {
	conn *libvirt.Connect
}

// NewConnection establishes the libvirt connection for this server.
// Note that if the daemon is not available, this method returns an
// error.
func NewConnection() (*Connection, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil, fmt.Errorf("can't connect to libvirt: %v", err)
	}
	virtConn := &Connection{conn}
	return virtConn, nil
}

// Close the connection with libvirt
func (c *Connection) Close() (int, error) {
	return c.conn.Close()
}

// GetVersion from the libvirt daemon. This method returns an error
// if the daemon is not available or reachable
func (c *Connection) GetVersion() (uint32, error) {
	return c.conn.GetVersion()
}
