package student

const (
	getAllStudents    = "select id, name, major, created_at from student"
	getStudentByID    = "select name, major, created_at from student where id = ?"
	insertStudent     = "insert into student (name, major, created_at) values (?, ?, ?)"
	deleteStudentByID = "delete from student where id = ?"
)
