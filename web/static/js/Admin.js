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