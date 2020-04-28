function showDashboard(){
   isAuth();
   token = getCookie("Token")
   citizens = getAllCitizen(token)
   citizens.then(data => {
        if (data instanceof Error){
            document.getElementById("numberCitizen").innerHTML = "-";
            return
        }
        console.log(data.length)
        var citizensCount = Object.keys(data).length
        document.getElementById("numberCitizen").innerHTML = citizensCount;
        citizenRecordsTableMapping(data)
   }).catch(error=>{
    console.log(error)
   })

   marriages = getAllMarriage(token);
   marriages.then(data => {
    if (data instanceof Error){
        document.getElementById("numberMarriage").innerHTML = "-"
        return
    }
    var marriagesCount = Object.keys(data).length
    document.getElementById("numberMarriage").innerHTML = marriagesCount;
    marriageRecordsTableMapping(data);
   });
   birth = getAllBirth(token);
   birth.then(data => {
    if (data instanceof Error){
        document.getElementById("numberBirth").innerHTML = "-";
        return
    }
    var birthCount = Object.keys(data).length;
    document.getElementById("numberBirth").innerHTML = birthCount;
    birthRecordTableMapping(data);
   });
   
}
