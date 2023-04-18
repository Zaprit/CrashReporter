function postNotice() {
    const form = document.getElementById("notice_form");

    const XHR = new XMLHttpRequest();

    // Bind the FormData object and the form element
    const FD = new FormData(form);

    // Define what happens on successful data submission
    XHR.addEventListener("load", async (_) => {
        await updateNotices();
    });

    // Define what happens in case of error
    XHR.addEventListener("error", (_) => {
        alert('Failed to post comment.');
    });

    // Set up our request
    XHR.open("POST", `/api/v1/notice`, true);

    // The data sent is what the user provided in the form
    XHR.send(FD);

    return false;
}

function deleteNotice(id) {
    const XHR = new XMLHttpRequest();

    XHR.open("DELETE", `/api/v1/notice/${id}`);

    XHR.addEventListener("load", ev => {
        if (XHR.status !== 200) {
            return;
        }
        const noticeID = `notice_${id}`;
        const elem = document.getElementById(noticeID);
        elem.remove();
    });

    XHR.send();
}

async function updateNotices() {

}