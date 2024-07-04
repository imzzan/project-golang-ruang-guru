package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	StudyName    string `json:"study_name"`
	StudyCreadit int    `json:"study_creadit"`
	Grade        string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {

	file, err := os.Open(filename)
	if err != nil {
		return Report{}, err
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return Report{}, err
	}
	report := Report{}
	err = json.Unmarshal([]byte(jsonData), &report)
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
		return 0
	}
	acumIp := 0.0
	acumCredit := 0
	for _, study := range report.Studies {
		gradeMap := map[string]float64{
			"A":  4.0,
			"AB": 3.5,
			"B":  3.0,
			"BC": 2.5,
			"C":  2.0,
			"CD": 1.5,
			"D":  1.0,
			"DE": 0.5,
			"E":  0,
		}

		acumIp += float64(study.StudyCreadit) * gradeMap[study.Grade]
		acumCredit += study.StudyCreadit
	}
	return acumIp / float64(acumCredit)
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
