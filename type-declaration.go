package main
import "fmt"
func main(){
	type NoKTP string //untuk alias suatu type data
	type married bool

	var noKtpiki NoKTP = "1124020482304"
	var marriedStatus married = false
	fmt.Println(noKtpiki)
	fmt.Println(marriedStatus)
}