const toggle_show = document.getElementById("show_password");
const pwd_field = document.getElementById("password");
toggle_show.addEventListener("change", () => {
    pwd_field.type = toggle_show.checked ? "text" : "password";
});

let selectedRole = "customer";
const chefAlert = document.querySelector('.chef-alert');
const adminAlert = document.querySelector('.admin-alert');


document.addEventListener("DOMContentLoaded", function () {
    const roleRadios = document.querySelectorAll('input[name="role"]');
    roleRadios.forEach(function (radio) {
        radio.addEventListener('change', function () {
            if (this.checked) {
                selectedRole = this.value;
                if (selectedRole === 'chef') {
                    chefAlert.classList.remove('d-none');
                    adminAlert.classList.add('d-none');
                } 
                else if (selectedRole === 'admin') {
                    adminAlert.classList.remove('d-none');
                    chefAlert.classList.add('d-none');
                }
                else {
                    adminAlert.classList.add('d-none');
                    chefAlert.classList.add('d-none');   
                }
            }
        });
    });
});

