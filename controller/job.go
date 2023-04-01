package controller

import (
	"encoding/json"
	"fmt"
	"go-theapi/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type R struct {
	Respo []Respo
}

type Respo struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	Created_at   string `json:"created_at"`
	Company      string `json:"company"`
	Company_url  string `json:"company_url"`
	Location     string `json:"location"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	How_to_apply string `json:"how_to_apply"`
	Company_logo string `json:"company_logo"`
}

func C_Get_Jobs(c echo.Context) error {
	result, err := fetchJobs(c)
	if err != nil {
		return c.JSON(result.Status, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func fetchJobs(c echo.Context) (model.ResponseData, error) {
	var res model.ResponseData
	var page string
	var err error
	var page_param, desc, loc, ft string
	var full_time bool
	questionmark := false

	page = c.QueryParam("page")
	if len(page) > 0 {
		if questionmark {
			page_param = "&page=" + page
		} else {
			page_param = "?page=" + page
			questionmark = true
		}
	}

	description := c.QueryParam("description")
	if len(description) > 0 {
		if questionmark {
			desc = "&description=" + description
		} else {
			desc = "?description=" + description
			questionmark = true
		}
	}

	location := c.QueryParam("location")
	if len(location) > 0 {
		if questionmark {
			loc = "&location=" + location
		} else {
			loc = "?location=" + location
			questionmark = true
		}
	}

	full_time, err = strconv.ParseBool(c.QueryParam("full_time"))
	if err == nil {
		if questionmark {
			ft = "&full_time=" + strconv.FormatBool(full_time)
		} else {
			ft = "?full_time=" + strconv.FormatBool(full_time)
			questionmark = true
		}
	}

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"
	endpoint := url + page_param + desc + loc + ft

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Respo
	json.Unmarshal(responseData, &responseObject)

	if responseObject == nil {
		res.Status = http.StatusNotFound
		res.Message = "Record Not Found"
		res.Data = responseObject
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = responseObject

	return res, nil
}

func C_Get_JobById(c echo.Context) error {
	result, err := getJobById(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func getJobById(c echo.Context) (model.ResponseData, error) {
	var res model.ResponseData

	jobid := c.Param("id")

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions/"
	endpoint := url + jobid

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Respo
	json.Unmarshal(responseData, &responseObject)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = responseObject

	return res, nil
}
