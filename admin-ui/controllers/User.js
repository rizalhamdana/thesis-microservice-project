function login(username, password){
    const data = {
        username: username,
        password: password
    };
    const url = "http://192.168.1.3:31679/auth"
    const other_params = {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },
        body: JSON.stringify(data),
        mode: "cors"
      };
      fetch(url, other_params)
        .then(function(response) {
          if (response.ok) {
            console.log(response);
            return response.text();
          } else if(response.status==404) {
            console.log(response);
            throw new Error("Invalid Credentials");
          } else {
            console.log(response);
            throw new Error("Could Not Reach API");
          }
        })
        .then(function(data) {
            jsonData = JSON.parse(data)
            document.cookie = "Token=" + jsonData["token"];
            document.location = "../index.html";
            // console.log(document.cookie);
        })
        .catch(function(error) {
          document.getElementById("message").innerHTML = error.message;
        });
      return true;
}

function logout(){
    document.cookie = "Token=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
    window.location = '../login.html';
}


function tokenCheck(){
    token = getCookie("Token");
    return token == undefined || token == '' ? false : true;
}

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
  }

function isAuth(){
  isAuthenticated = tokenCheck();
  if(!isAuthenticated){
       window.location = "../login.html"
  }
}