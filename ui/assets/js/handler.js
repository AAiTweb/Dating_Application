var questionBox=function(questionId,userQuestion,wishQuestion,answers){
   
   
    $("select").empty().html(' ');
    // var questionNumber=$(".question_number");
    // questionNumber.empty();
    // questionNumber.append()
    $(".question_number").html(questionId);
    $(".userQuestion").html(userQuestion);
    $(".wishQuestion").html(wishQuestion);

    // And add a new value
    // var value = "New value";
    $.each(answers,function(index,value){
      
        $("select").append(
            $("<option></option>").attr("value",value.answerId).text(value.answer)
          );

    });

    

    // Update the content clearing the caret
    $("select").material_select('update');
    $("select").closest('.input-field').children('span.caret').remove();
};


var answers=[];
var getSelected=function(){
  // console.log($("#own_choice").val());
  // console.log($("#wish_choice").val());
   console.log(id+1);
   var questionID=id+1;
   var ownChoiceId=parseInt($("#own_choice").val());
   var wishChoiceID=parseInt($("#wish_choice").val());
   console.log("logging from get select function");
   var answer={
     questionId:questionID,
     ownAnswerId:ownChoiceId,
     wishAnswerId:wishChoiceID
   }
   answers.push(answer);

   console.log(answers);

        

}




  
  // 2a) Whenever you do this --> add new option
  
  
  // 2b) Manually do this --> trigger custom event
  
