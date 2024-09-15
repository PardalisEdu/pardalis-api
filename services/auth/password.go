// Package auth 🐄 – Aquí te enseñamos a "esconder" contraseñas y a fingir que todo está bajo control.
// ¡Usamos bcrypt porque un hash nunca será lo suficientemente crujiente! 🥐
package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword 🐄 – La función que toma una contraseña y la transforma en una sopa de letras irreconocible,
// garantizando que ni siquiera tú puedas adivinarla. 😅
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Usamos bcrypt para garantizar que tu contraseña esté protegida... ¡incluso de ti mismo! 🔒
	if err != nil {
		return "", err // Si algo falla, te devolvemos un error, porque la vida no siempre es tan dulce. 🍬
	}

	return string(hash), nil // Retorna el hash de la contraseña, que ahora parece más una contraseña Wi-Fi imposible de recordar. 📶
}

// ComparePasswords 🐄 – La función que compara una contraseña "plana" con una "crujiente".
// Si coinciden, ¡bingo! Si no, pues... mejor suerte para la próxima. 🎯
func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain) // Compara el hash con la contraseña sin encriptar, como si fuera un examen sorpresa. 📋
	return err == nil                                           // Retorna true si pasan el examen, false si no. ¡A veces estudiar no es suficiente! 📚
}
