package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {

	if id == "" || len(id) == 0 && name == "" || len(name) == 0 {
		return "ID or Name is undefined!"
	}

	hasil := ""
	var idStudent []string
	var nameStudent []string
	for _, student := range Students {
		ids := strings.Split(student, "_")
		idStudent = append(idStudent, ids[0])
		nameStudent = append(nameStudent, ids[1])
	}

	for i := 0; i < len(idStudent); i++ {
		if id != idStudent[i] || name != nameStudent[i] {
			hasil = "Login gagal: data mahasiswa tidak ditemukan"
		} else {
			hasil = fmt.Sprintf("Login berhasil: %s", nameStudent[i])
			break
		}
	}

	return hasil

}

func Register(id string, name string, major string) string {
	if id == "" || len(id) == 0 || name == "" || len(name) == 0 || major == "" || len(major) == 0 {
		return "ID, Name or Major is undefined!"
	}

	if len(id) < 5 || len(id) > 5 {
		return "ID must be 5 characters long!"
	}
	var idStudent []string

	for _, student := range Students {
		a := strings.Split(student, "_")
		idStudent = append(idStudent, a[0])
	}
	hasil := ""
	for i := 0; i < len(idStudent); i++ {
		if id == idStudent[i] {
			hasil = "Registrasi gagal: id sudah digunakan"
			break
		} else {
			newStudent := fmt.Sprintf("%s_%s_%s", id, name, major)
			Students = append(Students, newStudent)
			hasil = fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
			break
		}
	}

	return hasil
}

func GetStudyProgram(code string) string {
	switch code {
	case "TI":
		return StudentStudyPrograms[code]
	case "TK":
		return StudentStudyPrograms[code]
	case "SI":
		return StudentStudyPrograms[code]
	case "MI":
		return StudentStudyPrograms[code]
	}

	return "Kode program studi tidak ditemukan"
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	var students []string
	for _, student := range Students {
		ids := strings.Split(student, "_")
		students = append(students, ids[1])
	}

	isSuccessChanges := false
	faildMessage := ""
	for i, value := range students {
		if nama == value {
			fn(programStudi, &Students[i])
			isSuccessChanges = true
		} else {
			faildMessage = "Mahasiswa tidak ditemukan."
		}
	}

	if isSuccessChanges {
		return "Program studi mahasiswa berhasil diubah."
	}

	return faildMessage
}

func UpdateStudyProgram(programStudi string, students *string) {
	student := strings.Split(*students, "_")
	*students = fmt.Sprintf("%s_%s_%s", student[0], student[1], programStudi)
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
