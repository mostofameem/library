// Function to fetch and display user requests
function fetchUserRequests() {
    axios.get('http://localhost:3001/get_user?is_active=false')
        .then(response => {
            const userRequestList = document.getElementById('user-request-list');
            userRequestList.innerHTML = '';
            response.data.data.forEach(request => {
                const item = document.createElement('div');
                item.className = 'item';
                item.innerHTML = `
                    <span>${request.Id} ${request.Name} ${request.Email}</span>
                    <div>
                        <button class="accept" data-name="${request.Name}" data-email="${request.Email}" data-id="${request.Id}">✅</button>
                        <button class="reject" data-name="${request.Name}" data-email="${request.Email}" data-id="${request.Id}">❌</button>
                    </div>`;
                userRequestList.appendChild(item);
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
                const item = document.createElement('div');
                item.className = 'item';
                item.innerHTML = `
                    <span>${request.user_id} ${request.book_title}</span>
                    <div>
                        <button class="accept" data-user-id="${request.user_id}" data-book-title="${request.book_title}">✅</button>
                        <button class="reject" data-user-id="${request.user_id}" data-book-title="${request.book_title}">❌</button>
                    </div>`;
                bookRequestList.appendChild(item);
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
            axios.patch('http://localhost:3001/approve-user', {
                id: parseInt(id),
            })
                .then(res => {
                    if (res.data.status == true) {
                        alert("User Added")
                        fetchUserRequests();
                    } else {
                        alert("Operation Failed")
                    }
                })
                .catch(err => {
                    console.log(err);
                });
        });
    });

    rejectButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const id = event.target.getAttribute('data-id');
            if (!id) {
                alert("No user id and book title found");
                return;
            }
            axios.delete(`http://localhost:3001/reject-user?id=${id}`)
                .then(res => {
                    if (res.data.status == true) {
                        alert("Rejected Done")
                        fetchUserRequests();
                    } else {
                        alert("Rejected Failed")
                    }

                })
                .catch(err => {
                    console.log(err);
                });
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
                    if (res.data.status == true) {
                        alert("Borrow Request Accepted")
                        fetchBookRequests();
                    } else {
                        alert("Operation Failed")
                    }
                })
                .catch(err => {
                    console.log(err);
                });

        });
    });

    rejectButtons.forEach(button => {
        button.addEventListener('click', (event) => {
            const userId = event.target.getAttribute('data-user-id');
            const bookTitle = event.target.getAttribute('data-book-title');
            if (!userId || !bookTitle) {
                alert("No user id and book title found");
                return;
            }
            axios.delete(`http://localhost:3002/reject-request?user_id=${userId}&book_title=${bookTitle}`)
                .then(res => {
                    if (res.data.status == true) {
                        alert("Rejected Done")
                        fetchBookRequests();
                    } else {
                        alert("Rejected Failed")
                    }

                })
                .catch(err => {
                    console.log(err);
                });
        });
    });
}

// Function to alter a message
function alterMessage(status, details) {
    alert(`Request has been ${status}: ${details}`);
}

// Call the functions to fetch data when the page loads
window.onload = function () {
    fetchUserRequests();
    fetchBookRequests();
};
