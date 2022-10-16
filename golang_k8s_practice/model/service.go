package model

type ServiceList struct {
	Services []*ServiceGet
	Err               error
}


type ServiceGet struct {
	Name      string
	Type      string
	ClusterIp string
	Port 	  int32
}
