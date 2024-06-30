// Function to fetch and display user requests
function fetchUserRequests() {
    axios.get('http://localhost:3001/get_user?is_active=false')
        .then(response => {
            const userRequestList = document.getElementById('user-request-list');
            userRequestList.innerHTML = '';
            response.data.data.forEach(request => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${request.Name}</td>
                    <td>${request.Email}</td>
                    <td class="actions">
                        <button class="accept" data-name="${request.Name}" data-email="${request.Email}" data-id="${request.Id}">✔</button>
                        <button class="reject" data-name="${request.Name}" data-email="${request.Email}" data-id="${request.Id}">✖</button>
                    </td>`;
                userRequestList.appendChild(row);
            });

            addUserEventListeners();
        })
        .catch(error => console.error('Error fetching user requests:', error));
}

// Function to fetch and display book requests
function fetchBookRequests() {
    axios.get('http://localhost:3002/get-borrow-request?is_active=false')
        .then(response => {
            const bookRequestList = document.getElementById('book-request-list');
            bookRequestList.innerHTML = '';
            response.data.data.forEach(request => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${request.user_id}</td>
                    <td>${request.book_title}</td>
                    <td class="actions">
                        <button class="accept" data-user-id="${request.user_id}" data-book-title="${request.book_title}">✔</button>
                        <button class="reject" data-user-id="${request.user_id}" data-book-title="${request.book_title}">✖</button>
                    </td>`;
                bookRequestList.appendChild(row);
            });

            addBookEventListeners();
        })
        .catch(error => console.error('Error fetching book requests:', error));
}

// Function to add event listeners to user request buttons
function addUserEventListeners() {
    const acceptButtons = document.querySelectorAll('#user-request .accept');
    const rejectButtons = document.querySelectorAll('#user-request .reject');

    acceptButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const id = event.target.getAttribute('data-id');
            axios.patch('http://localhost:3001/approve-user', { id: parseInt(id) })
                .then(res => {
                    if (res.data.status) {
                        alert("User Added");
                        fetchUserRequests();
                    } else {
                        alert("Operation Failed");
                    }
                })
                .catch(err => console.log(err));
        });
    });

    rejectButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const id = event.target.getAttribute('data-id');
            axios.delete(`http://localhost:3001/reject-user?id=${id}`)
                .then(res => {
                    if (res.data.status) {
                        alert("Rejected Done");
                        fetchUserRequests();
                    } else {
                        alert("Rejected Failed");
                    }
                })
                .catch(err => console.log(err));
        });
    });
}

// Function to add event listeners to book request buttons
function addBookEventListeners() {
    const acceptButtons = document.querySelectorAll('#book-request .accept');
    const rejectButtons = document.querySelectorAll('#book-request .reject');

    acceptButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const userId = event.target.getAttribute('data-user-id');
            const bookTitle = event.target.getAttribute('data-book-title');
            axios.patch('http://localhost:3002/approve-request', {
                user_id: parseInt(userId),
                book_title: bookTitle
            })
                .then(res => {
                    if (res.data.status) {
                        alert("Borrow Request Accepted");
                        fetchBookRequests();
                    } else {
                        alert("Operation Failed");
                    }
                })
                .catch(err => console.log(err));
        });
    });

    rejectButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const userId = event.target.getAttribute('data-user-id');
            const bookTitle = event.target.getAttribute('data-book-title');
            axios.delete(`http://localhost:3002/reject-request?user_id=${userId}&book_title=${bookTitle}`)
                .then(res => {
                    if (res.data.status) {
                        alert("Rejected Done");
                        fetchBookRequests();
                    } else {
                        alert("Rejected Failed");
                    }
                })
                .catch(err => console.log(err));
        });
    });
}

// Call the functions to fetch data when the page loads
window.onload = function () {
    fetchUserRequests();
    fetchBookRequests();
};

function fetchBookDetails() {
    window.location.href = "home.html";
}