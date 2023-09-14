
//   ------------- numbers counter -----------------//

    $(document).ready(function(){
        $(".counter").counterUp({
         delay: 10,
         time: 1200,
        });
    });

    // ----------------------- clients carousel -------------------//

  $('.owl-carousel.client').owlCarousel({
    loop:true,
    margin:10,
    nav:false,
    dots:false,
    autoplay:true,
    autoplayTimeout:1200,
    responsive:{
        0:{
            items:2
        },
        600:{
            items:3
        },
        1000:{
            items:6
        }
    }
})


// ---------------- clients slider ---------------------//

$('.owl-carousel.testimonial').owlCarousel({
  loop:true,
  margin:10,
  nav:false,
  dots:true,
  autoplay:true,
  autoplayTimeout:8000,
  responsive:{
      0:{
          items:1
      },
      600:{
          items:2
      },
      1000:{
          items:3
      }
  }
})

// ---------------- preloader -------------------- //

var loader = document.getElementById("preloader");
window.addEventListener("load", function(){
    loader.style.display = "none";
})


// ---------------- back to top button -------------------- //

let calcScrollValue = () => {
    let scrollProgress = document.getElementById("progress");
    let porogressValue = document.getElementById("progress-value");
    let pos = document.documentElement.scrollTop;
    let calcHeight = document.documentElement.scrollHeight - document.documentElement.clientHeight;
    let scrollValue =  Math.round((pos * 100) / calcHeight);
    if (pos > 100) {
        scrollProgress.style.display = "grid";
    }
    else {
        scrollProgress.style.display = "none";
    }
    scrollProgress.addEventListener("click", () => {
        document.documentElement.scrollTop = 0;
    });
    scrollProgress.style.background = `conic-gradient(#FF6B33 ${scrollValue}%, #d7d7d7 ${scrollValue}%)`;
};
window.onscroll = calcScrollValue;
window.onload =  calcScrollValue;

