package vmachine

import "github.com/sirupsen/logrus"

type Vmachine struct {
	log *logrus.Logger
}

type VmConfig struct {
	name         string
	memoryMiB    uint64
	cpu          uint64
	cloudInitUrl string
	cloudInitId  string
	network      string
}

func New(log *logrus.Logger) *Vmachine {
	return &Vmachine{
		log: log,
	}
}

func (v *Vmachine) Create() error {

	return nil
}
