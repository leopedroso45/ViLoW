// Create a request variable and assign a new XMLHttpRequest object to it.
var request = new XMLHttpRequest();

// Open a new connection, using the GET request on the URL endpoint
request.open("GET", "http://localhost:8000/videos", true);

request.onload = function() {
  // Begin accessing JSON data here
  var data = JSON.parse(this.response);

  if (request.status >= 200 && request.status < 400) {
    data.forEach(video => {
      console.log(video);
      criaComponent(video);
    });
  } else {
    console.log("error");
  }
};

// Send request
request.send();

function criaComponent(video) {
  var div1 = document.createElement("div");
  var center = document.createElement("center");
  var div2 = document.createElement("div");
  var videohtml = document.createElement("video");
  var div3 = document.createElement("div");
  var p = document.createElement("p");
  var div4 = document.createElement("div");
  var div5 = document.createElement("div");
  var btn1 = document.createElement("BUTTON");
  var btn2 = document.createElement("BUTTON");
  var sourceVi = document.createElement("source");

  div5.appendChild(btn1);
  div5.appendChild(btn2);
  div4.appendChild(div5);
  div3.appendChild(p);
  div3.appendChild(div4);
  videohtml.appendChild(sourceVi);
  div2.appendChild(videohtml);
  div2.appendChild(div3);
  center.appendChild(div2);
  div1.appendChild(center);

  var att = document.createAttribute("class");
  att.value = "col-md-4";
  div1.setAttributeNode(att);
  var att2 = document.createAttribute("class");
  att2.value = "card mb-4 shadow-sm";
  div2.setAttributeNode(att2);
  att = document.createAttribute("width");
  att.value = "320";
  videohtml.setAttributeNode(att);
  att = document.createAttribute("height");
  att.value = "240";
  videohtml.setAttributeNode(att);
  att = document.createAttribute("controls");
  videohtml.setAttributeNode(att);
  att = document.createAttribute("src");
  att.value = "app:8000/" + video.PathVideo;
  sourceVi.setAttributeNode(att);
  var textName = document.createTextNode(video.DescVideo);
  sourceVi.appendChild(textName);
  var att3 = document.createAttribute("class");
  att3.value = "card-body";
  div3.setAttributeNode(att3);
  var att4 = document.createAttribute("class");
  att4.value = "card-text";
  p.setAttributeNode(att4);
  var textDesc = document.createTextNode(video.NameVideo);
  p.appendChild(textDesc);
  var att = document.createAttribute("class");
  att.value = "btn-group";
  div5.setAttributeNode(att);

  var att = document.createAttribute("type");
  att.value = "button";
  btn1.setAttributeNode(att);
  var att = document.createAttribute("type");
  att.value = "button";
  btn2.setAttributeNode(att);

  var att = document.createAttribute("class");
  att.value = "btn btn-sm btn-outline-secondary";
  btn1.setAttributeNode(att);
  var att = document.createAttribute("class");
  att.value = "btn btn-sm btn-outline-secondary";
  btn2.setAttributeNode(att);

  document.getElementById("demo").appendChild(div1);
}
