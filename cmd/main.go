package cmd

type Player struct {
	Name string `json:"name"` binding:"required"`
	Age int    `json:"age"` binding:"required"`
	CreationTime time.Time `json:"-"`
}

func CreatePlayer(c *gin.Context) {
	// Create a new player


}
