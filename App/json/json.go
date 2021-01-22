package json

type Status struct {
	Id       int    `json:"id"`
	HostName string `json:"hostname"`
	CPU      string `json:"cpu"`
	IP       string `json:"ip"`
	PORT     int    `json:"port"`
}

func json() {
}
