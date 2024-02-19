package vmachine

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/vmultihost/server/internal/cloudinit"
)

type HttpCloudInit struct {
	host      string
	port      string
	cloudInit *cloudinit.CloudInit
	log       *logrus.Logger
}

func NewHttpCloudInit(
	host string,
	port string,
	cloudInit *cloudinit.CloudInit,
	log *logrus.Logger,
) *HttpCloudInit {
	return &HttpCloudInit{
		host:      host,
		port:      port,
		cloudInit: cloudInit,
		log:       log,
	}
}

func (c *HttpCloudInit) DataSource() *DataSource {
	return NewDataSource(c.host, c.port)
}

func (c *HttpCloudInit) AddVmConfig(
	hostName string,
	userName string,
	password string,
	sshAuthKeys []string,
) string {
	instanceId := uuid.New().String()

	c.cloudInit.AddMetaData(instanceId, hostName)
	c.cloudInit.AddUserData(instanceId, userName, password, sshAuthKeys)

	return instanceId
}
