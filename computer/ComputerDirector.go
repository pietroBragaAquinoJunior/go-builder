package computer

type Director struct {
	builder IComputerBuilder
}

func NewDirector(builder IComputerBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) SetBuilder(builder IComputerBuilder) {
	d.builder = builder
}

func (d *Director) CreateComputer() ComputerStruct {
	d.builder.SetCpu()
	d.builder.SetMemory()
	d.builder.SetVideo()
	return d.builder.GetComputer()
}
