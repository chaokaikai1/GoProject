package Models

//Student 学生
type Student struct {
	ID     int    `json:"id"` //  "ID"首字母大写是因为要序列化，必须大写，json包里才能访问。  `json:"id"`是让序列化后，"ID"用小写表示
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

//MClass 班级
type MClass struct {
	Title    string    `json:"title"`
	Students []Student `json:"students"`
}
