// usage:
//
// <a role="button" onclick="copyToClipboard(this)">
//   lorem ipsum
//   <i class="fa-solid fa-copy ms-1"></i>
// </a>
function copyToClipboard(elem){
	navigator.clipboard.writeText(elem.textContent.trim());
}

function scheduleReload() {
	setTimeout(function(){
		window.location.reload(1);
	}, 5*60*1000); // 5 minutes
}

document.addEventListener('DOMContentLoaded', function() {
	document.querySelectorAll(".d-js-none").forEach(e => e.remove());
	document.querySelectorAll(".d-nojs-none").forEach(e => e.classList.remove("d-nojs-none")); // style.css: .d-nojs-none { display: none; }
});
