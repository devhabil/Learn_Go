package main

import "fmt"

func main() {
	// Mendeklarasikan array bertipe string dengan 6 elemen
	names := [...]string{"Muhammad", "Habil", "Arifin", "Fajar", "Joko", "Anwar"}

	// Membuat slice dari elemen indeks 4 hingga sebelum indeks 6 (yaitu: "Joko", "Anwar")
	slice1 := names[4:6]
	fmt.Println(slice1) // Output: [Joko Anwar]

	// Membuat slice dari awal array hingga sebelum indeks 3 (yaitu: "Muhammad", "Habil", "Arifin")
	slice2 := names[:3]
	fmt.Println(slice2) // Output: [Muhammad Habil Arifin]

	// Membuat slice dari indeks 3 hingga akhir array (yaitu: "Fajar", "Joko", "Anwar")
	slice3 := names[3:]
	fmt.Println(slice3) // Output: [Fajar Joko Anwar]

	// Membuat slice dari seluruh array
	slice4 := names[:]
	fmt.Println(slice4) // Output: [Muhammad Habil Arifin Fajar Joko Anwar]

	// Deklarasi array hari
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// Membuat slice dari elemen indeks 5 hingga akhir (yaitu: "Sabtu", "Minggu")
	daySlice1 := days[5:]
	fmt.Println(daySlice1) // Output: [Sabtu Minggu]

	// Mengubah isi slice — ini juga mempengaruhi array asli karena slice menunjuk ke array
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	// Menampilkan slice dan array setelah perubahan
	fmt.Println(daySlice1) // Output: [Sabtu Baru Minggu Baru]
	fmt.Println(days)      // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// Menambahkan elemen ke slice menggunakan append
	daySlice2 := append(daySlice1, "Libur Baru")

	// Mengubah elemen pertama dari slice hasil append
	daySlice2[0] = "Sabtu Lama"

	// Menampilkan isi kedua slice dan array
	fmt.Println(daySlice1) // Output: [Sabtu Lama Minggu Baru]
	fmt.Println(daySlice2) // Output: [Sabtu Lama Minggu Baru Libur Baru]
	fmt.Println(days)      // Output bisa sama atau berbeda tergantung apakah append menyebabkan alokasi array baru

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Habil"
	newSlice[1] = "Habil"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Habil")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Eko"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice :=[]int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}