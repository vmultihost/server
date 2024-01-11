package hypervisor

import (
	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/sirupsen/logrus"
)

type Hypervisor struct {
	config Config
	virt   *libvirt.Libvirt
	log    *logrus.Logger
}

func New(config Config, log *logrus.Logger) *Hypervisor {
	timeoutOption := dialers.WithLocalTimeout(config.socketTimeout)
	socketOption := dialers.WithSocket(config.socketName)
	dialer := dialers.NewLocal(timeoutOption, socketOption)
	virt := libvirt.NewWithDialer(dialer)

	return &Hypervisor{
		config: config,
		virt:   virt,
		log:    log,
	}
}
func (h *Hypervisor) Connect() error {
	if err := h.virt.Connect(); err != nil {
		return err
	}

	h.log.Info("connected to hypervisor")
	return nil
}

func (h *Hypervisor) Disconnect() error {
	if err := h.virt.Disconnect(); err != nil {
		return err
	}

	h.log.Info("disconnected from hypervisor")
	return nil
}
