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

// GetVersion from the libvirt daemon and hypervisor. This method returns an error
// if the daemon is not available or reachable. It contains two keys, version for
// the current hypervisor version and libvirt for the libvirt library version.
func (c *Connection) GetVersion() (map[string]uint32, error) {
	libVersion, err := c.conn.GetLibVersion()
	if err != nil {
		return nil, err
	}
	version, err := c.conn.GetVersion()
	if err != nil {
		return nil, err
	}
	data := make(map[string]uint32)
	data["version"] = version
	data["libvirt"] = libVersion
	return data, nil
}
