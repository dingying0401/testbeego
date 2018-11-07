<html>

<head>
  <meta charset="UTF-8">
  <title>贩卖小哔仔主页</title>
  <meta name="description" content="Login - Register Template">
  <meta name="author" content="Lorenzo Angelino aka MrLolok">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="static/css/main.css">
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <style>
    body {
      background-color: #303641;
    }
  </style>
</head>

<body>
<div id="container-login">
</br>
  <h3>小哔仔多少钱一斤</h3>
  <form action="/home" method="post">
    {{.Form | renderform}}
    </br>
    <input type="submit" value="预算评估">
      <form enctype="multipart/form-data" method="post">
        <input type="file" name="uploadname" />
        <input type="submit" value="上传文件"/>
      </form>
    </form>
   <form action="/detail" method="get">
   <input type="submit" value="宝贝详情">
   </form>
</div>
<script src="/static/js/reload.min.js"></script>
</body>
</html>
