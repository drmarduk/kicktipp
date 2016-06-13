package main

var htmllogin string = `
<html>
	<body>
		<form action="/login" method="post">
			<input type="text" name="user" />
			<input type="password" name="pass" />
			<input type="submit" value="Login" />
		</form>
	</body>
</html>
`
