package vmachine

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type Domain struct {
	XMLName xml.Name `xml:"domain"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Memory  Memory   `xml:"memory"`
	Vcpu    Vcpu     `xml:"vcpu"`
	Os      Os       `xml:"os"`
	Sysinfo Sysinfo  `xml:"sysinfo"`
	Devices Devices  `xml:"devices"`
}

type Memory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}

type Vcpu struct {
	Text      string `xml:",chardata"`
	Placement string `xml:"placement,attr"`
}

type Os struct {
	Type   string `xml:"type"`
	Smbios smbios `xml:"smbios"`
}

type smbios struct {
	Mode string `xml:"mode,attr"`
}

type Sysinfo struct {
	Type    string  `xml:"type,attr"`
	System  System  `xml:"system"`
	Chassis Chassis `xml:"chassis"`
}

type System struct {
	Entry Entry `xml:"entry"`
}

type Entry struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type Chassis struct {
	Entry Entry `xml:"entry"`
}

type Devices struct {
	Disk      Disk      `xml:"disk"`
	Interface Interface `xml:"interface"`
}

type Disk struct {
	Type   string `xml:"type,attr"`
	Device string `xml:"device,attr"`
	Driver Driver `xml:"driver"`
	Source Source `xml:"source"`
	Target Target `xml:"target"`
}

type Driver struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Cache string `xml:"cache,attr"`
	Io    string `xml:"io,attr"`
}

type Source struct {
	File string `xml:"file,attr"`
}

type SourceNet struct {
	Network string `xml:"network,attr"`
}

type Target struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type Interface struct {
	Type   string    `xml:"type,attr"`
	Source SourceNet `xml:"source"`
}

func CreateDomainCfg(
	name string,
	memoryMiB uint64,
	cpu uint64,
	imgPath string,
	cloudInitUrl string,
	cloudInitId string,
	network string,
) (string, error) {
	cloudInitPath := fmt.Sprintf("ds=nocloud;s=%s/__dmi.chassis-serial-number__/", cloudInitUrl)

	dto := &Domain{
		Name: name,
		Type: "kvm",
		Memory: Memory{
			Unit: "MiB",
			Text: strconv.FormatUint(memoryMiB, 10),
		},
		Vcpu: Vcpu{
			Placement: "static",
			Text:      strconv.FormatUint(cpu, 10),
		},
		Os: Os{
			Type: "hvm",
			Smbios: smbios{
				Mode: "sysinfo",
			},
		},
		Sysinfo: Sysinfo{
			Type: "smbios",
			System: System{
				Entry: Entry{
					Name: "serial",
					Text: cloudInitPath,
				},
			},
			Chassis: Chassis{
				Entry: Entry{
					Name: "serial",
					Text: cloudInitId,
				},
			},
		},
		Devices: Devices{
			Disk: Disk{
				Type:   "file",
				Device: "disk",
				Driver: Driver{
					Type:  "qcow2",
					Name:  "qemu",
					Cache: "none",
					Io:    "native",
				},
				Source: Source{
					File: imgPath,
				},
				Target: Target{
					Dev: "vda",
					Bus: "virtio",
				},
			},
			Interface: Interface{
				Type: "network",
				Source: SourceNet{
					Network: network,
				},
			},
		},
	}

	xmlText, err := xml.MarshalIndent(dto, " ", " ")
	if err != nil {
		return "", errors.New("failed to marshal domain")
	}

	return string(xmlText), nil
}
