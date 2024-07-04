package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var UserLogin = make(map[string]model.User)

// DESC: func Auth is a middleware to check user login id, only user that already login can pass this middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_login_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if _, ok := UserLogin[c.Value]; !ok || c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // your code here }) // TODO: replace this
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
		return
	}

	payload := model.UserLogin{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if payload.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "ID or name is empty"})
		return
	}

	file, err := os.Open("../data/users.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	dataUser, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	lines := strings.Split(string(dataUser), "\n")
	var users []model.User

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 3 {
			user := model.User{
				ID:        fields[0],
				Name:      fields[1],
				StudyCode: fields[2],
				Role:      fields[3],
			}
			users = append(users, user)
		}
	}

	var foundUser model.User
	isUserExist := false
	for _, user := range users {
		if user.ID == payload.ID && user.Name == payload.Name {
			foundUser = user
			isUserExist = true
			break
		} else {
			isUserExist = false
		}
	}

	if !isUserExist {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user id not found"})
		return
	}

	UserLogin[foundUser.ID] = foundUser

	http.SetCookie(w, &http.Cookie{
		Name:  "user_login_id",
		Value: foundUser.ID,
	})

	// Memberikan cookie dengan key `user_login_role` dan value `<role user>`
	http.SetCookie(w, &http.Cookie{
		Name:  "user_login_role",
		Value: foundUser.Role,
	})

	res := model.SuccessResponse{
		Username: payload.Name,
		Message:  "Successfully",
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(res)
	if err != nil {
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// userID := r.Context().Value("userID").(string)

	// TODO: answer here
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	// list study program
	// TODO: answer here
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, r *http.Request) {
	// var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}
