// $(document).ready(function(){
//     // $('.sidenav').sidenav();
//     // alert("button");
//     // $sideNav();
//     // alert("modal");
//     $('.modal').modal();
//   });
     

$(document).ready(function(){
  $('ul.tabs').tabs();
  $('.modal').modal();
  
  // $('.modal-trigger').leanModal();
 
  $(".modal-content .exit").click(function(){
    // $(this).addClass("circle red");
    $(".modal").modal("close");

  })

  
  
});
      