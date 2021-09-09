package switchcase

import (
	"fmt"
	"runtime"
	"time"
)

func Switch() {
	exemplo()
	exemplo1()
	exemplo2()
	exemplo3()
	exemplo4()
}

func exemplo() {
	fmt.Print("[Switch] Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func exemplo1() {
	i := 2
	fmt.Print("[Switch] Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("[Switch] one")
	case 2:
		fmt.Println("[Switch] two")
	case 3:
		fmt.Println("[Switch] three")
	}
}

func exemplo2() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("[Switch] It's the weekend")
	default:
		fmt.Println("[Switch] It's a weekday")
	}
}

func exemplo3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("[Switch] It's before noon")
	default:
		fmt.Println("[Switch] It's after noon")
	}
}

// Type Switch
func exemplo4() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("[Switch] I'm a bool")
		case int:
			fmt.Println("[Switch] I'm an int")
		default:
			fmt.Printf("[Switch] Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("[Switch] hey")
}
