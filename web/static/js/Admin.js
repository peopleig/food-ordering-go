document.querySelectorAll('.approve-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
        const userId = btn.getAttribute('data-user-id');
        const res = await fetch(`/admin/${userId}`, { method: "PATCH" }); 
        if (res.ok) {
            alert("Approved!")
            window.location.href = '/admin';
        } else {
            alert("Approval failed.");
        }
    });
});

document.querySelectorAll('.decline-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
        const userId = btn.getAttribute('data-user-id');
        const res = await fetch(`/admin/${userId}`, { method: "DELETE" }); 
        if (res.ok) {
            alert("Deleted Request!")
            window.location.href = '/admin';
        } else {
            alert("Decline failed.");
        }
    });
});

document.querySelectorAll('.approve-payment-btn').forEach(btn => {
    btn.addEventListener('click', async () => {
        const orderId = btn.getAttribute('data-order-id');
        const res = await fetch(`/admin/payment/${orderId}`, { method: "POST" }); 
        if (res.ok) {
            alert("Approved Payment!")
            window.location.href = '/admin';
        } else {
            alert("Approval failed.");
        }
    });
});

const rows = document.querySelectorAll("#orders_table_body tr");
const filter_buttons = document.querySelectorAll(".order_buttons");
filter_buttons.forEach(btn => {
    btn.addEventListener("click", () => {
        const status = btn.dataset.status;
        rows.forEach(row => {
            const rowStatus = row.dataset.status;
            if (status === "all" || rowStatus === status) {
                row.classList.remove("d-none");
            } else {
                row.classList.add("d-none");
            }
        });
        filter_buttons.forEach(b => b.classList.remove("active"));
        btn.classList.add("active");
    });
});

const item_filter_buttons = document.querySelectorAll(".item_buttons");
item_filter_buttons.forEach(button => {
    button.addEventListener("click", () => {
        const filter = button.dataset.chef;
        document.querySelectorAll("#items_table_body tr").forEach(row => {
            const chef_id = parseInt(row.dataset.chef_id);
            row.style.display = (
                filter === "all" ||
                (filter === "unassigned" && chef_id === 1) ||
                (filter === "assigned" && chef_id !== 1)
            ) ? "table-row" : "none";
        });
        item_filter_buttons.forEach(b => b.classList.remove("active"));
        button.classList.add("active");
    });
});
