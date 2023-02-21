const params = new Proxy(new URLSearchParams(window.location.search), {
    get: (searchParams, prop) => searchParams.get(prop),
});

function buildComment(commentData) {
    let comment = document.createElement("div")

    comment.setAttribute("class", "box")
    comment.innerHTML = `
        <div class="Box-header">
            <img class="avatar avatar-5 mr-2" alt="User avatar" src="${commentData.poster_avatar}" />
            ${commentData.poster}
            <span class="branch-name float-right pt-1">${time}</span>
        </div>
        <div class="Box-body">
            ${comment.content}
        </div>
    `;

    return comment
}

async function updateComments() {

    let comments = await fetch(`/api/v1/report/${params.id}/comments`).then((res) => {
        if (res.status !== 200) {
            return "Failed to load comments";
        }
        return res.json();
    });

    let commentContainer = document.getElementById("comment_container");
    commentContainer.replaceChildren();


    for (let comment in comments) {
        let commentElement = buildComment(comment);
        commentContainer.append(commentElement);
    }
}

function postComment() {
    const form = document.getElementById("comment_form");

    const XHR = new XMLHttpRequest();

    // Bind the FormData object and the form element
    const FD = new FormData(form);

    // Define what happens on successful data submission
    XHR.addEventListener("load", (event) => {

    });

    // Define what happens in case of error
    XHR.addEventListener("error", (event) => {
        alert('Oops! Something went wrong.');
    });

    // Set up our request
    XHR.open("POST", `/api/v1/report/${params.id}/post_comment`, true);

    // The data sent is what the user provided in the form
    XHR.send(FD);

    return false
}

window.setInterval(updateComments, 15000);