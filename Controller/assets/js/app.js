function PrintFunc() {
    if (window.innerWidth < 960) {
        document.getElementById("width-window").value = window.innerWidth - 120;
    } else {
        document.getElementById("width-window").value = window.innerWidth - 460;
    }
    HTMLFormElement.submit()
}
function ReverseFunc(url) {
    console.log(url);
    let Window;
    if (window.innerWidth < 960) {
        Window = window.innerWidth - 80;
    } else {
        Window = window.innerWidth - 410;
    }
    window.location.href = url + "&w=" + Window;
}