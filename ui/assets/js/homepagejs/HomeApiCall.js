let GetMatches = function (id) {
    $.ajax({
       url:`/matches/user/${id}`,
        type:"GET",
        success:function (data) {
            // let parsedData = JSON.parse(data)
            $("#dashboard").empty();
            data.forEach(function (item) {
                $("#dashboard").append(userCard(item.UserId,item.PicturePath,item.UserName,item.Age,item.City,item.Country,item.MatchPercentage));
            })
        }

    });
}

GetMatches(1);