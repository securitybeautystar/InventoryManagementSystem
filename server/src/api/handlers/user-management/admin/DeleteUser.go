package admin

import (
	"net/http"
)

type AdminDeleteUserMgmt struct {
	Username string `json:"username"`
}

func AdminDeleteUser(w http.ResponseWriter, req *http.Request) {
	// // Set Headers
	// w.Header().Set("Content-Type", "application/json")
	// var adminDeleteUser AdminDeleteUserMgmt

	// // Reading the request body and UnMarshal the body to the AdminDeleteUserMgmt struct
	// bs, _ := io.ReadAll(req.Body)
	// if err := json.Unmarshal(bs, &adminDeleteUser); err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal Server Error in Unmarshal JSON body in AdminDeleteUser:", err)
	// 	return
	// }

	// // Check User Group Admin
	// if !auth_management.RetrieveIssuer(w, req) {
	// 	return
	// }
	// if !utils.CheckUserGroup(w, w.Header().Get("Username"), "Admin") {
	// 	return
	// }

	// username := adminDeleteUser.Username

	// // Check username format
	// if !auth_management.UsernameFormValidation(w, username) {
	// 	return
	// }

	// // // Check if username exists in the database
	// if !database.GetUsername(username) {
	// 	utils.WriteJSON(w, http.StatusNotFound, "Username does not exist. Please try again.")
	// 	return
	// }

	// err := database.DeleteUserByID(username)
	// if err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal Server Error in deleting user from accounts table:", err)
	// 	return
	// }

	// utils.WriteJSON(w, http.StatusOK, "Successfully Deleted User!")
}