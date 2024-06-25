package main

type ComputerLinuxBuilder struct {
	memory string
	video  string
	cpu    string
}

func newComputerLinuxBuilder() *ComputerLinuxBuilder {
	return &ComputerLinuxBuilder{}
}

func (c *ComputerLinuxBuilder) setMemory() {
	c.memory = "16GB"
}

func (c *ComputerLinuxBuilder) setVideo() {
	c.video = "NVIDIA RTX 3060"
}

func (c *ComputerLinuxBuilder) setCpu() {
	c.cpu = "intel i7 11ºGen"
}

func (c *ComputerLinuxBuilder) getComputer() Computer {
	return Computer{
		cpu:    c.cpu,
		memory: c.memory,
		video:  c.video,
	}
}
