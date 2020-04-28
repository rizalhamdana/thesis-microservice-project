function getAllBirth(token){
    const url = "http://192.168.1.3:31679/birth"
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
        console.log(error.message)
        return error     
      });
    return result;
}
function getOneBirth(){

}
function insertBirth(){

}
function deleteBirth(){

}
function updateOneBirth(){

}
function verifyOneBirth(){

}