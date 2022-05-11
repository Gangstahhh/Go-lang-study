package main 

import "fmt"

func main(){
	var name string // membuat variabel harus spesifik dengan tipe data. kalau ingin string valuenya otomatis harus string juga tipe datanya

	name = "Muhammad Risky" 
	fmt.Println(name)

	var friendname = "Bambang" // bisa juga dengan memasukan valuenya langsung jadi bahasa go ini pintar jadi langsung otomatis mengikuti valuenya
	fmt.Println(friendname)	 

	var age  = 16 
	fmt.Println(age)

	nama := "Muhammad Risky" // bisa di sederhanakan lagi seperti ini dengan nama variabel di ikuti := (titik dua sama dengan)
	fmt.Println(nama)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstname = "Muhammad"
		lastname = "Risky"
	)
	fmt.Println(firstname,lastname)

}