const order_filters = document.querySelectorAll(".order-filter");
order_filters.forEach(button => {
    button.addEventListener("click", () => {
        const status = button.dataset.status;
        if (status == "all") {
            document.querySelectorAll(".orders-card").forEach(card => {
                card.classList.remove("d-none");
                card.classList.add("d-inline-block");
            })
        }
        else {
            document.querySelectorAll(".orders-card").forEach(card => {    
                if (card.dataset.status === status) {
                    card.classList.remove("d-none");
                    card.classList.add("d-inline-block");
                } else {
                    card.classList.remove("d-inline-block");
                    card.classList.add("d-none");
                }
            });
        }
        order_filters.forEach(b => b.classList.remove("active"));
        button.classList.add("active");
    });
});