$(document).ready(function(){
   
   
    $('.modal').modal();
    $('select').formSelect();
    

    $('.modal-trigger').click(function(){
        // alert("hello");
        $("#user_form_modal").modal("open");

    });
    var currYear = (new Date()).getFullYear();
    $(".datepicker").datepicker({
        defaultDate: new Date(currYear,1,31),
        // setDefaultDate: new Date(2000,01,31),
        maxDate: new Date(currYear,12,31),
        yearRange: [1928, currYear],
        format: "yyyy-mm-dd"    
      });
   
    
    

    // $('.questionnarie-trigger').leanModal();
    // var count=0;
    
    
    function questionnarie(){
        $("#questionnarie-modal").modal("open");

       
    }
    function close_user_form(){
        $('#user_form_modal').modal("close");
    }
    $('.questionnarie-trigger').click(function(){
        questionnarie();
        close_user_form();
    });
    // $('select').material_select();

    


});

