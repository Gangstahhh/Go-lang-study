package main

import "fmt"

func main(){ // Constant sama saja fungsinya seperti variabel bedanya constant valuenya tidak bisa di ubah
	const firstname string = "Muhammad"
	const lastname ="Risky"
	const value = 2000

	const ( //Versi Simpel Menulis constant
		country string = "Indonesia"
		daerah string  = "Jakarta"
	)

	fmt.Println(firstname,lastname,value)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)
	fmt.Println(country)
	fmt.Println(daerah)
}
