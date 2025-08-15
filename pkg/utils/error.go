package utils

func AddDishErrors(errorMsg string) (bool, string) {
	showToast := false
	var message string
	switch errorMsg {
	case "img":
		showToast = true
		message = "Invalid image file. Only PNG/JPG/WEBP allowed"
	case "server":
		showToast = true
		message = "Sorry, server error. Please try again!"
	case "req":
		showToast = true
		message = "Incorrect Request Sent!"
	case "dish":
		showToast = true
		message = "Dish with such a name already exists!"
	}
	return showToast, message
}
