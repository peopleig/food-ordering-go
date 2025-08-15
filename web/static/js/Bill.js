const initialTotal = document.getElementById("initial")
const orderElement = document.getElementById("order_id")
const orderId = parseInt(orderElement.innerText)
const baseTotal = parseFloat(initialTotal.innerText);
const tipInput = document.getElementById("tip_input");
const finalTotal = document.getElementById("final_total");
const payBtn = document.getElementById("pay_btn");
tipInput.addEventListener("input", () => {
    let tip = parseFloat(tipInput.value) || 0;
    if (tip < 0) {
        tipInput.value = 0;
        tip = 0;
    }
    finalTotal.textContent = (baseTotal + tip);
});
payBtn.addEventListener("click", async () => {
    let tip = parseFloat(tipInput.value) || 0;
    let url = "/bill/"+orderElement.innerText
    const response = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            order_id: orderId,
            tip: tip
        })
    });
    if (!response.ok) {
        alert("Unable to complete Payment. Please try again!")
        throw new Error(`Server returned ${response.status}`);
    }
    const data = await response.json();
    if (data.redirect) {
        window.location.href = data.redirect;
    }
});

