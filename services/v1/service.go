package v1

import (
	"encoding/json"
	"net/http"

	"github.com/rapidvirt/host-agent/virt"
)

// Service base struct
type Service struct {
	Conn *virt.Connection
}

// Version of the libvirt daemon and hypervisor. This service can return an error
// if there are any problem with the libvirt daemon.
func (s *Service) Version(w http.ResponseWriter, r *http.Request) {
	version, err := s.Conn.GetVersion()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	content, err := json.Marshal(version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
