package vmachine

import (
	"fmt"
	"net/url"
)

type DataSource struct {
	host string
	port string
}

func NewDataSource(host string, port string) *DataSource {
	return &DataSource{
		host: host,
		port: port,
	}
}

func (d *DataSource) String() string {
	fullAddress := fmt.Sprintf("%s:%s", d.host, d.port)
	u := url.URL{
		Scheme: "http",
		Host:   fullAddress,
		Path:   "/__dmi.chassis-serial-number__/",
	}

	return fmt.Sprintf("ds=nocloud;s=%s", u.String())
}
