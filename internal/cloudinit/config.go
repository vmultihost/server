// config for cloud-init data source and http server
package cloudinit

import (
	"fmt"
	"net/url"
)

type HttpConfig struct {
	instanceId string
	host       string
	port       uint16
}

func NewHttpConfig(instanceId, host string, port uint16) *HttpConfig {
	return &HttpConfig{
		// todo: remove
		instanceId: instanceId,
		host:       host,
		port:       port,
	}
}

func (c *HttpConfig) InstanceId() string {
	return c.instanceId
}

func (c *HttpConfig) DataSource() string {
	host := fmt.Sprintf("%s:%d", c.host, c.port)
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/__dmi.chassis-serial-number__/",
	}

	return fmt.Sprintf("ds=nocloud;s=%s", u.String())
}
