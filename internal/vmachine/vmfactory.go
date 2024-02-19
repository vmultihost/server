package vmachine

import (
	"github.com/digitalocean/go-libvirt"
	"github.com/sirupsen/logrus"
)

type VmFactory struct {
	imgPath    string
	dataSource *DataSource
	virt       *libvirt.Libvirt
	log        *logrus.Logger
}

func NewVmFactory(
	imgPath string,
	dataSource *DataSource,
	virt *libvirt.Libvirt,
	log *logrus.Logger,
) *VmFactory {
	return &VmFactory{
		imgPath:    imgPath,
		dataSource: dataSource,
		virt:       virt,
		log:        log,
	}
}

func (f *VmFactory) CreateVm(
	instanceId string,
	hostName string,
	network string,
	cpu uint64,
	memoryMiB uint64,
) *Vmachine {
	domain := NewDomain(
		instanceId,
		f.dataSource,
		hostName,
		cpu,
		memoryMiB,
		f.imgPath,
		network,
	)

	return New(domain, f.virt, f.log)
}
