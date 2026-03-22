package schema

var Permission = Resource{
	Key:  "permission",
	Name: "权限",
	Fields: map[string]Field{
		"id": {
			Key:  "id",
			Name: "ID",
		},
		"page": {
			Key:  "page",
			Name: "页面",
		},
		"action": {
			Key:  "action",
			Name: "操作",
		},
		"columns": {
			Key:  "columns",
			Name: "列名",
		},
	},
}
