package computer

type ComputerAppleBuilder struct {
	memory string
	cpu    string
	video  string
}

func NewComputerAppleBuilder() *ComputerAppleBuilder {
	return &ComputerAppleBuilder{}
}

func (c *ComputerAppleBuilder) SetMemory() {
	c.memory = "8GB"
}

func (c *ComputerAppleBuilder) SetCpu() {
	c.cpu = "Intel"
}

func (c *ComputerAppleBuilder) SetVideo() {
	c.video = "Ryzen X"
}

func (c *ComputerAppleBuilder) GetComputer() ComputerStruct {
	return ComputerStruct{Memory: c.memory, Cpu: c.cpu, Video: c.video}
}
