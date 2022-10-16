package model

type NamespaceList struct {
	Namespaces []*NamespaceGet
	Err               error
}


type NamespacePhase string

type NamespaceGet struct {
	Name      string
	Status	  NamespacePhase
}
