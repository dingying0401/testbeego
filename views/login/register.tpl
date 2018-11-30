<html>

<head>
    <meta charset="UTF-8">
    <title>Register</title>
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
    <div id="title">
        <i class="material-icons lock">lock</i> 注册
    </div>

    <form method="post" action="/register">
        <div class="input">
            <div class="input-addon">
                <i class="material-icons">face</i>
            </div>
            <input id="membername" name="membername" placeholder="用户名" type="text" required class="validate" autocomplete="off">
        </div>

        <div class="clearfix"></div>

        <div class="input">
            <div class="input-addon">
                <i class="material-icons">vpn_key</i>
            </div>
            <input id="memberpassword" name="memberpassword" placeholder="密码" type="password" required class="validate" autocomplete="off">
        </div>

    </br>
        <input type="submit" value="注册" />
    </br>
    </form>
</body>

</html>