function generateReportId(min, max) {
    return Math.random() * (max - min) + min;
}

//Resize window
function mobile() {
    if ($(window).width() < 700) {
        $('#ReportContainer').removeClass('col-10');
        $('#Sidebar').hide();
    } else {
        $('#ReportContainer').addClass('col-10');
        $('#Sidebar').show();
    }
}

window.onload = function () {
    mobile();
}

window.onresize = function () {
    mobile();
}