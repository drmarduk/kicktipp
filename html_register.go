package main

var htmlregister string = `
<html>
	<body>
		<form action="/register" method="post">
			<input type="text" name="user" />
			<input type="text" name="email" />
			<input type="password" name="pass" />
			<input type="submit" value="Login" />
		</form>
	</body>
</html>
`
