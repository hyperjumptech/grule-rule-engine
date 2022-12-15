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

function MarkUp(src) {
    let re =  /rule/gi;
    let poses = src.match(re);
    let marks = []
    if (poses) {
        while ((match = re.exec(src)) != null) {
            marks.push(
                {
                    start: match.index,
                    end: match.index+4,
                    text: src.substring(match.index,match.index+4),
                    kind: "rule"
                }
            );
        }
    }
    return marks;
}

function MarkDown(src) {
    return src;
}

function Mark() {
    let markers = MarkUp($("#grleditor").text());
    // do nothing
}

function JsonEditorTab(e) {
    if (e.keyCode === 9) { // tab key
        e.preventDefault();  // this will prevent us from tabbing out of the editor

        // now insert four non-breaking spaces for the tab key
        var editor = document.getElementById("jsoneditor");
        var doc = editor.ownerDocument.defaultView;
        var sel = doc.getSelection();
        var range = sel.getRangeAt(0);

        var tabNode = document.createTextNode("\u00a0\u00a0\u00a0\u00a0");
        range.insertNode(tabNode);

        range.setStartAfter(tabNode);
        range.setEndAfter(tabNode);
        sel.removeAllRanges();
        sel.addRange(range);
    }
}

function GrlEditorTab(e) {
    if (e.keyCode === 9) { // tab key
        e.preventDefault();  // this will prevent us from tabbing out of the editor

        // now insert four non-breaking spaces for the tab key
        var editor = document.getElementById("grleditor");
        var doc = editor.ownerDocument.defaultView;
        var sel = doc.getSelection();
        var range = sel.getRangeAt(0);

        var tabNode = document.createTextNode("\u00a0\u00a0\u00a0\u00a0");
        range.insertNode(tabNode);

        range.setStartAfter(tabNode);
        range.setEndAfter(tabNode);
        sel.removeAllRanges();
        sel.addRange(range);
    }
}

