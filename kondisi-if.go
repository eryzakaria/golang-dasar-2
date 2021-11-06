package main

import "fmt"

func main() {
	nilai := 4

	if nilai == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if nilai > 70 {
		fmt.Println("Lulus")
	} else if nilai > 50 {
		fmt.Println("Nilai Kurang Sedikit")
	} else {
		fmt.Println("Murid Goblok")
	}

	point := 8840.0

	if check := point / 100; check > 100 {
		fmt.Printf("Hasil Memuaskan \n")
	} else if check > 70 {
		fmt.Printf("Anda Dinyatakan Lulus (%.1f) \n",check)
	} else {
		fmt.Printf("Anda Tidak Lulus \n")
	}

	ujian := 100
	absensi := 90

	if ujian >= 80 && absensi >= 85  {
		fmt.Println("Anda Dinyatakan lulus dengan predikat Terbaik")
	} else if ujian >= 70 && absensi >= 70 {
		fmt.Println("Anda Dinyatakan Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

} 