# 😎 Pardalis - Porque Aprender Inglés en Primaria Nunca Fue Tan "Divertido" (O eso esperamos)

## 🚀 Descripción 

**Pardalis** es la plataforma educativa que nadie sabía que necesitaba, pero aquí estamos. Diseñada con sudor, lágrimas y probablemente muchas tazas de café,<> **Pardalis** tiene como objetivo (o al menos eso creemos) hacer que los niños de primaria baja aprendan inglés de una manera más... digamos... interactiva y estimulante. Ya sabes, esas cosas que se dicen cuando intentas que alguien use tu plataforma.

En el corazón de **Pardalis** hay una profunda filosofía educativa. O quizás solo queríamos sonar serios, quién sabe. La plataforma está llena de actividades lúdicas (porque "jugar es aprender" o algo así) y lecciones estructuradas (suena aburrido, pero prometemos que no lo es... tanto). Todo esto para hacer que los niños de primaria vean el inglés con ojos brillantes y llenos de esperanza, en lugar de con pánico absoluto. 😱

## 🙌 Justificación (Sí, tenemos una razón... o varias)

¿Y por qué estamos haciendo esto? Buena pregunta. 🤔 Pues resulta que el equipo **Ponchoides**/**Eqüipito** (sí, somos tan geniales como suena) decidió que la educación primaria es la raíz de todos los problemas educativos del mundo (ok, quizás no todos, pero si de México). Nos dimos cuenta de que el inglés en la primaria era, en nuestra experiencia, inexistente o mal enseñado. Así que decidimos hacer algo al respecto... ¡Bam! Nace **Pardalis**.

Queremos que los niños aprendan inglés desde pequeños para que, cuando lleguen a la secundaria, no piensen que "English es una pesadilla". Y quién sabe, quizás hasta les guste. P.D: El programador principal no sabe inglés, solo poncho que paso el espa de inglés.

## 🎓 El Impacto en la Educación (o eso nos gusta pensar)

Nos dirigimos a la educación primaria porque creemos que es **el momento crucial** para desarrollar habilidades lingüísticas, cognitivas y... ¿qué más? Ah, sí, digitales. 💻 Porque si no aprendes a programar y hablar inglés, ¿estás realmente viviendo en el siglo XXI? Queremos dar a los estudiantes las herramientas para conquistar el futuro y, de paso, que no sufran como nosotros lo hicimos intentando entender esos exámenes de inglés llenos de "fill in the blanks". 

## 📋 Características Principales (Aquí es donde nos sentimos orgullosos)

- **Registro y Login** 🛂: Porque necesitamos saber quién eres antes de dejarte entrar en este paraíso digital. Además, las bases de datos tienen hambre de tus datos. 🍽️
- **Lecciones Interactivas** 📖: Aprender inglés ya no será aburrido (esperamos, pero no prometemos nada). 
- **Actividades Lúdicas** 🎮: Porque sabemos que los niños prefieren jugar antes que leer... obvio. 
- **Autenticación JWT** 🔒: No, no es un combo de comida rápida. Es seguridad de nivel casi militar para que nadie más se haga pasar por ti. 
- **API REST** 🛠️: Porque en el fondo, somos unos buenos programadores y necesitábamos decir que tenemos una API REST 🤓☝️.

## 🛠️ Instalación (Para el valiente que quiera probar)

1. Clona este repositorio como si no hubiera un mañana:

   ```bash
   git clone https://codeberg.org/Pardalis/pardalis-api.git
   ```
   
2. ¡Configura las variables de entorno! ¿Por qué? Porque si no, nada funcionará. Simple, pero cierto.

   ```bash
   export PORT=8080
   export JWT_SECRET="SuperSecretTokenNoTanSecreto"
   ```
   
3. Ahora instala las dependencias, porque aunque sea doloroso, es necesario:

   ```bash
   go mod tidy
   ```
   
4. Levanta el servidor y observa cómo la magia (o el caos) sucede ante tus ojos:

   ```bash
   go run main.go
   ```

## 📚 Endpoints (Aquí es donde la API cobra vida)

- **POST /login**: Porque, obviamente, necesitas autenticación para empezar.
- **POST /register**: Para esos niños que se registran y esperan aprender inglés sin darse cuenta de lo que les espera.
- **GET /users/{userApodo}**: Porque todo el mundo necesita su apodo genial en la plataforma.

## 📦 Tecnologías Usadas (Tranquilo, todo de código abierto. Nada que debas pagar)

- **Golang** 💻: Porque queríamos ser modernos y cool, pero no podíamos usar algo demasiado difícil.
- **SQLite** 🗄️: Bases de datos locales para que puedas decir "Mira mamá, sin servidores".
- **JWT** 🔒: Para asegurarnos de que nadie más use tu cuenta y robe tu progreso en inglés (porque eso sería triste).
- **Gorilla Mux** 🦍: Enrutamiento salvaje para tu experiencia de API.

## 🙃 Contribuciones (Como si alguien fuera a contribuir...)

Si por alguna razón cósmica sientes que debes contribuir, ¡adelante! Haz un fork, abre un pull request y nosotros, con suerte, lo leeremos antes de rechazarlo.

## 😜 Licencia

Este proyecto está bajo la Licencia GPL v3. Así que, si lo modificas y distribuyes, más te vale compartir esos cambios con el mundo 🌍 (o los abogados de Stallman te buscarán 😅). Y si te haces millonario vendiéndolo, al menos asegúrate de compartir también tus ganancias, no solo el código. 💸

---

Gracias por leer este README... si es que alguien lo leyó. Si llegaste hasta aquí, te mereces una galleta virtual. 🍪
Menos tu Jos por rechazar ser integrante de este proyecto. No cierto, pero no tienes galleta virtual. 😡
