package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    switch os.Args[1] {
    case "run":
        parent()
    case "child":
        child()
    default:
        panic("wat should I do")
    }
}

func parent() {

	for i,v:= range os.Args{
        fmt.Printf("parent index %d value %s\n",i,v)
    }
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[1:]...)...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }else{
		fmt.Println(" go parent ok .....")
		for{
		}
		fmt.Println(" go parent ok .....333")
	}
}

func child() { 
	for i,v:= range os.Args{
        fmt.Printf(" child  index %d value %s\n",i,v)
    }
	// child 执行完之后再去执行 parent else 部分
	// for{	}
}
