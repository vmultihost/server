package vmachine

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type cloudInit struct {
	instanceId  string
	hostName    string
	userName    string
	password    string
	sshAuthKeys []string
	log         *logrus.Logger
}

type metaDataDto struct {
	InstanceId    string `yaml:"instance-id,omitempty"`
	LocalHostname string `yaml:"local-hostname,omitempty"`
}

type userDataDto struct {
	Users []userDto `yaml:"users,omitempty"`
}

type userDto struct {
	Name              string   `yaml:"name,omitempty"`
	PlainTextPasswd   string   `yaml:"plain_text_passwd,omitempty"`
	SshAuthorizedKeys []string `yaml:"ssh_authorized_keys,omitempty"`
	Sudo              []string `yaml:"sudo,omitempty,flow"`
	Groups            string   `yaml:"groups,omitempty"`
	Shell             string   `yaml:"shell,omitempty"`
}

func NewCloudInit(hostName, userName, password string, sshAuthKeys []string, log *logrus.Logger) *cloudInit {
	return &cloudInit{
		instanceId:  uuid.New().String(),
		hostName:    hostName,
		userName:    userName,
		password:    password,
		sshAuthKeys: sshAuthKeys,
		log:         log,
	}
}

func (c *cloudInit) GetMetaDataYaml() (string, error) {
	metaData := &metaDataDto{
		InstanceId:    c.instanceId,
		LocalHostname: c.hostName,
	}

	data, err := yaml.Marshal(metaData)
	if err != nil {
		return "", errors.New("failed to marshal metaData")
	}

	return string(data), nil
}

func (c *cloudInit) GetUserDataYaml() (string, error) {
	userData := &userDataDto{
		Users: []userDto{
			{
				Name:              c.userName,
				PlainTextPasswd:   c.password,
				SshAuthorizedKeys: c.sshAuthKeys,
				Sudo:              []string{"ALL=(ALL) NOPASSWD:ALL"},
				Groups:            "sudo",
				Shell:             "/bin/bash",
			},
		},
	}

	data, err := yaml.Marshal(userData)
	if err != nil {
		return "", errors.New("failed to marshal metaData")
	}

	result := "#cloud-config\n" + string(data)
	return result, nil
}
