package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/vmultihost/server/internal/cloudinit"
	"github.com/vmultihost/server/internal/httpserver"
	"github.com/vmultihost/server/internal/hypervisor"
	"github.com/vmultihost/server/internal/vmachine"
)

var (
	configPath string
)

const (
	socketName    = "/var/run/libvirt/libvirt-sock"
	socketTimeout = 15 * time.Second
	imgPath       = "/var/lib/libvirt/images"
	imgSource     = "/home/kostuyn/Downloads/ubuntu-22.04-server-cloudimg-amd64.img"
	volumeSizeKB  = 10 * 1024 * 1024 * 1024
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	fmt.Println("Hello from server!")

	log := logrus.New()
	vmTest(log)
	cfg := cloudinit.NewHttpConfig(uuid.NewString(), "localhost", 8080)
	// yamlTest(cfg, log)
	domainXmlTest(cfg, log)

	flag.Parse()

	config := httpserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server := httpserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	// domain := new(domain.Domain)
	// err := xml.Unmarshal([]byte(data), domain)
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	return
	// }

	// xmlText, err := xml.MarshalIndent(domain, " ", " ")
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// }
	// fmt.Printf("%s\n", string(xmlText))
}

func domainXmlTest(config *cloudinit.HttpConfig, log *logrus.Logger) {
	xml, err := vmachine.CreateDomainCfg(
		"instanceId",
		"data_source",
		"server_name",
		2000,
		2,
		"/var/lib/libvirt/images/vol1",
		"network1",
	)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Println(xml)
}

func yamlTest(config *cloudinit.HttpConfig, log *logrus.Logger) {
	cloudInit := cloudinit.NewCloudInit(
		"instanceId",
		"server1",
		"vmuser",
		"p@ssword",
		[]string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDZ2uffyx6IblR6dbmdtFdrqdUmH410HW7TsMSnVzsgi9wAjGheG0mYCw+dQHQzmmO+ZIDoIQuF4TE4d1PUdHzErgSkF6PNf/1Hq5+ycDqbg9qWvfsnpfkrZ8ZXkac1cEwvUwP8+aknpUMvdd1Tb1KJGvGbNFgjRWWdmB7QeobweY+6/SoV/c9n0lWRLjzlr/xXwRNgq924DrKVQPXQnoD2UFX/K8QTDaROJ7kh4x/hooPrsp9scfazwMQ+g9FS9isDXnR7HMiuJ2R2LFy4AS3E8Nh4g+3ywYcFMhL2W6ZSVWpTpKHh2x8ac+ZXxl6KTD2HdkPoHoc6Tkr4o4kexpO/ k.altuhov@MacBook-Pro-admin-6.local"},
		log,
	)

	metaDataStr, err := cloudInit.GetMetaDataYaml()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	userDataStr, err := cloudInit.GetUserDataYaml()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info(metaDataStr)
	log.Info(userDataStr)
}

func vmTest(log *logrus.Logger) {
	hypervisorConfig := hypervisor.NewConfig(socketName, socketTimeout, imgPath)
	hypervisor := hypervisor.New(hypervisorConfig, log)

	// if err := hypervisor.Connect(); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	// if err := hypervisor.CopyImg(imgSource, volumeSizeKB); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	// if err := hypervisor.Disconnect(); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	cloudInitServer := cloudinit.NewServer(
		"localhost",
		8081,
		log,
	)

	go func() {
		if err := cloudInitServer.Start(); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	log.Info("create vm")
	instanceId := "instanceId"

	vm := vmachine.New(
		cloudInitServer.DataSource(),
		hypervisor.ImgPath(),
		// todo: get from hypervisor
		"network1",
		log,
	)

	cloudInit := cloudinit.NewCloudInit(
		instanceId,
		"vm1",
		"user",
		"password",
		[]string{"rsa list"},
		log,
	)

	cloudInitServer.AddCloudInit(cloudInit)

	vm.Create(
		instanceId,
		"vm1",
		2000,
		2,
	)

	// vm.Start()
}

