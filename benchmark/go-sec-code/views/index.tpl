<!DOCTYPE html>
<html>
<head>
  <title>GO-SEC-CODE</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://cdn.staticfile.org/twitter-bootstrap/5.1.1/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.staticfile.org/twitter-bootstrap/5.1.1/js/bootstrap.bundle.min.js"></script>
</head>
<body>

<div class="container-fluid p-5 bg-primary text-white text-center">
  <h1>GO-SEC-CODE</h1>
  <p>Go-sec-code is a  project for learning Go vulnerability code.</p> 
</div>
  
<div class="container mt-4">
  <div class="row">
    <div class="col-sm-4">
      <h3>CommandInject</h3>
      <h4><a href="/commandInject/vuln?dir=%2F">vuln:unchecked get param</a></h4>
      <h4><a href="/commandInject/vuln/host">vuln:unchecked host</a></h4>
      <h4><a href="/commandInject/vuln/git?repoUrl=--upload-pack=$(open /)">vuln:unchecked git param</a></h4>
      <h4><a href="/commandInject/safe?dir=.%3Bwhoami">safe:filter</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>Cors</h3>
      <h4><a href="/cors/vuln/reflect">vuln:reflect</a></h4>
      <h4><a href="/cors/vuln/any-origin-with-credential">vuln:any-origin-with-credential</a></h4>
      <h4><a href="/cors/safe">safe:whilelist</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>CRLFInjection</h3>        
      <h4><a href=/crlfInjection/safe?header>safe:unexploitable</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>FileUpload</h3>
      <h4><a href=/fileUpload/vuln>vuln:unchecked post param</a></h4>
      <h4><a href=/fileUpload/safe>safe:filter</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>JSONP</h3>        
      <h4><a href="/jsonp/vuln/noCheck?callback=jsonp">vuln:unchecked host</a></h4>
      <h4><a href="/jsonp/vuln/emptyReferer?callback=jsonp">vuln:empty referer bypass</a></h4>
      <h4><a href="/jsonp/safe?callback=jsonp">safe:whitelist</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>PathTraversal</h3>        
      <h4><a href="/pathTraversal/vuln?file=../../../../../../../etc/passwd">vuln:unchecked get param</a></h4>
      <h4><a href="/pathTraversal/vuln?file=../../../../../../../etc/passwd">vuln:use Clean() improperly</a></h4>
      <h4><a href="/pathTraversal/safe/filter?file=../../../../../../../etc/passwd">safe:filter</a></h4>
      <h4><a href="/pathTraversal/safe/check?file=../../../../../../../etc/passwd">safe:check</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>SQLInjection</h3>
      <h4><a href="/sqlInjection/native/vuln/integer?id=1">vuln:integer</a></h4>
      <h4><a href="/sqlInjection/native/vuln/string?username=admin">vuln:string</a></h4>
      <h4><a href="/sqlInjection/orm/vuln/xorm?field=username&username=admin">vuln:use orm improperly</a></h4>
      <h4><a href="/sqlInjection/generator/vuln/squirrel?order=id&username=admin">vuln:use generator improperly</a></h4>
      <h4><a href="/sqlInjection/native/safe/integer?id=1">safe:integer</a></h4>
      <h4><a href="/sqlInjection/native/safe/string?username=admin">safe:string</a></h4>
      <h4><a href="/sqlInjection/orm/safe/beego?field=username&username=admin">safe:orm</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>SSRF</h3>
      <h4><a href="/ssrf/vuln?url">vuln:unchecked get param</a></h4>
      <h4><a href="/ssrf/vuln/obfuscation?url">vuln:obfuscation bypass</a></h4>
      <h4><a href="/ssrf/vuln/302?url">vuln:302 bypass</a></h4>
      <h4><a href="/ssrf/safe/whitelists?url">safe:whitelist</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>SSTI</h3>
      <h4><a href="/ssti/vuln?template=&#123;&#123;env &#34;PATH&#34;&#125;&#125;">vuln:concat</a></h4>
      <h4><a href="/ssti/safe?template">safe:no concat</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>XSS</h3>
      <h4><a href="/xss/vuln?xss">vuln:reflect</a></h4>
      <h4><a href="/xss/vuln/store?xss">vuln:store</a></h4>
      <h4><a href="/xss/vuln/svg">vuln:svg</a></h4>
      <h4><a href="/xss/vuln/pdf">vuln:pdf</a></h4>
      <h4><a href="/xss/safe?xss">safe:filter</a></h4>
      <h4><a href="/xss/safe/svg">safe:CSP</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>XXE</h3>
      <h4><a href="/xxe/vuln">vuln:libxml2</a></h4>
      <h4><a href="/xxe/safe">safe:unexploitable</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>ZipSlip</h3>
      <h4><a href="/zipslip/vuln">vuln</a></h4>
    </div>
  </div>
</div>

</body>
</html>