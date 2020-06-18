package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	//router.HandleFunc("/api/client", middleware.GetAllLawyerEmails).Methods("GET", "OPTIONS")
	router.HandleFunc("/client/api/client", middleware.CreateClientsInfo).Methods("POST", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/sign_up/api/signup", middleware.LawyerSignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/api/signin", middleware.LawyerSignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/api/signin", middleware.SendAuth).Methods("GET", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/api/opencases", middleware.GetOpenCase).Methods("GET", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/api/mycases", middleware.GetMyCase).Methods("GET", "OPTIONS")
	router.HandleFunc("/lawyerdashboard/api/takecase/{id}", middleware.CaseComplete).Methods("PUT", "OPTIONS")
	return router
}
