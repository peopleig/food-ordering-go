package types

type NewUser struct {
	Role       string `json:"role" schema:"role"`
	First_name string `json:"first_name" schema:"first_name"`
	Last_name  string `json:"last_name" schema:"last_name"`
	Email      string `json:"email" schema:"email"`
	Mobile     string `json:"mobile" schema:"mobile"`
	Password   string `json:"password" schema:"password"`
}

type Item struct {
	Item_id       int     `json:"item_id"`
	Item_name     string  `json:"item_name"`
	Category_name string  `json:"category_name"`
	Price         float32 `json:"price"`
	Description   string  `json:"description"`
	Item_img      string  `json:"item_img_url"`
	Is_veg        bool    `json:"is_veg"`
	Spice_level   int     `json:"spice_level"`
}

type MenuData struct {
	Title string
	Items []Item
}

type OrderRequest struct {
	Cart                 []CartItem `json:"cart"`
	Special_instructions string     `json:"special_instructions"`
	Order_type           string     `json:"order_type"`
	Table_number         string     `json:"table_number"`
}

type CartItem struct {
	Item_name string `json:"itemName"`
	Quantity  int    `json:"quantity"`
	Item_id   int    `json:"itemId"`
	Price     int    `json:"price"`
}

type ChefAssignRequest struct {
	ChefID  int `json:"chefId"`
	OrderID int `json:"orderId"`
	ItemID  int `json:"itemId"`
}

type OrdersData struct {
	Title string    `json:"title"`
	Items []Ordered `json:"items"`
}

type Ordered struct {
	OrderId      int    `json:"order_id"`
	ItemId       int    `json:"item_id"`
	ItemName     string `json:"item_name"`
	Quantity     int    `json:"quantity"`
	ChefId       int    `json:"chef_id"`
	ChefName     string `json:"chef_name"`
	Instructions string `json:"instructions"`
	Order_type   string `json:"order_type"`
	Assigned     bool   `json:"assigned"`
}

type AdminData struct {
	Title   string           `json:"title"`
	Items   []Ordered        `json:"items"`
	Orders  []Order          `json:"orders"`
	Uausers []UnApprovedUser `json:"uauser"`
}

type Order struct {
	OrderId      int    `json:"order_id"`
	UserId       int    `json:"user_id"`
	Status       string `json:"status"`
	Order_type   string `json:"order_status"`
	Table_number int    `json:"table_number"`
}

type MyBills struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
	Price   int    `json:"price"`
}

type UnApprovedUser struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type BillData struct {
	Title  string    `json:"title"`
	MyBill []MyBills `json:"all_bills"`
}

type BillPay struct {
	Tip     int `json:"tip"`
	OrderId int `json:"order_id"`
}
