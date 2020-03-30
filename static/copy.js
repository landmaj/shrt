const delay = ms => new Promise(res => setTimeout(res, ms));

const animateButton = async (button) => {
    button.classList.add("is-loading");
    await delay(300);
    button.classList.add("is-success");
    button.classList.remove("is-loading");
    button.classList.remove("is-link");
    button.innerHTML = "Copied!";
    await delay(1500);
    button.classList.add("is-link");
    button.classList.remove("is-success");
    button.innerHTML = "Copy";
};


function copyToClipboard() {
    let link = document.getElementById("link");
    link.select();
    link.setSelectionRange(0, 99999); /*For mobile devices*/
    document.execCommand("copy");
    animateButton(document.getElementById("copy"))
}

