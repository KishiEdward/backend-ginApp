package handlers 

import ( 
	"net/http" 
	"time" 

	"github.com/gin-gonic/gin" 
	"github.com/KishiEdward/backend_ginApp/services" 
) 

type AuthHandler struct { 
	authService *services.AuthService 
} 

func NewAuthHandler() *AuthHandler { 
	return &AuthHandler{authService: services.NewAuthService()} 
} 