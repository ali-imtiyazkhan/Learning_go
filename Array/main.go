package main
import "fmt"

func main(){
    var booking [50]string 
	booking[0] = "John Doe"
	booking[1] = "Jane Doe"
	booking[2] = "Bob Smith"
	booking[49] ="Imtiyaz Khan"


    fmt.Println(booking)
	fmt.Printf("The size if array %T",len(booking))

	// problem with array

	// Alist is more dynamic in size 

}