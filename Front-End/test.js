document.addEventListener("DOMContentLoaded", function() {
  const books = ["Book 1", "Book 2", "Book 3", "Book 4", "Book 5"];
  const bookList = document.getElementById("book-list");

  books.forEach(book => {
      const li = document.createElement("li");
      li.textContent = book;
      bookList.appendChild(li);
  });
});
