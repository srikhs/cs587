package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr, "Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


	data    := MyType {}
	data.x = -3
	data.y = 5
	data.x1 = 2
	data.y1 = 1
	efmt.Println("The input coordinates are") 
	fd, status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
//	data.WriteVar(path +"foobar")
	efmt.Println(data)
	ypoints := data.y1 - data.y
	xpoints := data.x1 - data.x
	slope := ypoints/xpoints
	efmt.Println(slope)
//	efmt.Println("saiiiiiiiiiii")
}
