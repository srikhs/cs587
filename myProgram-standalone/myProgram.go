package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
	"math"
)

func New(x int64, y int64) MyType{
	return MyType{x,y}
}

func (data1 MyType) Distance(data2 MyType) float64{
	first := math.Pow(float64(data2.x - data1.x),2)
	second := math.Pow(float64(data2.y - data1.y),2)
	return math.Sqrt(first + second)
} 

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr, "Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


//	data    := MyType {}
//	data.x = 4
//	data.y = 2
//	data.x1 = 10
//	data.y1 = 4
	
	p1 := New(37, -76)
	p2 := New(26, -80)

	dist := p1.Distance(p2)
	efmt.Println(dist)

	efmt.Println("The input coordinates are") 
	fd, status := ethos.OpenDirectoryPath(path)
	o1.Write(fd)
	p1.WriteVar(path +"foobar")
//	efmt.Println(data)
//	ypoints := data.y1 - data.y
//	xpoints := data.x1 - data.x
	//slope := ypoints/xpoints
//	efmt.Println(data.y1 - data.y/data.x1 - data.x)
//	efmt.Println("saiiiiiiiiiii")
	
	


}
