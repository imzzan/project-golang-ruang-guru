package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {

	Student := strings.Split(Students, ",")

	if id == "" || len(id) == 0 && name == "" || len(name) == 0 {
		return "ID or Name is undefined!"
	}

	if len(id) < 5 || len(id) > 5 {
		return "ID must be 5 characters long!"
	}

	hasil := ""
	var idStudent []string
	var programStudi []string
	var nameStudent []string
	for _, student := range Student {
		ids := strings.Split(student, "_")
		idStudent = append(idStudent, ids[0])
		nameStudent = append(nameStudent, ids[1])
		programStudi = append(programStudi, ids[2])
	}

	ids := strings.Join(idStudent, "")
	idStudent = strings.Split(ids, " ")
	for i := 0; i < len(idStudent); i++ {
		if id != idStudent[i] || name != nameStudent[i] {
			hasil = "Login gagal: data mahasiswa tidak ditemukan"
		} else {
			hasil = fmt.Sprintf("Login berhasil: %s (%s)", nameStudent[i], programStudi[i])
			break
		}
	}

	return hasil
}

func Register(id string, name string, major string) string {

	Student := strings.Split(Students, ",")

	if id == "" || len(id) == 0 || name == "" || len(name) == 0 || major == "" || len(major) == 0 {
		return "ID, Name or Major is undefined!"
	}

	if len(id) < 5 || len(id) > 5 {
		return "ID must be 5 characters long!"
	}

	for _, studentId := range Student {
		if studentId == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	var idStudent []string

	for _, student := range Student {
		a := strings.Split(student, "_")
		idStudent = append(idStudent, a[0])
	}
	idStudentStr := strings.Join(idStudent, "")
	idStudent = strings.Split(idStudentStr, " ")

	for i := 0; i < len(idStudent); i++ {
		if id == idStudent[i] {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {

	prgmStudy := []string{}
	programStudi := strings.Split(StudentStudyPrograms, ",")
	for _, s := range programStudi {
		prgmStdi := strings.Split(s, "_")
		prgmStudy = append(prgmStudy, prgmStdi[0])
	}

	prgmStudyStr := strings.Join(prgmStudy, "")
	programStudiStudent := strings.Split(prgmStudyStr, " ")
	for i := 0; i < len(programStudiStudent); i++ {
		if code == programStudiStudent[0] {
			return "Teknik Informatika"
		} else if code == programStudiStudent[1] {
			return "Teknik Komputer"
		} else if code == programStudiStudent[2] {
			return "Sistem Informasi"
		} else if code == programStudiStudent[3] {
			return "Manajemen Informasi"
		}
	}

	return "Code is undefined!"

}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

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
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
