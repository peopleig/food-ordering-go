package types

import "html/template"

type NewUser struct {
	Role       string `json:"role" schema:"role"`
	First_name string `json:"first_name" schema:"first_name"`
	Last_name  string `json:"last_name" schema:"last_name"`
	Email      string `json:"email" schema:"email"`
	Mobile     string `json:"mobile" schema:"mobile"`
	Password   string `json:"password" schema:"password"`
}

type Item struct {
	Item_id       int    `json:"item_id"`
	Item_name     string `json:"item_name"`
	Category_name string `json:"category_name"`
	Price         uint   `json:"price"`
	Description   string `json:"description"`
	Item_img      string `json:"item_img_url"`
	Is_veg        bool   `json:"is_veg"`
	Spice_level   int    `json:"spice_level"`
}

type MenuData struct {
	Title      string
	Items      map[int]Item
	Categories []Categories
	Role       string
}

type GroupedData struct {
	Title        string
	GroupedItems map[string][]Item
}

type OrderRequest struct {
	Cart                 []CartItem `json:"cart"`
	Special_instructions string     `json:"special_instructions"`
	Order_type           string     `json:"order_type"`
	Table_number         string     `json:"table_number"`
}

type CartItem struct {
	Quantity uint `json:"quantity"`
	Item_id  int  `json:"itemId"`
}

type Categories struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type GetAddDishData struct {
	Title      string       `json:"title"`
	Categories []Categories `json:"categories"`
	Error      bool         `json:"error"`
	Message    string       `json:"message"`
	Role       string       `json:"role"`
}

type ChefAssignRequest struct {
	ChefID  int `json:"chefId"`
	OrderID int `json:"orderId"`
	ItemID  int `json:"itemId"`
}

type OrdersData struct {
	Title  string    `json:"title"`
	Items  []Ordered `json:"items"`
	UserId int       `json:"user_id"`
	Role   string    `json:"role"`
}

type Ordered struct {
	OrderId      int    `json:"order_id"`
	ItemId       int    `json:"item_id"`
	ItemName     string `json:"item_name"`
	Quantity     uint   `json:"quantity"`
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
	Show    bool             `json:"show"`
	Message string           `json:"message"`
	Role    string           `json:"role"`
}

type Order struct {
	OrderId      int    `json:"order_id"`
	UserId       int    `json:"user_id"`
	Status       string `json:"status"`
	Order_type   string `json:"order_status"`
	Table_number int    `json:"table_number"`
}

type BillOrder struct {
	OrderId      int    `json:"order_id"`
	UserId       int    `json:"user_id"`
	UserName     string `json:"name"`
	Status       string `json:"status"`
	OrderType    string `json:"order_type"`
	TableNumber  int    `json:"table_number"`
	Instructions string `json:"instructions"`
	TotalCost    uint   `json:"total_cost"`
}

type CompleteBill struct {
	OrderId      int    `json:"order_id"`
	UserId       int    `json:"user_id"`
	UserName     string `json:"name"`
	Status       string `json:"status"`
	OrderType    string `json:"order_type"`
	TableNumber  int    `json:"table_number"`
	Instructions string `json:"instructions"`
	TotalCost    uint   `json:"total_cost"`
	AmtPaid      uint   `json:"amt_paid"`
	TipPaid      uint   `json:"tip_paid"`
}

type OrderContents struct {
	ItemName string `json:"item_name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}

type FinalBill struct {
	Title    string          `json:"title"`
	Contents []OrderContents `json:"contents"`
	Order    CompleteBill    `json:"order"`
	Show     bool            `json:"show"`
	Message  string          `json:"message"`
	Role     string          `json:"role"`
}

type SingleBill struct {
	Title    string          `json:"title"`
	Order    BillOrder       `json:"order"`
	Contents []OrderContents `json:"contents"`
	Role     string          `json:"role"`
}

type MyBills struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
	Price   uint   `json:"price"`
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
	Tip     uint `json:"tip"`
	OrderId int  `json:"order_id"`
}

type LoginData struct {
	LoginType  string `json:"login_type" schema:"login_type"`
	Identifier string `json:"identifier" schema:"identifier"`
	Password   string `json:"password" schema:"password"`
}

type NewDish struct {
	DishName    string `schema:"dish_name"`
	Category    string `schema:"category"`
	Price       uint   `schema:"price"`
	Description string `schema:"description"`
	IsVeg       bool   `schema:"is_veg"`
	SpiceLevel  int    `schema:"spice_level"`
}

type ErrorPageData struct {
	Title   string        `json:"title"`
	Status  string        `json:"status"`
	Message template.HTML `json:"message"`
	Role    string        `json:"role"`
}

type ShortBillForm struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	TotalCost   uint   `json:"total_cost"`
	TableNumber int    `json:"table_number"`
	OrderType   string `json:"order_type"`
}

type ShortBillData struct {
	Title      string          `json:"title"`
	ShortBills []ShortBillForm `json:"short_bills"`
	Role       string          `json:"role"`
	Show       bool            `json:"show"`
}
