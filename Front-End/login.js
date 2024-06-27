const weapper = document.querySelector('.wrapper');

function login() {
  // change page on buttonclick
  window.location.href = "registration.html";
}
// Add event listener for the close button
document.getElementById("closeButton").addEventListener("click", function () {

  window.location.href = "start.html"; // Replace with the URL of the page you want to open
});
function Login() {

  const email = document.getElementById('Email').value;
  const password = document.getElementById('Password').value;
  // Sending user data in the body of the POST request
  axios.post('http://localhost:3001/login', {
    email: email,
    password: password
  })
    .then(res => {
      if (res.data.status == true) {
        if (res.data.message == "User") {
          // Store the JWT token in localStorage
          localStorage.setItem('jwtToken', res.data.data);
          console.log(res.data)
          alert("Log IN")
          // Redirect to home page
          window.location.href = "home.html";
        }else{

        }
      } else {
        alert("Wrong email/password");
        window.location.href = "start.html";
      }
    })
    .catch(err => {
      console.log(err);
    });
}