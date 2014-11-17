package main
import "fmt"
import "math"
func main(){
	fmt.Printf("Escape Velocity\n")
	G := 6.67 * math.Pow(10, -11) 
	M := 5.9736 * math.Pow(10,24)
	r := 6.378* math.Pow(10,6)
	es := (2*G*M)/r
	fmt.Println(es)
	esv := math.Sqrt(es)
	fmt.Println(esv)
	//fmt.Println(M)
	//fmt.Println(G)
}
