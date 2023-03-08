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
    let grl = $("#grleditor").text();
    let grlB64 = btoa(grl);
    let size = $('div#panelContext').find('div.jsoneditor').length;
    let jsonBlob = new Array(size);
    $('div#panelContext').find('div.jsoneditor').each(function(index, editor) {
        let json = $(editor).text();
        let jsonB64 = btoa(json);
        jsonBlob[index] = jsonB64;
    })
    $.post( "/evaluate", JSON.stringify({"grlText": grlB64, "jsonInput": jsonBlob})  , function( data, status ) {
        $("#response").text(status + " : " + JSON.stringify(data) );
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

function AddData(e) {
    var div = document.createElement('div');
    var divBody = document.createElement('div');
    divBody.className = "card-body";
    var editor = document.createElement('div');
    editor.className = "jsoneditor";
    editor.setAttribute("contenteditable", true);
    editor.setAttribute("style", style="white-space: pre-wrap; font-family: monospace; border: solid; padding: 15px 15px 15px 15px;");
    editor.setAttribute("role","textbox");
    divBody.appendChild(editor);

    var divFooter = document.createElement('div');
    divFooter.className = "card-footer";
    var button = document.createElement('button');
    button.setAttribute("type","button");
    button.className="btn btn-danger delete";
    button.innerHTML = "Delete Data";
    divFooter.appendChild(button);

    div.appendChild(divBody);
    div.appendChild(divFooter);
    $('div#panelContext').find("div.card-header").after(div);
}

function JsonEditorTab(editor) {
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

$(document).ready(function() {
    $("div#panelContext").on('click','button.delete', function(event) {
        event.preventDefault();
        $(this).parent().parent().remove();
    })

    $("div#panelContext").on('keydown', 'div.jsoneditor', function(event) {
        if (event.keyCode === 9) {
            event.preventDefault();
            JsonEditorTab(this);
        }
    })
})

