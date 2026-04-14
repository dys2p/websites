function addToValue(id, addend) {
	let elem = document.getElementById(id);
	if(elem) {
		elem.value = Math.max(parseInt(elem.value) + addend, 0);
	}
}

// usage:
//
// <a role="button" onclick="copyToClipboard(this)">
//   lorem ipsum
//   <i class="fa-solid fa-copy ms-1"></i>
// </a>
var copyToClipboardIcon = null;
function copyToClipboard(elem, success, failed){
	// restore previous icon
	if(copyToClipboardIcon) {
		copyToClipboardIcon.classList.remove("fa-check");
		copyToClipboardIcon.classList.remove("fa-xmark");
		copyToClipboardIcon.classList.add("fa-copy");
	}

	copyToClipboardIcon = elem.querySelector('i');
	navigator.clipboard.writeText(elem.textContent.trim()).then(
		function() {
			copyToClipboardIcon.classList.remove("fa-copy");
			copyToClipboardIcon.classList.add("fa-check");
		},
		function() {
			copyToClipboardIcon.classList.remove("fa-copy");
			copyToClipboardIcon.classList.add("fa-xmark");
		},
	);
}

function scheduleReload() {
	setTimeout(function(){
		window.location.reload(1);
	}, 60*1000); // 60 seconds
}

document.addEventListener('DOMContentLoaded', function() {
	document.querySelectorAll(".d-js-none").forEach(e => e.remove());
	document.querySelectorAll(".d-nojs-none").forEach(e => e.classList.remove("d-nojs-none")); // style.css: .d-nojs-none { display: none; }
	document.querySelectorAll(".d-scroll-into-view").forEach(e => e.scrollIntoView({block: "nearest"}));
});
