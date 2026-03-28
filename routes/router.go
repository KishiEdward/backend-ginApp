package routes // [cite: 1364]

import ( // [cite: 1365]
	"github.com/gin-gonic/gin" // [cite: 1367]
	"github.com/yourusername/backend_ginApp/handlers" // [cite: 1368]
	"github.com/yourusername/backend_ginApp/middleware" // [cite: 1368]
) // [cite: 1366]

func SetupRouter() *gin.Engine { // [cite: 1369]
	// gin.Default() sudah include Logger & Recovery middleware // [cite: 1370]
	r := gin.Default() // [cite: 1370]