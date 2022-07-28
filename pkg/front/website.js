document.getElementById('updload-website').addEventListener('click', openDialog);
document.getElementById('fileid').addEventListener('change', handleFileSelect, true);

getWallets();

let wallets = [];

// open file upload
function openDialog() {
	document.getElementById('fileid').value = null;
	document.getElementById('fileid').click();
}

// Handle event on file selecting
function handleFileSelect(evt) {
	let files = evt.target.files; // get files
	let f = files[0];
	const reader = new FileReader();
	reader.onload = (event) => importWallet(JSON.parse(event.target.result)); // desired file content
	reader.onerror = (error) => reject(error);
	reader.readAsText(f);
}

// Import a wallet through PUT query
async function importWallet(wallet) {
	axios
		.put('/mgmt/wallet', wallet)
		.then((_) => {
			tableInsert(wallet);
			wallets.push(wallet);
		})
		.catch((e) => {
			errorAlert(e.response.data.code);
		});
}

function errorAlert(error) {
	document.getElementsByClassName('alert-danger')[0].style.display = 'block';

	document.getElementsByClassName('alert-danger')[0].innerHTML = error;

	setTimeout(function () {
		document.getElementsByClassName('alert-danger')[0].style.display = 'none';
	}, 5000);
}
