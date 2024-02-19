package vmachine

import (
	"github.com/digitalocean/go-libvirt"
	"github.com/sirupsen/logrus"
)

type Vmachine struct {
	domain *domain
	virt   *libvirt.Libvirt
	log    *logrus.Logger
}

func New(
	domain *domain,
	virt *libvirt.Libvirt,
	log *logrus.Logger,
) *Vmachine {
	return &Vmachine{
		domain: domain,
		virt:   virt,
		log:    log,
	}
}

func (v *Vmachine) Create() error {
	domainXml, err := v.domain.ToXml()

	if err != nil {
		return err
	}

	// todo: create vm using libvirt
	v.log.Info(domainXml)
	return nil
}

func (v *Vmachine) Start() error {
	v.log.Info("")
	return nil
}
