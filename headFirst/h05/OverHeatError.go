package main

import (
	"fmt"
	"log"
)

type OverHeatError float64

// Error 满足error接口
func (o OverHeatError) Error() string {
	// 在错误信息中使用温度
	return fmt.Sprintf("OverHeating by %0.2f degrees", o)
}

// String 实现了Stringer接口，允许任何类型在决定输出时如何展示
func (o OverHeatError) String() string {
	return "OK"
	//return strconv.FormatFloat(float64(o), 'E', -1, 32) + "degrees"
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverHeatError(excess)
	}
	return nil
}

func main() {
	over := OverHeatError(3.1415926)
	log.Println(over)

	// error类型像int或string一样是一个“预定义标识符”，它不属于任何包。它是“全局块”的一部分，这意味着它在任何地方可用，不用考虑当前包的信息
	err := checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
