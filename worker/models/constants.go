package models

const TODO_RESOURCE = "/internal/todos"
const DATE_FORMAT = "2006-01-02"
const EMAIL_BODY_TEMPLATE = `
<html>
	<body>
		<p>Hello!</p>
		<p>Below are the todos due Today</p>
		<p>%s</p>
	</body>
</html>
`
const EMAIL_SUBJECT_TEMPLATE = `[Important] You Have %d Todo/s Due Today`
