<html>
 <head> <title></title> </head> <body>
<form action="/login" method="post" enctype="multipart/form-data" > 
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
用户名:<input type="text" name="username"> 
密码:<input type="password" name="password">

<input type="file" name="uploadfile" />
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="登陆">
</form> </body>
</html>