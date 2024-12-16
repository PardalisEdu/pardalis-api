// Package types 🐄 – Porque obviamente necesitas un paquete completo solo para definir un par de estructuras y una interfaz.
// Aquí, definimos tipos que probablemente complicarán tu vida más de lo necesario. ¡Disfruta! 🥳
package types

// UserStore 🐄 – La interfaz que promete gestionar a tus usuarios con métodos que
// probablemente no implementaste correctamente. Pero oye, la intención es lo que cuenta. 🎯
type UserStore interface {
	GetUserByApodo(apodo string) (*User, error)   // GetUserByApodo 🐄 – Encuentra al usuario por su apodo... suponiendo que el apodo sea lo suficientemente único y memorable como para ser útil. 🤔
	GetUserByCorreo(correo string) (*User, error) // GetUserByCorreo 🐄 – Encuentra al usuario por su correo electrónico, porque la gente ama recordar múltiples credenciales. 🔍
	CreateUser(User) error                        // CreateUser 🐄 – Crea un usuario, o al menos lo intenta, hasta que las validaciones fallan y todo explota. 💣
}

type BlogStore interface {
	GetBlogBySlug(slug string) (*Blog, error)
	GetBlogs(page, limit int, categoria string) ([]Blog, error)
	CreateBlog(blog Blog) error
	UpdateBlog(blog Blog) error
	DeleteBlog(id string) error
	GetBlogTags(blogID string) ([]string, error)
	AddBlogTag(blogID string, tag string) error
	RemoveBlogTag(blogID string, tag string) error
	GetBlogByID(id string) (*Blog, error)
}
