$(document).ready(function(){
    // alert("hello modal");
    $('.modal-trigger').leanModal();
    
    

    // $('.questionnarie-trigger').leanModal();
    var count=0;
    
    function getJsonData(id){

    }
    function questionnarie(){
        $("#questionnarie-modal").openModal();

       
    }
    function close_user_form(){
        $('#user_form_modal').closeModal();
    }
    $('.questionnarie-trigger').click(function(){
        questionnarie();
        close_user_form();
    });
    $('select').material_select();
    $("#nextQuestion").click(function(){
        alert("next question");
    });



});

