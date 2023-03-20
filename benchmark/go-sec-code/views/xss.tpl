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
        <a href="/" style="color: white;">back to home</a>
    </div>
    <div class="container mt-5">
            <h3>post me a xss to store</h3>
    </div>
    <div class="container mt-5">
        <form class="container" method="post">
            <input name=xss value="233" style="width: 250px;height: 38px;">
            <button type="submit" class="btn btn-primary">Submit</button>
        </form>
    </div>
    <div class="container mt-5">
        <p>xss content:{{.xss}}</p>
    </div>
  </body>