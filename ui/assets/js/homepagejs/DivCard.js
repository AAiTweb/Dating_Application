let userCard = function(userid,imagename,username,age,city,country,matchPercentage){
    let colContainer = $(`<div class="col m4 l2"></div>`);
    let card = $(`<div class="card"></div>`);
    let card_image = $(`<div class="card-image"></div>`);
    let image = ` <a href="/profile/${userid}"><img height="200px" src="../assets/images/${imagename}" /></a>`;
    card_image.append(image);
    let card_content = $(`<div class="card-content"></div>`);
    let div_content_p1 = `<p><b><span class="uname">${username}</span>, <span class="age">${age}</span></b></p>`;
    let div_content_p2 = `<p><span class="city">${city}</span>, <span class="country">${country}</span> </p>`;
    let div_content_p3 = `<div class="progress"><div class="determinate" style="width: ${matchPercentage}%"></div></div>`;
    card_content.append(div_content_p1,div_content_p2,div_content_p3);
    let cardAction = `<div class="card-action"><a class="secondLink bothLink"></a> </div>`;
    card.append(card_image,card_content,cardAction);
    colContainer.append(card);
    return colContainer;
}
// $("#dashboard").append(userCard(1,"default.jpg","abebe",12,"Nazreth","Ethiopia",89));
