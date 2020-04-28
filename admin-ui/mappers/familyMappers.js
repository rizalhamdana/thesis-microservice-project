function familyRecordsTableMapping(data){
    var limit = data.lenght >= 5 ? 5 : data.length;
    for(var i=0;i<limit;i++){
        var tableRow = document.createElement("tr");
        
        var familyCardNumber = document.createElement("th");
        familyCardNumber.innerHTML = data[i].family_card_number;

        var headOfHousehold = document.createElement("td");
        headOfHousehold.innerHTML = data[i].head_of_household;

        var familyMembersCount = document.createElement("td");
        familyMembersCount.innerHTML = data[i].family_members.length;        
        
        var address = document.createElement("td");
        address.innerHTML = data[i].address;

        tableRow.appendChild(familyCardNumber);
        tableRow.appendChild(headOfHousehold);
        tableRow.appendChild(familyMembersCount);
        tableRow.appendChild(address);

        document.getElementById("familyRecords").appendChild(tableRow);

    }
}
