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
	{Id: 2, Name: "Agus", Address: "Denpasar", Job: "Mahasiswa", Reason: "ingin menambah ilmu"},
	{Id: 3, Name: "Agus", Address: "Sumatera", Job: "Mahasiswa", Reason: "ingin mengembangkan hard skill"},
	{Id: 4, Name: "Adi", Address: "Sulawesi", Job: "Mahasiswa", Reason: "untuk menambah wawasan koding"},
	{Id: 5, Name: "Bagas", Address: "Bandung", Job: "Mahasiswa", Reason: "ingin mendapat pengalaman baru"},
	{Id: 6, Name: "Berry", Address: "Papua", Job: "Mahasiswa", Reason: "karena termotivasi dari orang hebat"},
	{Id: 7, Name: "Caca", Address: "Palembang", Job: "Mahasiswa", Reason: "karena ingin meningkatkan skill"},
	{Id: 8, Name: "Cinta", Address: "Kupang", Job: "Mahasiswa", Reason: "karena ingin mendapat pengakuan yang lebih"},
	{Id: 9, Name: "Cahya", Address: "Ponorogo", Job: "Mahasiswa", Reason: "ingin meningkatkan produktivitas"},
	{Id: 10, Name: "Dinda", Address: "Madiun", Job: "Mahasiswa", Reason: "ingin belajar sesuatu yang baru"},
	{Id: 11, Name: "Diah", Address: "Jogja", Job: "Mahasiswa", Reason: "ingin menambah pengetahuan"},
	{Id: 12, Name: "Dono", Address: "Semarang", Job: "Mahasiswa", Reason: "karena ingin memiliki jenjang karir yang bagus"},
	{Id: 13, Name: "Erwin", Address: "Surabaya", Job: "Mahasiswa", Reason: "karena ingin di terima di perusahaan besar"},
	{Id: 14, Name: "Eca", Address: "Banten", Job: "Mahasiswa", Reason: "karena ingin mencari soft skill"},
	{Id: 15, Name: "Edi", Address: "Aceh", Job: "Mahasiswa", Reason: "karena rasa ingin belajar yang besar"},
	{Id: 16, Name: "Fani", Address: "Gorontalo", Job: "Mahasiswa", Reason: "karena melihat peluang kerja yang bagus"},
	{Id: 17, Name: "Iriana", Address: "Jambi", Job: "Mahasiswa", Reason: "karena sedang banyak digunakan perusahaan"},
	{Id: 18, Name: "Gina", Address: "Bekasi", Job: "Mahasiswa", Reason: "karena menambah pengalaman"},
	{Id: 19, Name: "Hana", Address: "Bogor", Job: "Mahasiswa", Reason: "karena suka belajar hal baru"},
	{Id: 20, Name: "Hema", Address: "Depok", Job: "Mahasiswa", Reason: "karena ingin mendapatkan ilmu baru"},
	{Id: 21, Name: "Kayla", Address: "Sukabumi", Job: "Mahasiswa", Reason: "karena rasa ingin tahu yang tinggi"},
	{Id: 22, Name: "Mario", Address: "Banjar", Job: "Mahasiswa", Reason: "karena peluang karir yang tinggi"},
	{Id: 23, Name: "Lusi", Address: "Malang", Job: "Mahasiswa", Reason: "karena memiliki berbagai manfaat"},
	{Id: 24, Name: "Nina", Address: "Kediri", Job: "Mahasiswa", Reason: "karena dibutuhkan perusahaan"},
	{Id: 25, Name: "Rahayu", Address: "Pontianak", Job: "Mahasiswa", Reason: "karena berguna saat mencari pekerjaan"},
	{Id: 26, Name: "Surya", Address: "Banjar", Job: "Mahasiswa", Reason: "karena ingin menambah informasi pada bidang programming"},
	{Id: 27, Name: "Bima", Address: "Balikpapan", Job: "Mahasiswa", Reason: "karena ingin mengembangkan kemampuan diri"},
	{Id: 28, Name: "Kelvin", Address: "Samarinda", Job: "Mahasiswa", Reason: "karena ingin meningkatkan kualitas diri"},
	{Id: 29, Name: "Putra", Address: "Batam", Job: "Mahasiswa", Reason: "karena ingin menjadi tenaga kerja yang terampil"},
	{Id: 30, Name: "Wisnu", Address: "Ambon", Job: "Mahasiswa", Reason: "karena ingin menambah wawasan dan keterampilan pada bidang programming"},
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
