package cloudinit

import "github.com/sirupsen/logrus"

type Store struct {
	metaDataStore map[string]*metaDataDto
	userDataStore map[string]*userDataDto
	log           *logrus.Logger
}

func NewStore(log *logrus.Logger) *Store {
	return &Store{
		metaDataStore: map[string]*metaDataDto{},
		userDataStore: map[string]*userDataDto{},
		log:           log,
	}
}

func (s *Store) AddMetaData(instanceId string, data *metaDataDto) {
	s.metaDataStore[instanceId] = data
}

func (s *Store) AddUserData(instanceId string, data *userDataDto) {
	s.userDataStore[instanceId] = data
}

func (s *Store) GetMetaData(instanceId string) *metaDataDto {
	data, ok := s.metaDataStore[instanceId]
	if !ok {
		s.log.Errorf("not found meta-data by %s", instanceId)
		return &metaDataDto{}
	}

	return data
}

func (s *Store) GetUserData(instanceId string) *userDataDto {
	data, ok := s.userDataStore[instanceId]
	if !ok {
		s.log.Errorf("not found user-data by %s", instanceId)
		return &userDataDto{}
	}

	return data
}
