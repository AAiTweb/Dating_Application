// $(document).ready(function(){
    var api_url="/user/questionnarie/questions";
    var id=-1;
   
    // var ownId=1;
    // var wishId=1;

    // var answersList=[];
    // $("#own_choice").change(function(){
    //     ownId = $(this).children("option:selected").val();   
        
    // });
    // $("#wish_choice").change(function(){
    //     wishId = $(this).children("option:selected").val();    
       
    // });
    
    // $('#own_choice').on('change', function() {
    //     console.log($(this).val());
    //     $(this).material_select();  
    // });
    
    // $('select').formSelect('methodName');
    // $('select').formSelect('methodName', paramName);
    
    
    
    // var ownId=$(".user-select")
    
    // alert("api documentation");
    // var appendAnswers=function(data){

    //     answers.append(data);

    // }
    var getQuestion=function(data,id){
        // var userQuestion=$(".user_question");
        // userQuestion.empty();
        var question=data[id];
        questionBox(question.questionId,question.UserQuestion,question.wishQuestion,question.answers);
      

    };
   
    var ajaxGetCall=function(){
        $.ajax({
            url:api_url,
            contentType:"application/json",
            type:'GET',
            success:function(data,status){
                // console.log(id);
                getQuestion(data,id);
                // id++;
            },
        });

    }; 
    var ajaxGetUser=function(){
        $.ajax({
            url:"/user/profile/get/",
            contentType:"application/json",
            type:'GET',
            success:function(data,status){
                // console.log(id);
                // getQuestion(data,id);
                // id++;
            },
        });

    } 
    $(".prev").click(function(){
        
        // alert("prev clicked");
        
        id--;
        if(id >=0){
            ajaxGetCall();
        }
        else if(id == -1){
            $('#user_form_modal').openModal();
            $("#questionnarie-modal").closeModal();
        }
        else{
            alert("no turning back");
        }
        if( answers.length > 0){
           answers.pop();
           console.log("poping");
           console.log(answers);
        }
        
       


    });
    // $("#own_choice").change(function(){
        
    // });
   
    $(".next-question").click(function(){
        getSelected();
        // if (id=5){
        //     ajaxPost();
        // }
        if (id <5){
            ajaxGetCall();
        }
        else if(id==5){
            ajaxPost();
            
        }   
        else{
            alert("done");
        }
      
        // var selected=1;
        // selected=$("#own_choice").val();
        // console.log(selected);
        
        
       
        
        // var answer={
        //     questionId:id,
        //     userChoiceId:ownId,
        //     wishChoiceId:wishId

        // }
        
        // answersList.append(answer);
        
        id++;
      
        
        // $.ajax({
        //     url:api_url,
        //     contentType:"application/json",
        //     type:'GET',
        //     success:function(data,status){
        //         getQuestion(data,id);
        //         id++;
        //     },
        // });

    });
    $(".questionnarie-trigger").click(function(){
        getUserForm();
        // $("#user_form").submit();
        ajaxGetCall();
        ajaxFormPost();
        id++;
        
        // $.ajax({
        //     url:api_url,
        //     contentType:"application/json",
        //     dataType:"json",
        //     type:'GET',
        //     success:function(data,status){
        //         getQuestion(data,0);
        //     },
        // });

    });



    // console.log(answersList);
    // $.ajax({
    //     url:api_url,
    //     contentType:"application/json",
    //     dataType:'json',
    //     success:function(data){
    //         console.log(data);
    //     }
    // });
// });