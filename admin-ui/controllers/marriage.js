function getAllMarriage(token){
    const url = "http://192.168.1.3:31679/married"
    const other_params = {
        method: "GET",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          "Token": token
        },
        mode: "cors"
      };
    var result = fetch(url, other_params)
      .then(function(response) {
        if (response.ok) {
          
          return response.text();
        } else if(response.status==401) {   
          throw new Error("Unauthorized");
        } else {  
          throw new Error("Could Not Reach API");
        }
      })
      .then(function(data) {
            jsonData = data == null || data == '' || data == undefined ? [] : JSON.parse(data);
            
            return jsonData;

      })
      .catch(function(error) {
        return error     
      });
    return result;
}
function getOneMarriage(token, certificateNumber){
  const url = "http://192.168.1.3:31679/married/"+certificateNumber;
  const other_params = {
    method: "GET",
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        "Token": token
    },
    mode: "cors"
  };
var result = fetch(url, other_params)
  .then(function(response) {
    if (response.ok) {
    
      return response.text();
    } else if(response.status==401) {
      throw new Error("Unauthorized");
    } else if(response.status==404){
      throw new Error("Not Found")
    } else {
      throw new Error("Could Not Reach API");
    }
  })
  .then(function(data) {
        jsonData = JSON.parse(data)
        return jsonData
  })
  .catch(function(error) {
    console.log(error.message)
    return error     
  });
return result;

}
function insertMarriage(){

}
function deleteMarriage(){

}
function updateOneMarriage(){

}
function verifyOneMarriage(token, regisNumber){
  const url = "http://192.168.1.3:31679/married/verif/"+regisNumber;
  const other_params = {
    method: "PUT",
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        "Token": token
    },
  };
var result = fetch(url, other_params)
  .then(function(response) {
    if (response.ok) {
      return response.text();
    } else if(response.status==401) {
      throw new Error("Unauthorized");
    } else if(response.status==404){
      throw new Error("Not Found")
    } else {
      throw new Error("Could Not Reach API");
    }
  })
  .then(function(data) {
        console.log(data)
        return data
  })
  .catch(function(error) {
    console.log(error.message)
    return error     
  });
return result;
}