let cart = [];

function revertAddButton(itemId) {
    const buttons = document.querySelectorAll(`.add-buttons`);
    buttons.forEach(btn => {
        if (btn.getAttribute('data-item-id') === itemId.toString()) {
            btn.textContent = "Add to Cart";
            btn.classList.remove("btn-success");
            btn.classList.add("btn-primary");
        }
    });
}

function addToCart( itemId, itemName, price, button) {
    let item = cart.find(i => i.itemName === itemName);
    if (item) {
        item.quantity++;
    } else {
        cart.push({itemName, quantity: 1, itemId: parseInt(itemId), price: parseInt(price)});
        button.textContent = "Added to Cart";
        button.classList.remove("btn-primary");
        button.classList.add("btn-success");
    }
    renderCart();
}

function incrementItem(itemName) {
    let item = cart.find(i => i.itemName === itemName);
    if (item) {
        item.quantity++;
        renderCart();
    }
}


function decrementItem(itemName) {
    let item = cart.find(i => i.itemName === itemName);
    if (item) {
        item.quantity--;
        if (item.quantity <= 0) {
            cart = cart.filter(i => i.itemName !== itemName);
            revertAddButton(item.itemId);
        }
        renderCart();
    }
}

function renderCart() {
    const cartList = document.getElementById("cart-list");
    const cartTotal = document.getElementById("cart-total");
    cartList.innerHTML = "";
    let total = 0;

    cart.forEach(item => {
        total += item.price * item.quantity;
        let li = document.createElement("li");
        li.classList.add("cart-item");
        let nameSpan = document.createElement("span");
        nameSpan.classList.add("item-name");
        nameSpan.textContent = item.itemName;
        let qtySpan = document.createElement("span");
        qtySpan.classList.add("item-qty");
        qtySpan.textContent = `Qty: ${item.quantity}`;
        let priceSpan = document.createElement("span");
        priceSpan.classList.add("item-price");
        priceSpan.textContent = `₹${item.price * item.quantity}`;
        let btnContainer = document.createElement("div");
        btnContainer.classList.add("item-controls");
        let incrementBtn = document.createElement("button");
        incrementBtn.textContent = "+";
        incrementBtn.classList.add("btn", "btn-sm", "btn-outline-success");
        incrementBtn.onclick = () => incrementItem(item.itemName);
        let decrementBtn = document.createElement("button");
        decrementBtn.textContent = "-";
        decrementBtn.classList.add("btn", "btn-sm", "btn-outline-danger", "ms-1");
        decrementBtn.onclick = () => decrementItem(item.itemName);
        btnContainer.appendChild(incrementBtn);
        btnContainer.appendChild(decrementBtn);
        li.appendChild(nameSpan);
        li.appendChild(qtySpan);
        li.appendChild(priceSpan);
        li.appendChild(btnContainer);
        cartList.appendChild(li);
    });

    cartTotal.textContent = `Total: ₹${total}`;
    console.log(cart);
}


const table = document.getElementById("tableNo");
let tableNumber;
table.addEventListener("input", () => {
    tableNumber = table.value;
    if (tableNumber < 0) {
        tableNumber = 0;
        table.value = 0;
    }
    else if (tableNumber > 20) {
        tableNumber = 20;
        table.value = 20;
    }
});

async function placeOrder() {
    try {
        const specialInstructions = document.getElementById("specialInstructions").value;
        const orderType = document.getElementById("orderType").value;
        const payload = {
            cart: cart,
            special_instructions: specialInstructions,
            order_type: orderType,
            table_number: tableNumber
        };

        const response = await fetch("/menu", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        });
        if (!response.ok){
            alert("Unable to Place Order");
            throw new Error("Order request failed");
        }

        const result = await response.json();
        console.log("Order placed successfully:", result);
        alert("Order placed successfully!");
        cart = [];
        renderCart();
        window.location.reload();
    } 
    catch (error) {
        console.error("Error placing order:", error);
    }
}

