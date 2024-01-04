package domain

import "encoding/xml"

type Domain struct {
	XMLName xml.Name `xml:"domain"`
	Type    string   `xml:"type,attr"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	Uuid    string   `xml:"uuid"`
	Memory  struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	} `xml:"memory"`
	CurrentMemory struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	} `xml:"currentMemory"`
	Vcpu struct {
		Text      string `xml:",chardata"`
		Placement string `xml:"placement,attr"`
	} `xml:"vcpu"`
	Resource struct {
		Partition string `xml:"partition"`
	} `xml:"resource"`
	Os struct {
		Type struct {
			Text    string `xml:",chardata"`
			Arch    string `xml:"arch,attr"`
			Machine string `xml:"machine,attr"`
		} `xml:"type"`
		Boot struct {
			Dev string `xml:"dev,attr"`
		} `xml:"boot"`
	} `xml:"os"`
	Features struct {
		Acpi string `xml:"acpi"`
		Apic string `xml:"apic"`
	} `xml:"features"`
	Cpu struct {
		Mode  string `xml:"mode,attr"`
		Match string `xml:"match,attr"`
		Check string `xml:"check,attr"`
		Model struct {
			Text     string `xml:",chardata"`
			Fallback string `xml:"fallback,attr"`
		} `xml:"model"`
		Vendor  string `xml:"vendor"`
		Feature []struct {
			Policy string `xml:"policy,attr"`
			Name   string `xml:"name,attr"`
		} `xml:"feature"`
	} `xml:"cpu"`
	Clock struct {
		Offset string `xml:"offset,attr"`
		Timer  []struct {
			Name       string `xml:"name,attr"`
			Tickpolicy string `xml:"tickpolicy,attr"`
			Present    string `xml:"present,attr"`
		} `xml:"timer"`
	} `xml:"clock"`
	OnPoweroff string `xml:"on_poweroff"`
	OnReboot   string `xml:"on_reboot"`
	OnCrash    string `xml:"on_crash"`
	Pm         struct {
		SuspendToMem struct {
			Enabled string `xml:"enabled,attr"`
		} `xml:"suspend-to-mem"`
		SuspendToDisk struct {
			Enabled string `xml:"enabled,attr"`
		} `xml:"suspend-to-disk"`
	} `xml:"pm"`
	Devices struct {
		Emulator string `xml:"emulator"`
		Disk     []struct {
			Type   string `xml:"type,attr"`
			Device string `xml:"device,attr"`
			Driver struct {
				Name  string `xml:"name,attr"`
				Type  string `xml:"type,attr"`
				Cache string `xml:"cache,attr"`
				Io    string `xml:"io,attr"`
			} `xml:"driver"`
			Source struct {
				File string `xml:"file,attr"`
			} `xml:"source"`
			BackingStore string `xml:"backingStore"`
			Target       struct {
				Dev string `xml:"dev,attr"`
				Bus string `xml:"bus,attr"`
			} `xml:"target"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Address struct {
				Type       string `xml:"type,attr"`
				Domain     string `xml:"domain,attr"`
				Bus        string `xml:"bus,attr"`
				Slot       string `xml:"slot,attr"`
				Function   string `xml:"function,attr"`
				Controller string `xml:"controller,attr"`
				Target     string `xml:"target,attr"`
				Unit       string `xml:"unit,attr"`
			} `xml:"address"`
			Readonly string `xml:"readonly"`
		} `xml:"disk"`
		Controller []struct {
			Type  string `xml:"type,attr"`
			Index string `xml:"index,attr"`
			Model string `xml:"model,attr"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Address struct {
				Type          string `xml:"type,attr"`
				Domain        string `xml:"domain,attr"`
				Bus           string `xml:"bus,attr"`
				Slot          string `xml:"slot,attr"`
				Function      string `xml:"function,attr"`
				Multifunction string `xml:"multifunction,attr"`
			} `xml:"address"`
			Master struct {
				Startport string `xml:"startport,attr"`
			} `xml:"master"`
		} `xml:"controller"`
		Interface struct {
			Type string `xml:"type,attr"`
			Mac  struct {
				Address string `xml:"address,attr"`
			} `xml:"mac"`
			Source struct {
				Network string `xml:"network,attr"`
				Bridge  string `xml:"bridge,attr"`
			} `xml:"source"`
			Target struct {
				Dev string `xml:"dev,attr"`
			} `xml:"target"`
			Model struct {
				Type string `xml:"type,attr"`
			} `xml:"model"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Address struct {
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"interface"`
		Serial struct {
			Type   string `xml:"type,attr"`
			Source struct {
				Path string `xml:"path,attr"`
			} `xml:"source"`
			Target struct {
				Type  string `xml:"type,attr"`
				Port  string `xml:"port,attr"`
				Model struct {
					Name string `xml:"name,attr"`
				} `xml:"model"`
			} `xml:"target"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
		} `xml:"serial"`
		Console struct {
			Type   string `xml:"type,attr"`
			Tty    string `xml:"tty,attr"`
			Source struct {
				Path string `xml:"path,attr"`
			} `xml:"source"`
			Target struct {
				Type string `xml:"type,attr"`
				Port string `xml:"port,attr"`
			} `xml:"target"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
		} `xml:"console"`
		Input []struct {
			Type  string `xml:"type,attr"`
			Bus   string `xml:"bus,attr"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
		} `xml:"input"`
		Graphics struct {
			Type       string `xml:"type,attr"`
			Port       string `xml:"port,attr"`
			Autoport   string `xml:"autoport,attr"`
			AttrListen string `xml:"listen,attr"`
			Listen     struct {
				Type    string `xml:"type,attr"`
				Address string `xml:"address,attr"`
			} `xml:"listen"`
		} `xml:"graphics"`
		Video struct {
			Model struct {
				Type    string `xml:"type,attr"`
				Vram    string `xml:"vram,attr"`
				Heads   string `xml:"heads,attr"`
				Primary string `xml:"primary,attr"`
			} `xml:"model"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Address struct {
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"video"`
		Memballoon struct {
			Model string `xml:"model,attr"`
			Alias struct {
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Address struct {
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"memballoon"`
	} `xml:"devices"`
	Seclabel []struct {
		Type       string `xml:"type,attr"`
		Model      string `xml:"model,attr"`
		Relabel    string `xml:"relabel,attr"`
		Label      string `xml:"label"`
		Imagelabel string `xml:"imagelabel"`
	} `xml:"seclabel"`
}
