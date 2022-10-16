package model

type PodList struct {
	PodGetlist []*PodGet
	Err               error
}

type PodGet struct {
	Name 	  string
	Namespace string
	ContainerNum  int
}
