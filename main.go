package main 

import "fmt"
import "github.com/lillesalt/rivercrossing/event"
import "bufio"
import "os"

/*
func main (){
	fmt.Println(state.ViewState())
	fmt.Println(event.PutInn("Kylling"))
	fmt.Println(event.CrossRiver("Kylling"))
	fmt.Println(event.TaUt("Kylling"))
}
*/

func main (){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Kva vil du sende over?")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(event.FirstPut(text))
		fmt.Println(event.GetInn(text))
		fmt.Println(event.CrossRiver(text))
		fmt.Println(event.TaUt(text))
		fmt.Println(event.GetOut(text))
		if(text == "Kylling"){
			break
		}
	}
}
