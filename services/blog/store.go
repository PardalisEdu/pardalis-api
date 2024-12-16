package blog

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"gitlab.com/pardalis/pardalis-api/types"
)

type Store struct {
	db *sql.DB
}

func NewBlogStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateBlog(blog types.Blog) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Printf("Error al iniciar la transacción: %v", err)
		return err
	}

	// Insertar el blog
	query := `
        INSERT INTO blogs (
            id, titulo, slug, contenido, extracto, 
            imagen_portada, fecha_publicacion, estado,
            categoria, tiempo_lectura, autor_apodo,
            meta_descripcion, meta_keywords
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	log.Printf("Ejecutando query: %s", query)
	log.Printf("Con valores: %+v", blog)

	_, err = tx.Exec(query,
		blog.ID, blog.Titulo, blog.Slug, blog.Contenido,
		blog.Extracto, blog.ImagenPortada, blog.FechaPublicacion,
		blog.Estado, blog.Categoria, blog.TiempoLectura,
		blog.AutorApodo, blog.MetaDescripcion, blog.MetaKeywords,
	)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		log.Printf("Error al insertar el blog: %v", err)
		return err
	}

	// Insertar tags
	if len(blog.Tags) > 0 {
		for _, tag := range blog.Tags {
			err = s.addTagTx(tx, blog.ID, tag)
			if err != nil {
				err := tx.Rollback()
				if err != nil {
					return err
				}
				log.Printf("Error al insertar tag %s: %v", tag, err)
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error al hacer commit de la transacción: %v", err)
		return err
	}

	log.Printf("Blog creado exitosamente en la base de datos")
	return nil
}

func (s *Store) UpdateBlog(blog types.Blog) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Actualizar el blog
	query := `
        UPDATE blogs 
        SET titulo = ?, slug = ?, contenido = ?, extracto = ?,
            imagen_portada = ?, estado = ?, categoria = ?,
            tiempo_lectura = ?, meta_descripcion = ?, meta_keywords = ?
        WHERE id = ?
    `

	result, err := tx.Exec(query,
		blog.Titulo, blog.Slug, blog.Contenido,
		blog.Extracto, blog.ImagenPortada, blog.Estado,
		blog.Categoria, blog.TiempoLectura,
		blog.MetaDescripcion, blog.MetaKeywords,
		blog.ID,
	)

	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	if rowsAffected == 0 {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return fmt.Errorf("blog not found")
	}

	// Eliminar tags existentes
	_, err = tx.Exec("DELETE FROM blog_posts_tags WHERE blog_id = ?", blog.ID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	// Insertar nuevos tags
	for _, tag := range blog.Tags {
		err = s.addTagTx(tx, blog.ID, tag)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) DeleteBlog(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Eliminar tags primero debido a la restricción de clave foránea
	_, err = tx.Exec("DELETE FROM blog_posts_tags WHERE blog_id = ?", id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	// Eliminar el blog
	result, err := tx.Exec("DELETE FROM blogs WHERE id = ?", id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	if rowsAffected == 0 {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return fmt.Errorf("blog not found")
	}

	return tx.Commit()
}

func (s *Store) AddBlogTag(blogID string, tag string) error {
	return s.addTagTx(s.db, blogID, tag)
}

func (s *Store) RemoveBlogTag(blogID string, tag string) error {
	query := `
        DELETE pt FROM blog_posts_tags pt
        JOIN blog_tags t ON pt.tag_id = t.id
        WHERE pt.blog_id = ? AND t.nombre = ?
    `

	result, err := s.db.Exec(query, blogID, tag)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tag not found")
	}

	return nil
}

// Función auxiliar para añadir tags dentro de una transacción
func (s *Store) addTagTx(tx interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}, blogID string, tag string) error {
	// Primero intentamos insertar el tag si no existe
	_, err := tx.Exec("INSERT IGNORE INTO blog_tags (id, nombre) VALUES (UUID(), ?)", tag)
	if err != nil {
		return err
	}

	// Luego vinculamos el tag con el blog
	query := `
        INSERT INTO blog_posts_tags (blog_id, tag_id)
        SELECT ?, id FROM blog_tags WHERE nombre = ?
    `

	_, err = tx.Exec(query, blogID, tag)
	return err
}

func (s *Store) GetBlogBySlug(slug string) (*types.Blog, error) {
	query := `
        SELECT 
            b.id, b.titulo, b.slug, b.contenido, b.extracto, 
            b.imagen_portada, b.fecha_publicacion, b.estado,
            b.categoria, b.tiempo_lectura, b.autor_apodo,
            b.meta_descripcion, b.meta_keywords
        FROM blogs b
        WHERE b.slug = ? AND b.estado = 'publicado'
    `

	blog := &types.Blog{}
	err := s.db.QueryRow(query, slug).Scan(
		&blog.ID, &blog.Titulo, &blog.Slug, &blog.Contenido,
		&blog.Extracto, &blog.ImagenPortada, &blog.FechaPublicacion,
		&blog.Estado, &blog.Categoria, &blog.TiempoLectura,
		&blog.AutorApodo, &blog.MetaDescripcion, &blog.MetaKeywords,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("blog not found")
		}
		return nil, err
	}

	// Obtener tags
	tags, err := s.GetBlogTags(blog.ID)
	if err != nil {
		return nil, err
	}
	blog.Tags = tags

	return blog, nil
}

func (s *Store) GetBlogs(page, limit int, categoria string) ([]types.Blog, error) {
	offset := (page - 1) * limit

	var query string
	var args []interface{}

	if categoria != "" && categoria != "Todos" {
		query = `
            SELECT 
                b.id, b.titulo, b.slug, b.extracto, 
                b.imagen_portada, b.fecha_publicacion,
                b.categoria, b.tiempo_lectura, b.autor_apodo
            FROM blogs b
            WHERE b.estado = 'publicado' AND b.categoria = ?
            ORDER BY b.fecha_publicacion DESC
            LIMIT ? OFFSET ?
        `
		args = []interface{}{categoria, limit, offset}
	} else {
		query = `
            SELECT 
                b.id, b.titulo, b.slug, b.extracto, 
                b.imagen_portada, b.fecha_publicacion,
                b.categoria, b.tiempo_lectura, b.autor_apodo
            FROM blogs b
            WHERE b.estado = 'publicado'
            ORDER BY b.fecha_publicacion DESC
            LIMIT ? OFFSET ?
        `
		args = []interface{}{limit, offset}
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var blogs []types.Blog
	for rows.Next() {
		var blog types.Blog
		err := rows.Scan(
			&blog.ID, &blog.Titulo, &blog.Slug, &blog.Extracto,
			&blog.ImagenPortada, &blog.FechaPublicacion,
			&blog.Categoria, &blog.TiempoLectura, &blog.AutorApodo,
		)
		if err != nil {
			return nil, err
		}

		// Obtener tags para cada blog
		tags, err := s.GetBlogTags(blog.ID)
		if err != nil {
			return nil, err
		}
		blog.Tags = tags

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (s *Store) GetBlogTags(blogID string) ([]string, error) {
	query := `
        SELECT t.nombre
        FROM blog_tags t
        JOIN blog_posts_tags pt ON t.id = pt.tag_id
        WHERE pt.blog_id = ?
    `

	rows, err := s.db.Query(query, blogID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (s *Store) GetBlogByID(id string) (*types.Blog, error) {
	query := `
        SELECT 
            b.id, b.titulo, b.slug, b.contenido, b.extracto, 
            b.imagen_portada, b.fecha_publicacion, b.estado,
            b.categoria, b.tiempo_lectura, b.autor_apodo,
            b.meta_descripcion, b.meta_keywords
        FROM blogs b
        WHERE b.id = ?
    `

	blog := &types.Blog{}
	err := s.db.QueryRow(query, id).Scan(
		&blog.ID, &blog.Titulo, &blog.Slug, &blog.Contenido,
		&blog.Extracto, &blog.ImagenPortada, &blog.FechaPublicacion,
		&blog.Estado, &blog.Categoria, &blog.TiempoLectura,
		&blog.AutorApodo, &blog.MetaDescripcion, &blog.MetaKeywords,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("blog not found")
		}
		return nil, err
	}

	tags, err := s.GetBlogTags(blog.ID)
	if err != nil {
		return nil, err
	}
	blog.Tags = tags

	return blog, nil
}
