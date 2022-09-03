package main

import "fmt"

func main() {

	point := 10

	if point == 10 {
		fmt.Println("Lulus dengan sempurna gan")
	} else if point > 5{
		fmt.Println("Lulus aja")
	} else if point == 4{
		fmt.Println("Dikit lagi lulus")
	} else{
		fmt.Printf("Sorry nih gan ente kaga lulus %d\n", point)
	}

	// * Temporary Variable

	pointBaru := 3340.0

	if percent := pointBaru / 100; percent >= 100{
		fmt.Printf("%.1f%s percent !\n", percent, "%")
	}else if percent >= 70{
		fmt.Printf("%.1f%s good\n", percent, "%")
	}else{
		fmt.Printf("%.1f%s not bad !\n", percent, "%")
	}

	// * Swtich-case

	point3 := 

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7,6,5,4:
		fmt.Println("Keren")
	default:
		{
			fmt.Println("its ok")
			fmt.Println("kapan kapan bisa lagi")
		}
	}

}