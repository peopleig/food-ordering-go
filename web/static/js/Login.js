const toggle_show = document.getElementById("show_password");
const pwd_field = document.getElementById("password");
const mobileOption = document.getElementById('mobile_option');
const emailOption = document.getElementById('email_option');
const identifierInput = document.getElementById('identifier');

toggle_show.addEventListener("change", () => {
    pwd_field.type = toggle_show.checked ? "text" : "password";
});

function updateIdentifierField() {
    if (mobileOption.checked) {
        identifierInput.placeholder = "Enter your mobile number";
        identifierInput.type = "tel";
        identifierInput.pattern = "\\d{10}"
    } 
    else if (emailOption.checked) {
        identifierInput.placeholder = "Enter your email address";
        identifierInput.type = "email";
    }
}
mobileOption.addEventListener('change', updateIdentifierField);
emailOption.addEventListener('change', updateIdentifierField);

updateIdentifierField();