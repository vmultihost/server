package vmachine

import (
	"github.com/sirupsen/logrus"
	"github.com/vmultihost/server/internal/cloudinit"
)

type Vmachine struct {
	dataSource string
	imgPath    string
	network    string
	log        *logrus.Logger
}

type VmConfig struct {
	name         string
	memoryMiB    uint64
	cpu          uint64
	cloudInitUrl string
	cloudInitId  string
	network      string
}

func New(
	dataSource string,
	imgPath string,
	network string,
	log *logrus.Logger,
) *Vmachine {
	return &Vmachine{
		dataSource: dataSource,
		imgPath:    imgPath,
		network:    network,
		log:        log,
	}
}

func (v *Vmachine) Create(
	instanceId string,
	name string,
	memoryMiB uint64,
	cpu uint64,
) error {
	domain, err := CreateDomainCfg(
		instanceId,
		v.dataSource,
		name,
		memoryMiB,
		cpu,
		v.imgPath,
		v.network,
	)
	if err != nil {
		return err
	}

	v.log.Info(domain)
	return nil
}

func (v *Vmachine) Configure(cloudInit *cloudinit.CloudInit) error {

	return nil
}
