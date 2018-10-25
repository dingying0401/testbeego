<html>

<head>
    <meta charset="UTF-8">
    <title>Login</title>
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
        <i class="material-icons lock">lock</i> 登录
    </div>

    <form>
        <div class="input">
            <div class="input-addon">
                <i class="material-icons">face</i>
            </div>
            <input id="username" placeholder="用户名" type="text" required class="validate" autocomplete="off">
        </div>

        <div class="clearfix"></div>

        <div class="input">
            <div class="input-addon">
                <i class="material-icons">vpn_key</i>
            </div>
            <input id="password" placeholder="密码" type="password" required class="validate" autocomplete="off">
        </div>

        <div class="remember-me">
            <input type="checkbox">
            <span style="color: #DDD">记住密码</span>
        </div>


        <input type="submit" value="登录" />
    </form>

    <div class="forgot-password">
        <a href="#">忘记密码</a>
    </div>
    <div class="privacy">
        <a href="#">隐私政策</a>
    </div>

    <div class="register">
        还没有账号？
        <a href="#"><button id="register-link">点击这里注册</button></a>
    </div>
</div>
</body>

</html>