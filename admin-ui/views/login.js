
function showLogin(){
    var isAuthenticated = tokenCheck();
    console.log(isAuthenticated)
    if(isAuthenticated){
        window.location = '../index.html';
    }
    
}

function loginRequest(){
    console.log("LOGIN REQUEST");
    username = document.getElementById("username").value;
    password = document.getElementById("password").value;    
    login(username, password)
}