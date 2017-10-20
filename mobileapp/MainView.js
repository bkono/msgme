var Observable = require("FuseJS/Observable");
var MsgMeSDK = require("MsgMeSDK");

var count = Observable(0);
var messages = Observable([]);

function test() {
  var result = MsgMeSDK.Test();
  console.dir(result)
  count.value += 1;
}

function send() {
  console.log("starting to send");
  MsgMeSDK.Send("some message");
  console.log("sent");
}

function listen() {
  console.log("starting to listen");
  MsgMeSDK.Listen();
  console.log("... listening");
}

MsgMeSDK.on("messageReceived", function (message) {
  console.log("Message received " + message);
  messages.add({ content: message });
});

module.exports = {
  count: count,
  test: test,
  send: send,
  listen: listen,
  messages: messages
};