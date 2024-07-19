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

// HTML noscript tag does not work with NoScript browser addon, so we apply class noscript and then call this function
function removeNoscript() {
	document.querySelectorAll(".noscript").forEach(e => e.remove());
}
