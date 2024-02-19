// creats cloud-init meta-data and user-data
package cloudinit

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type CloudInit struct {
	store *Store
	log   *logrus.Logger
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

func NewCloudInit(
	store *Store,
	log *logrus.Logger,
) *CloudInit {
	return &CloudInit{
		store: store,
		log:   log,
	}
}

// todo: remove ?
func (c *CloudInit) AddVmConfig(
	hostName string,
	userName string,
	password string,
	sshAuthKeys []string,
) string {
	instanceId := uuid.New().String()

	metaData := &metaDataDto{
		InstanceId:    instanceId,
		LocalHostname: hostName,
	}

	userData := &userDataDto{
		Users: []userDto{
			{
				Name:              userName,
				PlainTextPasswd:   password,
				SshAuthorizedKeys: sshAuthKeys,
				Sudo:              []string{"ALL=(ALL) NOPASSWD:ALL"},
				Groups:            "sudo",
				Shell:             "/bin/bash",
			},
		},
	}

	c.store.AddMetaData(instanceId, metaData)
	c.store.AddUserData(instanceId, userData)

	return instanceId
}

func (c *CloudInit) AddMetaData(
	instanceId string,
	hostName string,
) string {
	metaData := &metaDataDto{
		InstanceId:    instanceId,
		LocalHostname: hostName,
	}

	c.store.AddMetaData(instanceId, metaData)

	return instanceId
}

func (c *CloudInit) AddUserData(
	instanceId string,
	userName string,
	password string,
	sshAuthKeys []string,
) string {
	userData := &userDataDto{
		Users: []userDto{
			{
				Name:              userName,
				PlainTextPasswd:   password,
				SshAuthorizedKeys: sshAuthKeys,
				Sudo:              []string{"ALL=(ALL) NOPASSWD:ALL"},
				Groups:            "sudo",
				Shell:             "/bin/bash",
			},
		},
	}

	c.store.AddUserData(instanceId, userData)

	return instanceId
}

func (c *CloudInit) GetMetaDataYaml(instanceId string) (string, error) {
	metaData := c.store.GetMetaData(instanceId)

	data, err := yaml.Marshal(metaData)
	if err != nil {
		c.log.Errorf("failed to marshal meta-data for instanceId %s", instanceId)
		return "", errors.New("marshal failed")
	}

	return string(data), nil
}

func (c *CloudInit) GetUserDataYaml(instanceId string) (string, error) {
	userData := c.store.GetUserData(instanceId)

	data, err := yaml.Marshal(userData)
	if err != nil {
		c.log.Errorf("failed to marshal user-data for instanceId %s", instanceId)
		return "", errors.New("marshal failed")
	}

	result := "#cloud-config\n" + string(data)
	return result, nil
}
