package adapters

type Adapter interface {
	Create(project string, projectType string, data []byte) error
	CreateKey(project string, projectType string, key string, data []byte) error

	Update(project string, projectType string, data []byte) error
	UpdateKey(project string, projectType string, key string, data []byte) error

	Get(project string, projectType string) ([]byte, error)
	GetKey(project string, projectType string, key string) ([]byte, error)

	Delete(project string, projectType string) error
	DeleteKey(project string, projectType string, key string) error
}