const order_filters = document.querySelectorAll(".order-filter");
order_filters.forEach(button => {
    button.addEventListener("click", () => {
        const status = button.dataset.status;
        document.querySelectorAll(".orders-card").forEach(card => {    
            if (card.dataset.status === status) {
                card.classList.remove("d-none");
                card.classList.add("d-inline-block");
            } else {
                card.classList.remove("d-inline-block");
                card.classList.add("d-none");
            }
        });
        order_filters.forEach(b => b.classList.remove("active"));
        button.classList.add("active");
    });
});


document.addEventListener("DOMContentLoaded", () => {
    const surpriseBtn = document.getElementById("surprise-me-btn");
    const surprisePopup = document.getElementById("surprise-popup");
    const cancelBtn = document.getElementById("cancel-surprise");
    const generateBtn = document.getElementById("generate-surprise");

    surpriseBtn.onclick = () => {
        document.getElementById("people-count").value = "";
        surprisePopup.classList.remove("d-none");
    };

    cancelBtn.onclick = () => {
        surprisePopup.classList.add("d-none");
    };

    generateBtn.onclick = () => {
        const peopleCount = parseInt(document.getElementById("people-count").value);
        if (isNaN(peopleCount) || peopleCount <= 0) {
            alert("Enter a valid number!");
            return;
        }
        cart.forEach(item => revertAddButton(item.itemId));
        cart = [];
        renderCart();

        const allButtons = [...document.querySelectorAll(".add-buttons")];

        const allDishes = allButtons.map(btn => {
            const card = btn.closest(".card");
            const name = card.querySelector(".card-title").textContent.trim();
            const price = parseFloat(card.querySelector(".fw-bold").textContent.replace("₹", "").trim());
            const isVeg = card.querySelector(".badge.bg-success") !== null;
            const itemId = parseInt(btn.getAttribute("data-item-id"));
            const img = card.querySelector("img").src;
            const category = categorizeDish(card.querySelector(".badge.bg-warning.text-dark").textContent.trim());
            return { itemId, name, price, isVeg, img, btn, category };
        });
        const randomItem = arr => arr[Math.floor(Math.random() * arr.length)];
        const beverages = allDishes.filter(d => d.category === "beverage");
        const desserts = allDishes.filter(d => d.category === "dessert");
        const mains = allDishes.filter(d => d.category === "main");
        const selected = [];
        if (beverages.length > 0)
            selected.push({ ...randomItem(beverages), quantity: peopleCount });
        if (desserts.length > 0)
            selected.push({ ...randomItem(desserts), quantity: peopleCount });
        const shuffledMains = mains.sort(() => 0.5 - Math.random());
        selected.push(...shuffledMains.slice(0, peopleCount).map(d => ({ ...d, quantity: 1 })));
        selected.forEach(dish => {
            cart.push({
                itemName: dish.name,
                quantity: dish.quantity,
                itemId: dish.itemId,
                price: dish.price
            });
            dish.btn.textContent = "Added to Cart";
            dish.btn.classList.remove("btn-primary");
            dish.btn.classList.add("btn-success");
        });
        renderCart();
        surprisePopup.classList.add("d-none");
        document.getElementById("cart-list").scrollIntoView({ behavior: "smooth" });
    }
    function categorizeDish(name) {
        if (name === "Desserts") {
            return "dessert";
        }
        else if (name === "Beverages")
            return "beverage";
        return "main";
    }
});


const categoryFilter = document.getElementById("categoryFilter");
const cards = document.querySelectorAll(".menu-grid .card");
categoryFilter.addEventListener("change", () => {
    const selectedCategory = categoryFilter.value;
    cards.forEach(card => {
        const cardCategory = card.getAttribute("data-category");
        if (selectedCategory === "all" || cardCategory === selectedCategory) {
            card.style.display = "block";
        } else {
            card.style.display = "none";
        }
    });
});

document.addEventListener("DOMContentLoaded", function () {
    const orderTypeSelect = document.getElementById('orderType');
    const tableNoInput = document.getElementById('tableNo');
    function updateTableNo() {
        if (orderTypeSelect.value === 'takeaway') {
            tableNoInput.value = 0;
            tableNoInput.disabled = true;
        } else {
            tableNoInput.disabled = false;
            tableNoInput.value = '';
        }
    }
    updateTableNo();
    orderTypeSelect.addEventListener('change', updateTableNo);
});

