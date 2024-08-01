package main

import (
	"fmt"
	"builder/computer"
)

func main() {
	computerAppleBuilder := computer.NewComputerAppleBuilder()
	computerLinuxBuilder := computer.NewComputerLinuxBuilder()

	director := computer.NewDirector(computerAppleBuilder)
	computer := director.CreateComputer()

	fmt.Printf("Informações do computador, memory: %s, cpu: %s, video: %s", computer.Memory, computer.Cpu, computer.Video + "\n")

	director.SetBuilder(computerLinuxBuilder)
	computer = director.CreateComputer()
	fmt.Printf("Informações do computador, memory: %s, cpu: %s, video: %s", computer.Memory, computer.Cpu, computer.Video + "\n")

}
