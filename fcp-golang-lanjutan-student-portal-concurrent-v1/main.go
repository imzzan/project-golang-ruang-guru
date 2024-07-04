package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	failedLoginAttempts  map[string]int
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		failedLoginAttempts: map[string]int{},
	}
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	students := sm.students

	if students == nil {
		return []model.Student{}
	}

	return students
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if id == "" || len(id) == 0 || name == "" || len(name) == 0 {
		return "", errors.New("ID or Name is undefined!")
	}

	Students := sm.students

	hasil := ""
	var idStudent []string
	var nameStudent []string
	var programStudi []string
	for _, student := range Students {
		idStudent = append(idStudent, student.ID)
		nameStudent = append(nameStudent, student.Name)
		programStudi = append(programStudi, student.StudyProgram)
	}

	isLoginSuccess := false
	for i := 0; i < len(idStudent); i++ {
		if id != idStudent[i] || name != nameStudent[i] {
			isLoginSuccess = false
		} else {
			isLoginSuccess = true
			studiProgram, _ := sm.GetStudyProgram(programStudi[i])
			sm.failedLoginAttempts[id] = 0
			hasil = fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s", nameStudent[i], studiProgram)
			break
		}
	}

	if isLoginSuccess {
		return hasil, nil
	} else {
		sm.failedLoginAttempts[id] = sm.failedLoginAttempts[id] + 1
		if sm.failedLoginAttempts[id] > 3 {
			return "", errors.New("Login gagal: Batas maksimum login terlampaui")
		} else {
			return "", errors.New("Login gagal: data mahasiswa tidak ditemukan")
		}
	}

	return "", nil
}

func (sm *InMemoryStudentManager) RegisterLongProcess() {
	// 30ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// 30ms delay to simulate slow processing. DO NOT REMOVE THIS LINE
	sm.RegisterLongProcess()

	students := sm.students

	if id == "" || len(id) == 0 || name == "" || len(name) == 0 || studyProgram == "" || len(studyProgram) == 0 {
		return "", errors.New("ID, Name or StudyProgram is undefined!")
	}

	if studyProgram != "TI" && studyProgram != "TK" && studyProgram != "SI" && studyProgram != "MI" {
		return "", fmt.Errorf("Study program %s is not found", studyProgram)
	}

	hasil := ""
	for _, student := range students {
		if id == student.ID {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		} else {
			sm.students = append(sm.students, model.Student{
				ID:           id,
				Name:         name,
				StudyProgram: studyProgram,
			})
			hasil = fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram)
			break
		}
	}
	// Below lock is needed to prevent data race error. DO NOT REMOVE BELOW 2 LINES
	sm.Lock()
	defer sm.Unlock()

	return hasil, nil // TODO: replace this
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" || len(code) == 0 {
		return "", errors.New("Code is undefined!")
	}

	programStudi := sm.studentStudyPrograms
	isAvailableStudiProgram := false

	studiProgram := ""
	for key := range programStudi {
		if code != key {
			isAvailableStudiProgram = false
		} else {
			isAvailableStudiProgram = true
			break
		}
	}

	if isAvailableStudiProgram {
		studiProgram = programStudi[code]
	} else {
		return "", errors.New("Kode program studi tidak ditemukan")
	}

	return studiProgram, nil
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	students := sm.students
	isSuccessChanges := false

	for i, value := range students {
		if name != value.Name {
			isSuccessChanges = false
		} else {
			err := fn(&students[i])
			if err != nil {
				return "", errors.New("Sistem Error")
			}
			isSuccessChanges = true
			break
		}
	}

	if isSuccessChanges {
		return "Program studi mahasiswa berhasil diubah.", nil
	} else {
		return "", errors.New("Mahasiswa tidak ditemukan.")
	}
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		if _, ok := sm.studentStudyPrograms[programStudi]; !ok {
			return errors.New("Kode program studi tidak ditemukan")
		}
		s.StudyProgram = programStudi
		return nil
	}
}

func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
	for _, stucent := range filenames {
		stdnt, err := ReadStudentsFromCSV(stucent)
		if err != nil {
			return err
		}
		sm.students = append(sm.students, stdnt...)
	}

	start := time.Now()
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	sm.GetStudents()

	time.Sleep(200 * time.Millisecond)
	return nil
}

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// 3000ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {

	start := time.Now()

	for i := 0; i < numAssignments; i++ {
		for j := 1; j < 4; j++ {
			go sm.SubmitAssignmentLongProcess()
		}
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Submitting %d assignments took %s\n", numAssignments, elapsed)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "5":
			helper.ClearScreen()
			fmt.Println("=== Bulk Import Student ===")

			// Define the list of CSV file names
			csvFiles := []string{"students1.csv", "students2.csv", "students3.csv"}

			err := manager.ImportStudents(csvFiles)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful!")
			}

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')

		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")

			// Enter how many assignments you want to submit
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignments, _ := reader.ReadString('\n')

			// Convert the input to an integer
			numAssignments = strings.TrimSpace(numAssignments)
			numAssignmentsInt, err := strconv.Atoi(numAssignments)

			if err != nil {
				fmt.Println("Error: Please enter a valid number")
			}

			manager.SubmitAssignments(numAssignmentsInt)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
