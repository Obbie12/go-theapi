package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"go-theapi/db"
	"go-theapi/model"
	"go-theapi/repository"

	"github.com/labstack/echo/v4"
)

func C_login_post(c echo.Context) (err error) {
	u := new(model.M_login_input)
	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err := CheckLogin(*u, c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res)
}

func CheckLogin(user model.M_login_input, c echo.Context) (model.ResponseData, error) {
	var obj model.M_login_user
	var res model.ResponseData

	con := db.ConnectPostgresql()

	sqlStatement := "SELECT username, password, full_name FROM users WHERE username = $1"

	err := con.QueryRow(sqlStatement, user.Username).Scan(
		&obj.Username, &obj.Password, &obj.Displayname)
	con.Close()
	if err == sql.ErrNoRows {
		fmt.Println(err)
		res.Status = http.StatusNotFound
		res.Message = "Username not found"
		return res, err
	}

	if err != nil {
		fmt.Println(err)
		res.Status = http.StatusInternalServerError
		res.Message = "Query error"
		return res, err
	}

	if user.Password == obj.Password {
		result, err := repository.GenerateToken(c, user.Username)
		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
			return res, err
		}
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = &model.M_login_result{
			Displayname: obj.Displayname,
			Username:    obj.Username,
			AccessToken: result.AccessToken,
		}
	} else {
		fmt.Println("Incorrect password")
		res.Status = http.StatusNotFound
		res.Message = "Incorrect password"
	}

	return res, nil
}
