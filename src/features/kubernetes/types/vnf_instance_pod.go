package types

type VnfInstancePod struct {
	Id                  int
	Name                string
	Description         string
	Type                string
	Label               string
	VnfInfraName        string
	Discovered          bool
	ManagementInterface string
	ControlInterface    string
	Vendor              string
	Version             string
}
