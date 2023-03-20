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
    <form class="container" method="post" enctype="multipart/form-data">
        <div class="mb-3">
            <label for="formFile" class="form-label">Choose file to upload</label>
            <input class="form-control" type="file" id="formFile" name="file">
        </div>
        <input type="hidden" name=userid value="233">
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
    <div class="container mt-5">
        <h4>{{.savePath}}</h4>
    </div>
  </body>