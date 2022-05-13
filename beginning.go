/*
kullanıcıdan km başına yakılan yakıt miktarı istensin, bir de gidilecek kilometre istensin.
böylece yolun sonunda ne kadar masraf tutmuş oluyor bunu görelim.
*/

package main

import "fmt"

func main(){
	var kmBasinaYakitTL float32
	var kacKmGidilecek float32
	var toplamTutarTL float32
	
	fmt.Println("Kilometre Başına Kaç ₺ yakıt yakılmaktadır? : ")
	fmt.Scanln(&kmBasinaYakitTL)
	fmt.Println("Kaç Kilometre Yol Gidilecektir? : ")
	fmt.Scanln(&kacKmGidilecek)

	toplamTutarTL = kmBasinaYakitTL * kacKmGidilecek

	fmt.Println("Yol Sonu Toplam Tutar : ",toplamTutarTL,"₺'dir.")

}