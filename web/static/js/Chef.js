const my_id = parseInt(document.querySelector("input").dataset.user_id);
const filter_buttons = document.querySelectorAll(".filter-btn")
filter_buttons.forEach(button => {
    button.addEventListener("click", () => {
        const filter = button.dataset.chef;
        document.querySelectorAll("#orders_table_body tr").forEach(row => {
            const chef_id = parseInt(row.dataset.chef_id);
            row.style.display = (
                filter === "all" ||
                (filter === "unassigned" && chef_id === 1) ||
                (filter === "mine" && chef_id === my_id)
            ) ? "table-row" : "none";
        });
        filter_buttons.forEach(b => b.classList.remove("active"));
        button.classList.add("active");
    });
});


document.querySelectorAll('.approve-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
        const orderId = btn.getAttribute('data-order_id');
        const itemId = btn.getAttribute('data-item_id');
        const response = await fetch("/chef", {
            method: "PATCH",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                chefId : my_id,
                orderId: parseInt(orderId),
                itemId: parseInt(itemId),
            })
        });
        if (!response.ok) throw new Error("request failed");
        const result = await response.json();
        console.log(result)
        window.location.reload();
    });
});

document.querySelectorAll('.done-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
        const orderId = btn.getAttribute('data-order_id');
        const itemId = btn.getAttribute('data-item_id');
        const goAhead = await confirmComplete();
        if (!goAhead) return;

        try {
            const response = await fetch("/chef", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    chefId: my_id,
                    orderId: parseInt(orderId),
                    itemId: parseInt(itemId),
                })
            });

            if (!response.ok) throw new Error("Request failed");

            const result = await response.json();
            console.log(result);
            window.location.reload();
        } catch (err) {
            console.error("Error completing item:", err);
        }
    });
});


function confirmComplete() {
    return new Promise((resolve) => {
        confirmPopup.classList.remove("d-none");
        generateBtn.onclick = () => {
            confirmPopup.classList.add("d-none");
            resolve(true);
        };
        cancelBtn.onclick = () => {
            confirmPopup.classList.add("d-none");
            resolve(false);
        };
    });
}


const confirmPopup = document.getElementById("confirm-popup");
const cancelBtn = document.getElementById("cancel-confirm");
const generateBtn = document.getElementById("generate-confirm");

