var template = '<li>%s (<a href="/latest/%s/%s">%s</a> or <a href="/latest/%s/%s">%s</a>)</li>';

document.addEventListener('DOMContentLoaded', function(event) {
  var xhr = new XMLHttpRequest();
  xhr.open('GET', '/distributions');
  xhr.onload = function(e) {
    var list = document.querySelector("ul#js-downloads");
    var content = '';
    var data = JSON.parse(this.response);
    for (var i = 0; i < data.length; i++) {
      var dist = data[i];

      content += [
        dist.name,
        dist.os,
        dist.arches[0].id,
        dist.arches[0].name,
        dist.os,
        dist.arches[1].id,
        dist.arches[1].name
      ].reduce(
        function (p, c) {
          return p.replace(/%s/, c);
        },
        template
      );
    }
    list.innerHTML = content;
  };
  xhr.send();
});
