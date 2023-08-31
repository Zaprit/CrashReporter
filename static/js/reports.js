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

function ReopenReport(id) {
    const XHR = new XMLHttpRequest();

    XHR.open("POST", `/api/v1/report/${id}/reopen`);

    XHR.addEventListener("load", _ => {
        if (XHR.status !== 200) {
            alert("some kind of error has occurred");
        }
        window.location.reload();
    });

    XHR.send();
}

async function BanUser(ip) {
    const duration = window.prompt("Ban Duration, in days");
    if (!duration || !duration.trim() || isNaN(duration)) {
        alert("Invalid duration");
        return
    }
    const reason = window.prompt("Ban Reason:");
    let resp = await fetch("/api/v1/ban", {
        method: "POST",
        body: JSON.stringify({
            ip: ip,
            duration: duration,
            reason: reason,
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    });
    alert(resp.body);
}