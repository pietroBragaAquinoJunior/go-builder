package main

type IComputerBuilder interface {
	setMemory()
	setCpu()
	setVideo()
	getComputer() Computer
}
