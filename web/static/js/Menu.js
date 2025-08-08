let cart = [];

function addToCart( itemId, itemName, price) {
    let item = cart.find(i => i.itemName === itemName);
    if (item) {
        item.quantity++;
    } else {
        cart.push({itemName, quantity: 1, itemId: parseInt(itemId), price: parseInt(price)});
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
        }
        renderCart();
    }
}

function renderCart() {
    const cartList = document.getElementById("cart-list");
    cartList.innerHTML = "";
    cart.forEach(item => {
        let li = document.createElement("li");
        li.innerHTML = `${item.itemName} - Qty: ${item.quantity} 
            <button onclick="incrementItem('${item.itemName}')">+</button>
            <button onclick="decrementItem('${item.itemName}')">-</button>`;
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

        if (!response.ok) throw new Error("Order request failed");

        const result = await response.json();
        console.log("Order placed successfully:", result);
        alert("Order placed successfully!");
        cart = [];
        renderCart();
    } 
    catch (error) {
        console.error("Error placing order:", error);
    }
}