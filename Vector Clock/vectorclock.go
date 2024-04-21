package main

import (
	"fmt"
)

// VectorClock 结构表示向量时钟
type VectorClock struct {
	Clock []int
}

// NewVectorClock 创建一个新的向量时钟实例
func NewVectorClock(numProcesses int) *VectorClock {
	clock := make([]int, numProcesses)
	return &VectorClock{
		Clock: clock,
	}
}

// Increment 增加指定进程的时钟值
func (vc *VectorClock) Increment(processID int) {
	vc.Clock[processID]++
}

// Update 根据接收到的其他进程的时钟更新本地时钟
func (vc *VectorClock) Update(receivedClock []int) {
	for i, value := range receivedClock {
		if value > vc.Clock[i] {
			vc.Clock[i] = value
		}
	}
}

// Compare 比较本地时钟与另一个时钟的关系
func (vc *VectorClock) Compare(otherClock []int) string {
	if equal(vc.Clock, otherClock) {
		return "concurrent"
	}
	if happenedBefore(vc.Clock, otherClock) {
		return "happened-before"
	}
	if happenedBefore(otherClock, vc.Clock) {
		return "happened-after"
	}
	return "concurrent"
}

// equal 检查两个时钟是否相等
func equal(clock1, clock2 []int) bool {
	for i, value := range clock1 {
		if value != clock2[i] {
			return false
		}
	}
	return true
}

// happenedBefore 检查时钟 clock1 是否发生在 clock2 之前
func happenedBefore(clock1, clock2 []int) bool {
	for i, value := range clock1 {
		if value > clock2[i] {
			return false
		}
	}
	return true
}

func main() {
	// 创建三个进程的向量时钟
	clock1 := NewVectorClock(3)
	clock2 := NewVectorClock(3)
	clock3 := NewVectorClock(3)

	// 对各个时钟进行增加操作
	clock1.Increment(0)
	clock1.Increment(1)
	clock2.Increment(1)
	clock2.Increment(2)
	clock3.Increment(0)
	clock3.Increment(2)

	// 输出各个时钟的状态
	fmt.Println("Clock 1:", clock1.Clock)
	fmt.Println("Clock 2:", clock2.Clock)
	fmt.Println("Clock 3:", clock3.Clock)
}
