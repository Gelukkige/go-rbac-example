package schema

var User = Resource{
	Key:  "user",
	Name: "用户",
	Fields: map[string]Field{
		"id": {
			Key:  "id",
			Name: "ID",
		},
		"name": {
			Key:  "name",
			Name: "姓名",
		},
		"phone": {
			Key:  "phone",
			Name: "手机号",
		},
		"email": {
			Key:  "email",
			Name: "邮箱",
		},
	},
}
