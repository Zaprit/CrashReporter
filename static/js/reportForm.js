function sendData() {

    const form = document.getElementById("ReportForm");

    const XHR = new XMLHttpRequest();

    // Bind the FormData object and the form element
    const FD = new FormData(form);

    // Define what happens on successful data submission
    XHR.addEventListener("load", (event) => {
        let container = document.getElementById('ReportContainer');

        let toastFlag = "";

        let icon = `
            <span class="Toast-icon">
                <!-- <%= octicon "info" %>-->
                <svg width="14" height="16" viewBox="0 0 14 16" class="octicon octicon-info" aria-hidden="true">
                    <path
                        fill-rule="evenodd"
                        d="M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"
                    />
                </svg>
            </span>
            `;

        switch (XHR.status) {
            case 201:
                toastFlag = "Toast--success";
                icon = `
                <span class="Toast-icon">
                  <!-- <%= octicon "check" %> -->
                  <svg width="12" height="16" viewBox="0 0 12 16" class="octicon octicon-check" aria-hidden="true">
                    <path fill-rule="evenodd" d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5L12 5z" />
                  </svg>
                </span>
                `;
                form.reset();
                break
            case 400:
                toastFlag = "Toast--error";
                icon = `
                <span class="Toast-icon">
                  <!-- <%= octicon "stop" %> -->
                  <svg width="14" height="16" viewBox="0 0 14 16" class="octicon octicon-stop" aria-hidden="true">
                    <path
                      fill-rule="evenodd"
                      d="M10 1H4L0 5v6l4 4h6l4-4V5l-4-4zm3 9.5L9.5 14h-5L1 10.5v-5L4.5 2h5L13 5.5v5zM6 4h2v5H6V4zm0 6h2v2H6v-2z"
                    />
                  </svg>
                </span>
                `;
                break


        }

        let toast = document.getElementById("statusToast");
        if (toast != null) {
            toast.remove();
        }

        let toastEl = document.createElement("div");
        toastEl.setAttribute("id", "statusToast");
        toastEl.setAttribute("class", `Toast ${toastFlag}`);
        toastEl.innerHTML = `
            ${icon}
            <span class="Toast-content">${XHR.response}</span>
        `;

        container.prepend(toastEl);
    });

    // Define what happens in case of error
    XHR.addEventListener("error", (event) => {
        alert('Oops! Something went wrong.');
    });

    // Set up our request
    XHR.open("POST", "/api/v1/report", true);

    // The data sent is what the user provided in the form
    XHR.send(FD);

    return false
}

//window.addEventListener("load", () => {
//
//
//    // Get the form element
//
//
//    // Add 'submit' event handler
//    form.addEventListener("submit", (event) => {
//        event.preventDefault();
//
//        sendData();
//    }, false);
//});