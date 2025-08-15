const toggle_show = document.getElementById("show_password");
const pwd_field = document.getElementById("password");
toggle_show.addEventListener("change", () => {
    pwd_field.type = toggle_show.checked ? "text" : "password";
});