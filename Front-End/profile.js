document.getElementById('editUserInfo').addEventListener('click', function() {
    // Handle edit user info functionality here
    alert('Edit user info functionality is not implemented yet.');
});

document.getElementById('bookFilter').addEventListener('change', function() {
    const filter = this.value;
    const bookRows = document.querySelectorAll('.book-row');
    
    bookRows.forEach(row => {
        // Logic to filter books based on the selected filter
        // For simplicity, showing all books as an example
        if (filter === 'all') {
            row.style.display = 'flex';
        } else if (filter === 'current') {
            // Implement logic to filter current reading books
            row.style.display = 'flex';
        }
    });
});

document.querySelectorAll('.book-row').forEach(row => {
    row.addEventListener('click', function() {
        const bookName = this.getAttribute('data-book-name');
        alert('Book name: ' + bookName);
    });
});

// Fetch user info from API using Axios
const jwtToken = localStorage.getItem('jwtToken'); // Replace with your actual JWT token
async function fetchUserInfo() {
    try {
        const response = await axios.get('http://localhost:3001/get_user', {
            headers: {
                Authorization: `${jwtToken}`
            }
        });
        const data = response.data;
        if (data.status && data.data.length > 0) {
            const user = data.data[0];
            document.getElementById('userName').textContent = user.Name;
            document.getElementById('userEmail').textContent = user.Email;
        } else {
            alert('Failed to fetch user info');
        }
    } catch (error) {
        console.error('Error fetching user info:', error);
        alert('Error fetching user info');
    }
 }

// Fetch book list from API using Axios
// Function to fetch book list from API
async function fetchBookList() {
    try {
        const response = await axios.get('http://localhost:3002/profile', {
            headers: {
                Authorization: `${jwtToken}`
            }
        });
        const data = response.data;
        if (data.status && data.data.length > 0) {
            const bookTable = document.getElementById('bookTable');
            bookTable.innerHTML = ''; // Clear existing rows
            data.data.forEach(book => {
                const bookRow = document.createElement('div');
                bookRow.className = 'book-row';
                bookRow.setAttribute('data-book-title', book.book_title);
                bookRow.setAttribute('data-return-status', book.return_status);
                bookRow.innerHTML = `
                    <div class="book-details">
                        <h3>${book.book_title}</h3>
                    </div>
                    <div class="book-progress">
                        <p>Progress: ${((book.page_readed / book.total_page) * 100).toFixed(2)}%</p>
                    </div>
                    <div class="book-return-date">
                        <p>Issue Date: ${new Date(book.issue_date).toLocaleDateString()}</p>
                    </div>
                `;
                bookTable.appendChild(bookRow);

                // Add click event listener to navigate to book details page
                bookRow.addEventListener('click', function() {
                    // Navigate to book details page with book title as URL parameter
                    const bookTitle = encodeURIComponent(book.book_title);
                    window.location.href = `borrow-book-details.html?title=${bookTitle}`;
                });
            });
        } else {
            alert('Failed to fetch book list');
        }
    } catch (error) {
        console.error('Error fetching book list:', error);
        alert('Error fetching book list');
    }
}
fetchUserInfo();
fetchBookList();



