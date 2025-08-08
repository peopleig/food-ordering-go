package models

import "strconv"

type User struct {
	User_id  string `json:"id"`
	Role     string `json:"role"`
	Hash_pwd string `json:"hash"`
}

// func GetAllUsers() ([]User, error) {
// 	query := "SELECT user_id, first_name, last_name FROM User"
// 	rows, err := DB.Query(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("error querying users: %v", err)
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.User_id, &user.First_name, &user.Last_name)
// 		if err != nil {
// 			return nil, fmt.Errorf("error scanning user: %v", err)
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

func GetUserPwdatLogin(login_type string, identifier string) (User, error) {
	var query string
	var value interface{}
	if login_type == "email" {
		query = "SELECT user_id, role, password FROM User WHERE email_id = ?"
		value = identifier
	} else {
		value, _ = strconv.Atoi(identifier)
		query = "SELECT user_id, role, password FROM User WHERE mobile_number = ?"
	}
	var user User
	err := DB.QueryRow(query, value).Scan(&user.User_id, &user.Role, &user.Hash_pwd)
	if err != nil {
		return user, err
	}
	return user, nil
}
