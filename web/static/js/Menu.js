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
    cartList.innerHTML = "";
    cart.forEach(item => {
        let li = document.createElement("li");
        let itemText = document.createElement("span");
        itemText.textContent = `${item.itemName} - Qty: ${item.quantity}`;
        let btnContainer = document.createElement("div");
        let incrementBtn = document.createElement("button");
        incrementBtn.textContent = "+";
        incrementBtn.onclick = () => incrementItem(item.itemName);
        let decrementBtn = document.createElement("button");
        decrementBtn.textContent = "-";
        decrementBtn.onclick = () => decrementItem(item.itemName);
        btnContainer.appendChild(incrementBtn);
        btnContainer.appendChild(decrementBtn);
        li.appendChild(itemText);
        li.appendChild(btnContainer);
        cartList.appendChild(li);
    });
    console.log(cart);
}


async function placeOrder() {
    try {
        const specialInstructions = document.getElementById("specialInstructions").value;
        const orderType = document.getElementById("orderType").value;
        const tableNumber = document.getElementById("tableNo").value;

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

