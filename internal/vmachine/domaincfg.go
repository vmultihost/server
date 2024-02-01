package vmachine

import (
	"encoding/xml"
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

type Target struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type Interface struct {
	Type   string `xml:"type,attr"`
	Source Source `xml:"source"`
}

func createDomainCfg(name, memoryMiB, cpu string) (string, error) {
	dto := &Domain{
		Name: name,
		Type: "kvm",
		Memory: Memory{
			Unit: "MiB",
			Text: memoryMiB,
		},
		Vcpu: Vcpu{
			Placement: "static",
			Text:      cpu,
		},
	}
}
