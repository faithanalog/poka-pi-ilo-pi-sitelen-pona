
let page = require('webpage').create();
let webserver = require('webserver').create();
page.open("http://sitelen-pona.herokuapp.com/", function(status) {
  // TODO what if it doesn't load properly?
  
  // The only reason this mutex works is because javascript is single-threaded.
  let renderLock = false;
  function renderSitelenPona(tokipona, callback) {
    // If a render is already happening, defer to next event loop cycle
    if (renderLock) {
      setTimeout(function() {
        renderSitelenPona(tokipona, callback);
      });
      return;
    }
    renderLock = true;
    page.evaluate(function(toki) {
      let inputField = document.querySelector('textarea');
      inputField.value = toki;
      let evt = new Event('input', { bubbles: true });
      evt.simulated = true;
      inputField.dispatchEvent(evt);
    }, tokipona);
    window.setTimeout(function() {
      let bb = page.evaluate(function() {
        return document.querySelector('.tokipona').getBoundingClientRect();
      });
      page.clipRect = {
        top: bb.top,
        left: bb.left,
        width: bb.width,
        height: bb.height
      };
      let img = page.renderBase64('PNG');
      setTimeout(function() {
        renderLock = false;
      });
      callback(img);
    });
  }


  window.setTimeout(function() {
    webserver.listen('0.0.0.0:3002', function (req, res) {
      let toki = req.post;
      renderSitelenPona(toki, function(imgb64) {
        res.statusCode = 200;
        res.write(imgb64);
        res.close();
      });
    });
  }, 100);
});