var data = `
<domain type='kvm' id='9'>
  <name>server1</name>
  <uuid>b995e3e1-0e57-47c8-b5c2-d046e92cbf3c</uuid>
  <memory unit='KiB'>2097152</memory>
  <currentMemory unit='KiB'>2097152</currentMemory>
  <vcpu placement='static'>2</vcpu>
  <resource>
    <partition>/machine</partition>
  </resource>
  <os>
    <type arch='x86_64' machine='pc-i440fx-bionic'>hvm</type>
    <boot dev='hd'/>
  </os>
  <features>
    <acpi/>
    <apic/>
  </features>
  <cpu mode='custom' match='exact' check='full'>
    <model fallback='forbid'>IvyBridge-IBRS</model>
    <vendor>Intel</vendor>
    <feature policy='require' name='ss'/>
    <feature policy='require' name='vmx'/>
    <feature policy='require' name='pcid'/>
    <feature policy='require' name='hypervisor'/>
    <feature policy='require' name='arat'/>
    <feature policy='require' name='tsc_adjust'/>
    <feature policy='require' name='md-clear'/>
    <feature policy='require' name='ssbd'/>
    <feature policy='require' name='xsaveopt'/>
    <feature policy='require' name='ibpb'/>
    <feature policy='require' name='amd-ssbd'/>
  </cpu>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='none' io='native'/>
      <source file='/home/kostuyn/qemu/pool/server1.qcow2'/>
      <backingStore/>
      <target dev='vda' bus='virtio'/>
      <alias name='virtio-disk0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </disk>
    <disk type='file' device='cdrom'>
      <target dev='hda' bus='ide'/>
      <readonly/>
      <alias name='ide0-0-0'/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='ich9-ehci1'>
      <alias name='usb'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x7'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci1'>
      <alias name='usb'/>
      <master startport='0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0' multifunction='on'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci2'>
      <alias name='usb'/>
      <master startport='2'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x1'/>
    </controller>
    <controller type='usb' index='0' model='ich9-uhci3'>
      <alias name='usb'/>
      <master startport='4'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'>
      <alias name='pci.0'/>
    </controller>
    <controller type='ide' index='0'>
      <alias name='ide'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <interface type='network'>
      <mac address='52:54:00:06:7a:5c'/>
      <source network='network1' bridge='virbr1'/>
      <target dev='vnet0'/>
      <model type='rtl8139'/>
      <alias name='net0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <source path='/dev/pts/1'/>
      <target type='isa-serial' port='0'>
        <model name='isa-serial'/>
      </target>
      <alias name='serial0'/>
    </serial>
    <console type='pty' tty='/dev/pts/1'>
      <source path='/dev/pts/1'/>
      <target type='serial' port='0'/>
      <alias name='serial0'/>
    </console>
    <input type='mouse' bus='ps2'>
      <alias name='input0'/>
    </input>
    <input type='keyboard' bus='ps2'>
      <alias name='input1'/>
    </input>
    <graphics type='vnc' port='5900' autoport='yes' listen='127.0.0.1'>
      <listen type='address' address='127.0.0.1'/>
    </graphics>
    <video>
      <model type='cirrus' vram='16384' heads='1' primary='yes'/>
      <alias name='video0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <alias name='balloon0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
  <seclabel type='dynamic' model='apparmor' relabel='yes'>
    <label>libvirt-b995e3e1-0e57-47c8-b5c2-d046e92cbf3c</label>
    <imagelabel>libvirt-b995e3e1-0e57-47c8-b5c2-d046e92cbf3c</imagelabel>
  </seclabel>
  <seclabel type='dynamic' model='dac' relabel='yes'>
    <label>+64055:+132</label>
    <imagelabel>+64055:+132</imagelabel>
  </seclabel>
</domain>
`
