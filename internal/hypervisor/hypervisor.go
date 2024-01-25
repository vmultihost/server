package hypervisor

import (
	"fmt"
	"io"
	"os"

	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/sirupsen/logrus"
	"github.com/vmultihost/server/internal/hypervisor/xml_temp"
)

const (
	XML_VOL_PATH  = ""
	XML_POOL_PATH = ""
	POOL_NAME     = "pool1"
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

// todo: make async
func (h *Hypervisor) CopyImg(src string, volumeSizeKB uint64) error {
	if !h.virt.IsConnected() {
		return fmt.Errorf("not connected to hypervisor")
	}

	var err error
	var pool libvirt.StoragePool
	if pool, err = h.virt.StoragePoolLookupByName(POOL_NAME); err != nil {
		return err
	}

	var vol libvirt.StorageVol
	if vol, err = h.virt.StorageVolCreateXML(pool, xml_temp.VolXML, 0); err != nil {
		return err
	}
	h.log.Infof("volume is created %s", vol)

	var imgReader io.Reader
	if imgReader, err = os.Open(src); err != nil {
		return err
	}

	if err = h.virt.StorageVolUpload(vol, imgReader, 0, 0, libvirt.StorageVolUploadSparseStream); err != nil {
		return err
	}
	h.log.Info("volume is upladed")

	if err = h.virt.StorageVolResize(vol, volumeSizeKB, 0); err != nil {
		return err
	}
	h.log.Info("volume is resized")

	return nil
}
