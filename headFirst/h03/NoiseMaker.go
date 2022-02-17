package main

import "fmt"

type Robot string

// MakeSound 实现了NoiseMaker接口的方法
func (r Robot) MakeSound() {
	fmt.Println("Beep Bop")
}

// Walk 自有的方法
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	// 通过类型断言取回具体类型的值，可以调用那个类型上的方法，但这个方法并不属于接口
	robot, ok := noiseMaker.(Robot)
	if ok {
		robot.Walk()
	} else {
		fmt.Println("It's not Robot")
	}
}
