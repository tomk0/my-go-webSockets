var sock = null;
var wsuri = "ws://192.168.1.60:1234";

window.onload = function() {

    console.log("onload");

    sock = new WebSocket(wsuri);

    sock.onopen = function() {
        console.log("connected to " + wsuri);
    }

    sock.onclose = function(e) {
        console.log("connection closed (" + e.code + ")");
    }

    sock.onmessage = function(e) {
        var object = JSON.parse(e.data)
        console.log("message received: " + JSON.stringify(object));
    }
};

function send() {

    var msg = document.getElementById('message').value;
    sock.send(msg);
};