package main

type ComputerAppleBuilder struct {
	memory string
	cpu    string
	video  string
}

func newComputerAppleBuilder() *ComputerAppleBuilder {
	return &ComputerAppleBuilder{}
}

func (c *ComputerAppleBuilder) setMemory() {
	c.memory = "8GB"
}

func (c *ComputerAppleBuilder) setCpu() {
	c.cpu = "Intel"
}

func (c *ComputerAppleBuilder) setVideo() {
	c.video = "Ryzen X"
}

func (c *ComputerAppleBuilder) getComputer() Computer {
	return Computer{memory: c.memory, cpu: c.cpu, video: c.video}
}
