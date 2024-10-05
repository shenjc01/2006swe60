$(document).ready(function(){

    let $btns = $('.project-area .button-group button');
    $btns.click(function(e) {

        $('.project-area .button-group button').removeClass('active');
        e.target.classList.add('active');

        let selector = $(e.target).attr('data-filter');
        $('.project-area .grid').isotope({
            filter:selector
        });

        return false;
    })

    $('.project-area .button-group #btn1').trigger('click');

    $('.project-area .grid .test-popup-link').magnificPopup({
        type: 'image',
        gallery:{enabled:true}
      });
/*
    let nav_offset_top = $('.header_area').height() + 50;

    let navbar = document.getElementById("main-nav");
    let header = document.getElementById("header");
    let navPos = navbar.getBoundingClientRect().top;
    
    window.addEventListener("scroll", e => {
    let viewportHeight = window.innerHeight;
    let scrollPos = window.scrollY;
     if (scrollPos >= navPos) {
        navbar.classList.add('sticky');
        header.classList.add('navbarOffsetMargin');
    } else {
        navbar.classList.remove('sticky');
        header.classList.remove('navbarOffsetMargin');
    }
*/
    let navbar = document.getElementById("list-nav");
    let viewportHeight = window.innerHeight;
    let navHeight = document.getElementById("list-nav").offsetHeight;

    let navbarLinks = document.querySelectorAll("a.nav-link");

    window.addEventListener("scroll", e => {
        scrollpos = window.scrollY;
        navbarLinks.forEach(link => {
            let section = document.querySelector(link.hash);
            if (section.offsetTop <= scrollpos+75 &&
            section.offsetTop + section.offsetHeight > scrollpos+75) {
            link.classList.add("active");
            } else {
            link.classList.remove("active");
            }
        });
    });

    /* $('#list-nav').on('click','a', function (e) {
        e.preventDefault();
        $(this).parents('#list-nav').find('.active').removeClass('active').end().end().addClass('active');
        $(activeTab).show();
        }); */

    
});