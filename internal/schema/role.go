package schema

var Role = Resource{
	Key:  "role",
	Name: "角色",
	Fields: map[string]Field{
		"id": {
			Key:  "id",
			Name: "ID",
		},
		"name": {
			Key:  "name",
			Name: "名称",
		},
		"desc": {
			Key:  "desc",
			Name: "描述",
		},
	},
}
