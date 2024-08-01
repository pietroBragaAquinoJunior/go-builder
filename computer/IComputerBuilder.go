package computer

type IComputerBuilder interface {
	SetMemory()
	SetCpu()
	SetVideo()
	GetComputer() ComputerStruct
}
