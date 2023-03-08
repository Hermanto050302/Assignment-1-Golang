package main

import (
	"fmt"
	"os"
)

type Student struct {
	Id      int
	Name    string
	Address string
	Job     string
	Reason  string
}

var listStudens = []Student{
	{Id: 1, Name: "Thomas", Address: "Jakarta", Job: "Mahasiswa", Reason: "karena ingin belajar"},
	{Id: 2, Name: "Agus", Address: "Denpasar", Job: "Mahasiswa", Reason: "karena ingin belajar"},
	{Id: 3, Name: "Agus", Address: "Sumatera", Job: "Mahasiswa", Reason: "karena ingin belajar"},
	{Id: 4, Name: "Adi", Address: "Sulawesi", Job: "Mahasiswa", Reason: "karena ingin belajar"},
}

func main() {
	//mengambil argumen dari command line
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run biodata.go <name>")
		return
	}

	//mencari data
	name := cariData(os.Args[1])

	//menampilkan data
	tampilData(name)

}

func cariData(name string) []Student {
	var result []Student
	for _, peserta := range listStudens {
		if peserta.Name == name {
			result = append(result, peserta)
		}
	}
	return result
}

func tampilData(name []Student) {
	//jika data tidak ditemukan
	if len(name) == 0 {
		fmt.Println("Data tidak ditemukan")
		return
	}
	for _, peserta := range name {
		fmt.Println("Id:", peserta.Id)
		fmt.Println("Nama:", peserta.Name)
		fmt.Println("Alamat:", peserta.Address)
		fmt.Println("Pekerjaan:", peserta.Job)
		fmt.Println("Alasan:", peserta.Reason)
		fmt.Println()
	}
}
