document.getElementById('bookForm').addEventListener('submit', function(e) {
  e.preventDefault();
  
  const title = document.getElementById('title').value;
  const author = document.getElementById('author').value;
  const isbn = document.getElementById('isbn').value;
  const category = document.getElementById('category').value;

  addBookToList(title, author, isbn, category);

  document.getElementById('bookForm').reset();
});

function addBookToList(title, author, isbn, category) {
  const bookList = document.getElementById('bookList');

  const bookItem = document.createElement('div');
  bookItem.className = 'book-item';
  bookItem.innerHTML = `
      <h3>${title}</h3>
      <p><strong>Author:</strong> ${author}</p>
      <p><strong>ISBN:</strong> ${isbn}</p>
      <p><strong>Category:</strong> ${category}</p>
  `;

  bookList.appendChild(bookItem);
}
