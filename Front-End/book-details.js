
let book;

document.addEventListener("DOMContentLoaded", function () {
    // Retrieve the title from the URL parameters
    const params = new URLSearchParams(window.location.search);
    const title = params.get('title');

    if (title) {
        // Make an API call to fetch book details based on the title
        const query = `http://localhost:3002/showbooks?title=${encodeURIComponent(title)}`;
        console.log(`Sending request to: ${query}`);

        axios.get(query)
            .then(res => {
                console.log('Response received:', res); // Log the entire response
                if (res.data && Array.isArray(res.data.data)) {
                    book = res.data.data[0]; // Assuming the response contains a single book
                    displayBookDetails(book);
                } else {
                    console.error('Unexpected response format:', res);
                }
            })
            .catch(err => {
                console.error('Error fetching data:', err);
            });
    }
});

function displayBookDetails(book) {
    const bookDetailsDiv = document.getElementById('book-details');
    bookDetailsDiv.innerHTML = `
        <p><strong>ISBN:</strong> ${book.isbn}</p>
        <p><strong>Title:</strong> ${book.title}</p>
        <p><strong>Page:</strong> ${book.total_page}</p>
        <p><strong>Author:</strong> ${book.author}</p>
        <p><strong>Genres:</strong> ${book.genres}</p>
        <p><strong>Publication Date:</strong> ${book.publication_date}</p>
        <p><strong>Next Available:</strong> ${book.next_available}</p>
        <p><strong>Is Active:</strong> ${book.is_active}</p>
      `;
}

function BorrowBook() {
    // Get the JWT token from localStorage
    const token = localStorage.getItem('jwtToken');
    console.log(token);

    if (!token) {
        alert('You need to be logged in to borrow a book.');
        return;
    }
    // Get the current book details from localStorage
    const currentBook = book
    if (!currentBook) {
        alert('No book selected.');
        return;
    }
    // Retrieve the selected return date from the input field
    const date = document.getElementById('Date').value;

    if (!date) {
        alert('No Date selected.');
        return;
    }
    // Set the request body
    const requestBody = {
        book_title: currentBook.title,
        total_page:currentBook.total_page,
        Return_date: date // Set the desired return date
    };
    console.log(requestBody)

    // Make the API call to borrow the book
    axios.post('http://localhost:3002/borrow', requestBody, {
        headers: {
            Authorization: `${token}`
        }
    })
        .then(response => {
            alert('Book borrowed successfully.');
            window.location.href = "home.html";
        })
        .catch(error => {
            console.error(error);
        });

}