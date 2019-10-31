package virt

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

// Connection to libvirt
type Connection struct {
	conn *libvirt.Connect
}

// NewConnection establish
func NewConnection() (*Connection, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil, fmt.Errorf("can't connect to libvirt: %v", err)
	}
	virtConn := &Connection{conn}
	return virtConn, nil
}
