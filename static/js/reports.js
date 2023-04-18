function DismissReport(id) {
    const XHR = new XMLHttpRequest();

    XHR.open("DELETE", `/api/v1/report/${id}`);

    XHR.addEventListener("load", _ => {
        if (XHR.status !== 200) {
            alert("some kind of error has occurred");
        }
        window.location.reload();
    });

    XHR.send();
}