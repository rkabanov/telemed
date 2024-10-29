package main

type AppStore struct {
	*PatientStore
	*DoctorStore
}

func NewAppStore(ps *PatientStore, ds *DoctorStore) *AppStore {
	return &AppStore{
		PatientStore: ps,
		DoctorStore:  ds,
	}
}
