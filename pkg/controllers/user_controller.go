package controllers

type UserController struct{}

// func NewUserController() *UserController {
// 	return &UserController{}
// }

// func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := models.GetAllUsers()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }

// func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

// 	path := strings.TrimPrefix(r.URL.Path, "/users/")
// 	id, err := strconv.Atoi(path)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := models.GetUserByID(id)
// 	if err != nil {
// 		if err.Error() == "user not found" {
// 			http.Error(w, "User not found", http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }

// func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
// 		return
// 	}

// 	if user.Name == "" || user.Address == "" || user.Country == "" {
// 		http.Error(w, "Name, address, and country are required", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := models.CreateUser(user)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	response := map[string]interface{}{
// 		"message": "User created successfully",
// 		"id":      id,
// 	}
// 	json.NewEncoder(w).Encode(response)
// }

// func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPut {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	path := strings.TrimPrefix(r.URL.Path, "/users/update/")
// 	id, err := strconv.Atoi(path)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	var user models.User
// 	err = json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
// 		return
// 	}
// }
