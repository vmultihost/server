package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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
	ciHost        = "localhost"
	ciPort        = "12345"
	logLevel      = "info"
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	fmt.Println("Hello from server!")

	log := logrus.New()
	vmTest(log)

	flag.Parse()

	config, err := httpserver.NewConfigFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	server := httpserver.NewExample(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

func vmTest(log *logrus.Logger) {
	serverConfig := httpserver.NewConfig(ciHost, ciPort, logLevel)

	// todo: create cloudInit inside cloudinit.NewServer()? where create log & set logLevel?
	store := cloudinit.NewStore(log)
	cloudInit := cloudinit.NewCloudInit(store, log)

	cloudInitServer := cloudinit.NewServer(serverConfig, cloudInit)

	go func() {
		if err := cloudInitServer.Start(); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	hypervisorConfig := hypervisor.NewConfig(socketName, socketTimeout, imgPath)
	httpCloudInit := vmachine.NewHttpCloudInit(ciHost, ciPort, cloudInit, log)
	hv := hypervisor.New(hypervisorConfig, httpCloudInit, log)

	// if err := hv.Connect(); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	// if err := hv.CopyImg(imgSource, volumeSizeKB); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	// if err := hv.Disconnect(); err != nil {
	// 	log.Error(err)
	// 	os.Exit(1)
	// }

	log.Info("create vm")
	vm := hv.CreateVm(
		"vm1",
		"network1",
		2,
		2000,
		"user",
		"password",
		[]string{"rsa list"},
	)

	if err := vm.Create(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err := vm.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
