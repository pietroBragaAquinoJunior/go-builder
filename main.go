package main

import "fmt"

func main() {
	computerAppleBuilder := newComputerAppleBuilder()
	computerLinuxBuilder := newComputerLinuxBuilder()

	director := newDirector(computerAppleBuilder)
	computer := director.createComputer()

	fmt.Printf("Informações do computador, memory: %s, cpu: %s, video: %s", computer.memory, computer.cpu, computer.memory)

	director.setBuilder(computerLinuxBuilder)
	computer = director.createComputer()
	fmt.Printf("Informações do computador, memory: %s, cpu: %s, video: %s", computer.memory, computer.cpu, computer.memory)

}
