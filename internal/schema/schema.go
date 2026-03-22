package schema

var registry = make(map[string]Resource)

type Field struct {
	Key  string
	Name string
}

type Resource struct {
	Key    string
	Name   string
	Fields map[string]Field
}

func Init() {
	RegisterResource(User)
	RegisterResource(Role)
	RegisterResource(Permission)
}

func RegisterResource(resource Resource) {
	registry[resource.Key] = resource
}

func GetResource(key string) (Resource, bool) {
	resource, exists := registry[key]
	return resource, exists
}

func GetAllResources() []Resource {
	resources := make([]Resource, 0, len(registry))
	for _, resource := range registry {
		resources = append(resources, resource)
	}
	return resources
}
