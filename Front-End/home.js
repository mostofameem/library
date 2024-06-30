let books = []; // Declare an array to store the response data
document.addEventListener("DOMContentLoaded", function () {
    Get_Book_List();
});
function Get_Book_List() {
    const title = document.getElementById('title').value;
    const author = document.getElementById('author').value;
    const genre = document.getElementById('category').value;

    let query = "http://localhost:3002/showbooks?";
    let cnt = 0;
    if (title) {
        query += `title=${encodeURIComponent(title)}`;
        cnt++;
    }
    if (author) {
        if (cnt) query += `&`;
        query += `author=${encodeURIComponent(author)}`;
        cnt++;
    }
    if (genre) {
        if (cnt) query += `&`;
        query += `genres=${encodeURIComponent(genre)}`;
    }

    console.log(`Sending request to: ${query}`); // Log the query URL

    axios.get(query)
        .then(res => {
            console.log('Response received:', res); // Log the entire response
            if (res.data && Array.isArray(res.data.data)) {
                books = res.data.data; // Store the response array in the books array
                updateBookList(books); // Update the book list in the DOM
            } else {
                console.error('Unexpected response format:', res);
            }
        })
        .catch(err => {
            console.error('Error fetching data:', err);
        });
}

function updateBookList(books) {
    const bookList = document.getElementById("book-list");
    bookList.innerHTML = ''; // Clear the current list

    books.forEach(book => {
        const li = document.createElement("li");
        li.className = 'book-block';
        li.textContent = book.title; // Assuming 'title' is a property of each book object
        li.addEventListener('click', () => {
            // Navigate to the book details page and pass the book title via URL parameter
            const bookDetailsUrl = `book-details.html?title=${encodeURIComponent(book.title)}`;
            window.location.href = bookDetailsUrl;
        });
        bookList.appendChild(li);
    });
}
document.getElementById('logout').addEventListener('click', function(event) {
    // Prevent the default link action
    event.preventDefault();

    // Remove JWT token from localStorage
    localStorage.removeItem('jwtToken');

    // Navigate to the login page
    window.location.href = 'login.html';
});
