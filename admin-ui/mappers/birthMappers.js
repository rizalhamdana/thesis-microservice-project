function birthRecordTableMapping(data){
    var limit = data.lenght >= 5 ? 5 : data.length;
    for(var i=0;i<limit;i++){
        var tableRow = document.createElement("tr");
        
        var birthCertificateNumber = document.createElement("th");
        birthCertificateNumber.innerHTML = data[i].birth_regis_number;

        var name = document.createElement("td");
        name.innerHTML = data[i].name;

        var sex = document.createElement("td");
        sex.innerHTML = data[i].sex;      
        
        
        tableRow.appendChild(birthCertificateNumber);
        tableRow.appendChild(name);
        tableRow.appendChild(sex);

        document.getElementById("familyRecords").appendChild(tableRow);
    }
}


