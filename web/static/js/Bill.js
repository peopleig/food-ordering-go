const initialTotal = document.getElementById("initial")
const orderElement = document.getElementById("order_id")
const orderId = parseInt(orderElement.innerText)
const baseTotal = parseFloat(initialTotal.innerText);
const tipInput = document.getElementById("tip_input");
const finalTotal = document.getElementById("final_total");
const payBtn = document.getElementById("pay_btn");
tipInput.addEventListener("input", () => {
    let tip = parseFloat(tipInput.value) || 0;
    finalTotal.textContent = (baseTotal + tip);
});
payBtn.addEventListener("click", async () => {
    let tip = parseFloat(tipInput.value) || 0;
    const response = await fetch("/bill", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            order_id: orderId,
            tip: tip
        })
    });
    if (!response.ok) throw new Error("request failed");
    const result = await response.json();
    console.log(result)
    window.location.reload();
});

