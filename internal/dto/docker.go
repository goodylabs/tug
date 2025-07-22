package dto

type ServiceDTO struct {
	ID       string `json:"ID"`
	Image    string `json:"Image"`
	Mode     string `json:"Mode"`
	Name     string `json:"Name"`
	Ports    string `json:"Ports"`
	Replicas string `json:"Replicas"`
}

type ContainerDTO struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
}
