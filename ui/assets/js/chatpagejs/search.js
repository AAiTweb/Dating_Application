let id = $("#search");
id.on("keyup",function(event){
    let section = $(".section_left > p");
    let searchWord = $(this).val().toLowerCase();
    Array.from(section).forEach(function(val){
        if (! ($(val).text().toLowerCase().includes(searchWord))){
            $(val).parent().parent().hide();
        }else{
            $(val).parent().parent().show();
        }
    });
})