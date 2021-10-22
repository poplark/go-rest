package router

import (
	"log"
	"poplark/rest-blog/dbs"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Offset int64 `form:"offset"`
	Limit int64 `form:"limit"`
}
func getUsers(c *gin.Context) {
	var query Query
	if c.ShouldBindQuery(&query) == nil {
		log.Println(query.Limit)
		log.Println(query.Offset)
	}
	count := dbs.Count(false)
	var offset int64
	if query.Offset >= 0 {
		offset = query.Offset
	}
	limit := int64(10)
	if query.Limit > 0 {
		limit = query.Limit
	}
	users := dbs.Find(offset, limit, false)
	c.JSON(200, gin.H{"count": count, "data": users, "offset": offset, "limit": limit})
}

func Handler() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run(":8008")
}
