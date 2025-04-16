package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joaogabrielvianna/repository"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("minha_chave_secreta")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Usando o método GetUserByEmail para buscar o usuário pelo email
	user, err := repository.GetUserByEmail(creds.Email)
	if err != nil {
		// Logando o erro exato para verificar o que deu errado
		logger.ErrorLog(fmt.Sprintf("Erro ao buscar usuário: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Logando a senha fornecida pelo usuário
	logger.InfoLog(fmt.Sprintf("Senha fornecida: %s", creds.Password))

	// Logando a senha criptografada armazenada no banco
	logger.InfoLog(fmt.Sprintf("Senha criptografada armazenada: %s", user.Password))

	// Comparando a senha fornecida com a senha criptografada armazenada
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Erro ao comparar as senhas: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerando o token JWT
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
