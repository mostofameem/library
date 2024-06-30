document.addEventListener('DOMContentLoaded', function() {
    const urlParams = new URLSearchParams(window.location.search);
    const title = urlParams.get('title');
    const jwtToken = localStorage.getItem('jwtToken'); // Replace with your actual JWT token

    if (title) {
        // Make an API call to fetch book details based on the title
        const query = `http://localhost:3002/borrow-details?book_title=${encodeURIComponent(title)}`;
        console.log(`Sending request to: ${query}`);

        axios.get(query, {
            headers: {
                Authorization: `${jwtToken}`
            }
        })
        .then(res => {
            console.log('Response received:', res); // Log the entire response
            if (res.data && res.data.data) {
                const book = res.data.data; // Directly access the book details
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
    const issueDate = new Date(book.issue_date).toLocaleDateString();
    const returnDate = new Date(book.return_date).toLocaleDateString();

    bookDetailsDiv.innerHTML = `
        <p><strong>Title:</strong> ${book.book_title}</p>
        <p><strong>Pages Read:</strong> ${book.page_readed}</p>
        <p><strong>Total Pages:</strong> ${book.total_page}</p>
        <p><strong>Issue Date:</strong> ${issueDate}</p>
        <p><strong>Return Date:</strong> ${returnDate}</p>
        <p><strong>Return Status:</strong> ${book.return_status === "true" ? "Returned" : "Not Returned"}</p>
        <p><strong>Is Active:</strong> ${book.is_active ? "Yes" : "No"}</p>
    `;
}
