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
function copyToClipboard(elem, success, failed){
	let icon = elem.querySelector('i');
	navigator.clipboard.writeText(elem.textContent.trim()).then(
		function() {
			icon.classList.remove("fa-copy");
			icon.classList.add("fa-check");
			setTimeout(function() {
				icon.classList.remove("fa-check");
				icon.classList.add("fa-copy");
			}, 3000)
		},
		function() {
			icon.classList.remove("fa-copy");
			icon.classList.add("fa-xmark");
			setTimeout(function() {
				icon.classList.remove("fa-xmark");
				icon.classList.add("fa-copy");
			}, 3000)
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
