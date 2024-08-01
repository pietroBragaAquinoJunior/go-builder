package computer

type ComputerLinuxBuilder struct {
	memory string
	video  string
	cpu    string
}

func NewComputerLinuxBuilder() *ComputerLinuxBuilder {
	return &ComputerLinuxBuilder{}
}

func (c *ComputerLinuxBuilder) SetMemory() {
	c.memory = "16GB"
}

func (c *ComputerLinuxBuilder) SetVideo() {
	c.video = "NVIDIA RTX 3060"
}

func (c *ComputerLinuxBuilder) SetCpu() {
	c.cpu = "intel i7 11ºGen"
}

func (c *ComputerLinuxBuilder) GetComputer() ComputerStruct {
	return ComputerStruct{
		Cpu:    c.cpu,
		Memory: c.memory,
		Video:  c.video,
	}
}
