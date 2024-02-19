package vmachine

import (
	"encoding/xml"
	"errors"
	"strconv"
)

type domain struct {
	XMLName xml.Name `xml:"domain"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Memory  memory   `xml:"memory"`
	Vcpu    vcpu     `xml:"vcpu"`
	Os      osXml    `xml:"os"`
	Sysinfo sysinfo  `xml:"sysinfo"`
	Devices devices  `xml:"devices"`
}

type memory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}

type vcpu struct {
	Text      string `xml:",chardata"`
	Placement string `xml:"placement,attr"`
}

type osXml struct {
	Type   string `xml:"type"`
	Smbios smbios `xml:"smbios"`
}

type smbios struct {
	Mode string `xml:"mode,attr"`
}

type sysinfo struct {
	Type    string  `xml:"type,attr"`
	System  system  `xml:"system"`
	Chassis chassis `xml:"chassis"`
}

type system struct {
	Entry entry `xml:"entry"`
}

type entry struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type chassis struct {
	Entry entry `xml:"entry"`
}

type devices struct {
	Disk      disk         `xml:"disk"`
	Interface interfaceXml `xml:"interface"`
}

type disk struct {
	Type   string `xml:"type,attr"`
	Device string `xml:"device,attr"`
	Driver driver `xml:"driver"`
	Source source `xml:"source"`
	Target target `xml:"target"`
}

type driver struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Cache string `xml:"cache,attr"`
	Io    string `xml:"io,attr"`
}

type source struct {
	File string `xml:"file,attr"`
}

type sourceNet struct {
	Network string `xml:"network,attr"`
}

type target struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type interfaceXml struct {
	Type   string    `xml:"type,attr"`
	Source sourceNet `xml:"source"`
}

func NewDomain(
	instanceId string,
	dataSource *DataSource,
	hostName string,
	cpu uint64,
	memoryMiB uint64,
	imgPath string,
	network string,
) *domain {
	return &domain{
		Name: hostName,
		Type: "kvm",
		Memory: memory{
			Unit: "MiB",
			Text: strconv.FormatUint(memoryMiB, 10),
		},
		Vcpu: vcpu{
			Placement: "static",
			Text:      strconv.FormatUint(cpu, 10),
		},
		Os: osXml{
			Type: "hvm",
			Smbios: smbios{
				Mode: "sysinfo",
			},
		},
		Sysinfo: sysinfo{
			Type: "smbios",
			System: system{
				Entry: entry{
					Name: "serial",
					Text: dataSource.String(),
				},
			},
			Chassis: chassis{
				Entry: entry{
					Name: "serial",
					Text: instanceId,
				},
			},
		},
		Devices: devices{
			Disk: disk{
				Type:   "file",
				Device: "disk",
				Driver: driver{
					Type:  "qcow2",
					Name:  "qemu",
					Cache: "none",
					Io:    "native",
				},
				Source: source{
					File: imgPath,
				},
				Target: target{
					Dev: "vda",
					Bus: "virtio",
				},
			},
			Interface: interfaceXml{
				Type: "network",
				Source: sourceNet{
					Network: network,
				},
			},
		},
	}
}

func (d *domain) ToXml() (string, error) {
	xmlText, err := xml.MarshalIndent(d, " ", " ")
	if err != nil {
		return "", errors.New("failed to marshal domain")
	}

	return string(xmlText), nil
}
