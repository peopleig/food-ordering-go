package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/peopleig/food-ordering-go/pkg/types"
)

type User struct {
	User_id  string `json:"id"`
	Role     string `json:"role"`
	Hash_pwd string `json:"hash"`
	Approved bool   `json:"approved"`
}

func GetUserPwdatLogin(login_type string, identifier string) (User, error) {
	var query string
	var value interface{}
	if login_type == "email" {
		query = "SELECT user_id, role, password, approved FROM User WHERE email_id = ?"
		value = identifier
	} else {
		value, _ = strconv.Atoi(identifier)
		query = "SELECT user_id, role, password, approved FROM User WHERE mobile_number = ?"
	}
	var user User
	err := DB.QueryRow(query, value).Scan(&user.User_id, &user.Role, &user.Hash_pwd, &user.Approved)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateNewUser(new_user *types.NewUser) (bool, string, int64, error) {
	mobile_num, _ := strconv.Atoi(new_user.Mobile)
	query := `SELECT COUNT(*) FROM User WHERE (email_id = ? AND ? <> '') OR (mobile_number = ? AND ? <> '')`
	var num int
	err := DB.QueryRow(query, new_user.Email, new_user.Email, mobile_num, mobile_num).Scan(&num)
	if err != nil {
		fmt.Println(err)
		return false, "Error in accessing database", 0, err
	}
	if num > 0 {
		return false, "OOPS! This Email/Mobile No. is already in use", 0, nil
	}
	query = `INSERT INTO User (first_name, last_name, email_id, mobile_number, password, role, approved) VALUES (?, ?, ?, ?, ?, ?, ?)`
	approved := (new_user.Role == "customer")
	var res sql.Result
	res, err = DB.Exec(query, new_user.First_name, new_user.Last_name, new_user.Email, mobile_num, new_user.Password, new_user.Role, approved)
	if err != nil {
		return false, "Error in accessing database", 0, err
	}
	user_id, err := res.LastInsertId()
	if err != nil {
		return false, "Error in accessing database", 0, err
	}
	return true, "", user_id, nil
}

func CheckForUser(orderId int) (int, error) {
	query := `SELECT user_id FROM Orders WHERE order_id = ?`
	var userId int
	fmt.Println(orderId)
	err := DB.QueryRow(query, orderId).Scan(&userId)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return userId, nil
}
