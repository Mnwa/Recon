package adapters

type Adapter interface {
	Create(project string, projectType string, data []byte) error
	CreateKey(project string, projectType string, key string, data []byte) error

	Update(project string, projectType string, data []byte) error
	UpdateKey(project string, projectType string, key string, data []byte) error

	Get(project string, projectType string) string
	GetKey(project string, projectType string, key string) string

	Delete(project string, projectType string) error
	DeleteKey(project string, projectType string, key string) error
}
