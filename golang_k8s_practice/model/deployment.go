package model

type DeploymentList struct {
	DeploymentGetlist []*DeploymentGet
	Err               error
}


type DeploymentGet struct {
	Name 	  string
	Namespace string
	Replicas  int32
}
