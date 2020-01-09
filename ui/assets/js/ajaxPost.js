var user_id=1;
// var questionId=id+1;

var ajaxPost=function(){
    console.log(answers);
    
    $.each(answers,function(index,value){
        var x= JSON.stringify(value);
        
        console.log(x);
        console.log("strigify the data");
        // console.log("value of")
        // console.log(value);
        var api_post_url=`/user/questionnarie/answers/${user_id}/${index}`;
        // fetch(api_post_url,{
        //     method:'POST',
        //     headers:new Headers(),
        //     body:JSON.stringify(value)

        // }).then((res)=>res.json())
        //    .then((data)=> console.log(data))
        //    .catch((err)=>console.log(err)) 
        $.ajax({
            url:api_post_url,
            contentType:"application/json",
            data:JSON.stringify(value),
            type:"POST",
            success:function(data,status){
                console.log("ajax data posted");
                console.log(data);
    
            }
        

    
    });
});
    // $.ajax({
    //     url:api_post_url,
    //     contentType:"application/json",
    //     type:"POST",
    //     success:function(data,status){


    //     }
    // })
}