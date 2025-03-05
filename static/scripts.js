// usage:
//
// <a role="button" onclick="copyToClipboard(this)">
//   lorem ipsum
//   <i class="fa-solid fa-copy ms-1"></i>
// </a>
function copyToClipboard(elem, success, failed){
	navigator.clipboard.writeText(elem.textContent.trim()).then(
		function() {
			if(success) {
				success.classList.remove("d-none");
				setTimeout(function() {
					success.classList.add("d-none");
				}, 3000)
			}
		},
		function() {
			if(failed){
				failed.classList.remove("d-none");
				setTimeout(function() {
					failed.classList.add("d-none");
				}, 3000)
			}
		},
	);
}


function scheduleReload() {
	setTimeout(function(){
		window.location.reload(1);
	}, 5*60*1000); // 5 minutes
}

document.addEventListener('DOMContentLoaded', function() {
	document.querySelectorAll(".d-js-none").forEach(e => e.remove());
	document.querySelectorAll(".d-nojs-none").forEach(e => e.classList.remove("d-nojs-none")); // style.css: .d-nojs-none { display: none; }
	document.querySelectorAll(".d-scroll-into-view").forEach(e => e.scrollIntoView({block: "nearest"}));
});
