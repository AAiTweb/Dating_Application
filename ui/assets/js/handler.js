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
      // console.log(answers[-1].answer);
      
      
        // $("select").append(
        //     $("<option></option>").attr("value",value.answerId).text(value.answer)
        //   );
        var option=$("<option></option>").attr("value",value.answerId).text(value.answer);
        // $("select").closest('.input-field').children('span.caret').remove();
        $('select').formSelect().append(
          option
        );
        // $('select').formSelect().append($('<option>'+value.answer+'</option>'))
        // console.log(value);


    });$('select').formSelect().append(
      "<option>eeee</option>"
    );

    // $('select').formSelect().append(
    //   $("<option></option>").attr("value",answers[-1].answerId).text(answers[-1].answer)
    // );
    

    // Update the content clearing the caret
    // $("select").formSelect('update');
    // 
};


var answers=[];
var  fromValue={};
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
var getUserForm=function(){
  var fName=$("#fName").val();
  var lName=$("#lName").val();
  var country=$("#country").val();
  var city=$("#city").val();
  var dob=$("#dob").val();
  var sex=$("#sex").val();
  var bio=$("#bio").val();
  // var timeZone = dob.getTimezoneOffset();
  // console.log(timeZone);
  fromValue={
    firstName:fName,
    lastName:lName,
    country:country,
    city:city,
    // dob:dob,
    bio:bio,
    sex:sex,
    dob:dob
  }




}



  
  // 2a) Whenever you do this --> add new option
  
  
  // 2b) Manually do this --> trigger custom event
  
