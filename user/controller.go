package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vignesh.com/gocrudrest/common/log"
	"vignesh.com/gocrudrest/datasource"
)

func init() {
	log.Debug(context.Background(), "msg controller init")
}

func Routes(app *gin.Engine) {
	route := app.Group("/users")
	route.GET("/", getAllUsers)
	route.GET("/:id", getUser)
	route.POST("/", createUser)
	route.PUT("/:id", updateUser)
	route.DELETE("/:id", deleteUser)
}

func getAllUsers(ctx *gin.Context) {
	rows, err := datasource.GetConn().Queryx("select * from users")
	if err != nil {
		log.Error(ctx, "error fetching data : "+err.Error())
	}
	defer rows.Close()
	var response []User
	for rows.Next() {
		var u User
		if err := rows.StructScan(&u); err != nil {
			log.Error(ctx, "error scanning rows : "+err.Error())
		}
		log.Debug(ctx, u)
		response = append(response, u)
	}
	log.Infof(ctx, "Users count : %d", len(response))
	ctx.JSON(200, response)
}

func createUser(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	log.Infof(ctx, "input : %+v", u)
	stmt, err := datasource.GetConn().PrepareNamed("insert into users (name,gender) values (:name,:gender) Returning id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	err = stmt.Get(&u.Id, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	log.Infof(ctx, "userId: %d :: input : %+v", u.Id, u)
	ctx.JSON(http.StatusCreated, u)
}

func updateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var u User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	u.Id = id
	res, err := datasource.GetConn().NamedExec("update users set name=:name,gender=:gender where id=:id", &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	n, _ := res.RowsAffected()
	log.Debug(ctx, n)
	ctx.JSON(http.StatusOK, &u)
}

func getUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var u = User{}
	err = datasource.GetConn().Get(&u, "select * from users where id=$1", &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	log.Debugf(ctx, "user : %+v", u)
	ctx.JSON(http.StatusOK, &u)
}

func deleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := datasource.GetConn().Exec("delete from users where id=$1", &id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if n < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("%d does not exists", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": fmt.Sprintf("%d delete successfully", id),
	})
}
