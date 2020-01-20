// $(document).ready(function(){
//     // $('.sidenav').sidenav();
//     // alert("button");
//     // $sideNav();
//     // alert("modal");
//     $('.modal').modal();
//   });
     

$(document).ready(function(){
  
  $('.modal').modal();
  var currYear = (new Date()).getFullYear();
  

  $(".datepicker").datepicker({
      defaultDate: new Date(currYear,1,31),
      // setDefaultDate: new Date(2000,01,31),
      maxDate: new Date(currYear,12,31),
      yearRange: [1928, currYear],
      format: "yyyy-mm-dd"    
    });
    $('ul.tabs').tabs();
  
  // $('.modal-trigger').leanModal();
 
  $(".modal-content .exit").click(function(){
    // $(this).addClass("circle red");
    $(".modal").modal("close");

  });
//   $('input[type="file"]').change(function(e){
//     var fileName = e.target.files[0].name;
//     alert('The file "' + fileName +  '" has been selected.');
// });

  

  
  
});
      