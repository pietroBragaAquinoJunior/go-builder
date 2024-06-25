package main

type Director struct {
	builder IComputerBuilder
}

func newDirector(builder IComputerBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) setBuilder(builder IComputerBuilder) {
	d.builder = builder
}

func (d *Director) createComputer() Computer {
	d.builder.setCpu()
	d.builder.setMemory()
	d.builder.setVideo()
	return d.builder.getComputer()
}
