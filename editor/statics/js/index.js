function ShowGRL() {
    $("#navGRL").addClass("active");
    $("#navContext").removeClass("active");
    $("#navResult").removeClass("active");

    $("#panelGRL").css("display", "block");
    $("#panelContext").css("display", "none");
    $("#panelResult").css("display", "none");
}

function ShowContext() {
    $("#navContext").addClass("active");
    $("#navGRL").removeClass("active");
    $("#navResult").removeClass("active");

    $("#panelContext").css("display", "block");
    $("#panelGRL").css("display", "none");
    $("#panelResult").css("display", "none");
}

function ShowResult() {
    $("#navResult").addClass("active");
    $("#navContext").removeClass("active");
    $("#navGRL").removeClass("active");

    $("#panelResult").css("display", "block");
    $("#panelContext").css("display", "none");
    $("#panelGRL").css("display", "none");
}

function executeRule() {
    let grl = $("#grlText").val();
    let json = $("#jsonText").val()
    let grlB64 = btoa(grl);
    let jsonB64 = btoa(json);

    $.post( "/evaluate", JSON.stringify({"grlText": grlB64, "jsonText": jsonB64})  , function( data, status ) {
        $("#response").val(status + " : " + JSON.stringify(data) );
    }, "json") .fail(function(data) {
        $("#response").val( "Status " + data.status + " : " + data.statusText + ". ResponseText : " + data.responseText);
    });

}