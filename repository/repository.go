package repository

import (
	"errors"

	"go-theapi/model"
	"go-theapi/token"
	"go-theapi/util"

	"github.com/labstack/echo/v4"
)

func GenerateToken(c echo.Context, user string) (*model.M_Response_session, error) {
	var session model.M_Response_session

	config, _ := util.LoadConfig(".")
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		err := errors.New("cannot create token maker")
		return &session, err
	}

	accessToken, accessPayload, err := tokenMaker.CreateToken(
		user,
		config.AccessTokenDuration,
	)
	if err != nil {
		err := errors.New("cannot generate access token")
		return &session, err
	}

	result := &model.M_Response_session{
		SessionID:        accessPayload.ID,
		Username:         user,
		AccessToken:      accessToken,
		Access_ExpiresAt: accessPayload.ExpiredAt,
	}

	return result, nil
}

// func FetchAllJobs(params model.M_page_param, limit int, search string) ([]model.M_Job_List, int64, int, error) {
// 	var pp model.M_Job_List
// 	var data []model.M_Job_List
// 	var offset string
// 	var total int

// 	if params.PageNumber >= 2 {
// 		offset = " offset " + strconv.Itoa((params.PageNumber-1)*limit)
// 	} else {
// 		offset = " offset 0"
// 	}

// 	var dateCond string

// 	var sqlStatement string
// 	if len(search) != 0 {

// 		db := db.ConnectPostgresql()

// 		sqlStatement = `select p.id, p.p_name, p.p_code, to_char(p.created_at at time ZONE 'Asia/Jakarta', 'dd-mm-yyyy'), CASE WHEN p.valid_start = 'infinity' OR p.valid_end = 'infinity' THEN 'infinity - infinity' ELSE COALESCE(to_char(p.valid_start, 'dd-mm-yyyy') || ' - '  || to_char(p.valid_end, 'dd-mm-yyyy'),'') END ,
// 		COALESCE(pd.original_transaction_quota,0), COALESCE(pd.usage,0), CASE WHEN p.valid_start > now() THEN 'Upcoming' WHEN now() BETWEEN p.valid_start AND p.valid_end THEN 'Active' ELSE 'Ended' END AS status
// 		from promo_nextgen p inner join promo_detail_nextgen pd on p.id = pd.id WHERE (p.p_name ilike '%' || $1 || '%' OR p.p_code ilike '%' || $1 || '%')` + dateCond + `ORDER BY ` + sortQuery.Field + ` ` + sortQuery.Order + ` LIMIT ` + strconv.Itoa(limit) + offset

// 		rows, err := db.Query(sqlStatement, search)
// 		defer db.Close()
// 		if err != nil {
// 			if err != sql.ErrNoRows {
// 				return nil, 0, 0, err
// 			}
// 		}
// 		defer rows.Close()

// 		if err != sql.ErrNoRows {
// 			for rows.Next() {
// 				err = rows.Scan(&pp.Promo_Id, &pp.Promo_Name, &pp.Promo_Code, &pp.Created_Date, &pp.Period, &pp.Quota, &pp.Usage, &pp.Status)

// 				if err != nil {
// 					return nil, 0, 0, err
// 				}

// 				if pp.Quota == "0" {
// 					pp.Quota = "Unlimited"
// 				}

// 				data = append(data, pp)
// 			}
// 		}

// 		count := db.QueryRow(`select count(*) from  promo_nextgen p inner join promo_detail_nextgen pd on p.id = pd.id where (p.p_name ilike '%' || $1 || '%' OR p.p_code ilike '%' || $1 || '%')`, search)
// 		_ = count.Scan(&total)
// 	} else {
// 		if len(dp.DateFrom) != 0 && len(dp.DateTo) != 0 {
// 			dateCond = " WHERE p.valid_start >= '" + dp.DateFrom + "' AND p.valid_end <= '" + dp.DateTo + "T23:59:59' "
// 		} else {
// 			dateCond = " "
// 		}

// 		db := database.ConnectPostgresql()
// 		sqlStatement = `select p.id, p.p_name, p.p_code, to_char(p.created_at at time ZONE 'Asia/Jakarta', 'dd-mm-yyyy'), CASE WHEN p.valid_start = 'infinity' OR p.valid_end = 'infinity' THEN 'infinity - infinity' ELSE COALESCE(to_char(p.valid_start, 'dd-mm-yyyy') || ' - '  || to_char(p.valid_end, 'dd-mm-yyyy'),'') END,
// 		COALESCE(pd.original_transaction_quota,0), COALESCE(pd.usage,0), CASE WHEN p.valid_start > now() THEN 'Upcoming' WHEN now() BETWEEN p.valid_start AND p.valid_end THEN 'Active' ELSE 'Ended' END AS status
// 		from promo_nextgen p inner join promo_detail_nextgen pd on p.id = pd.id` + dateCond + `ORDER BY ` + sortQuery.Field + ` ` + sortQuery.Order + ` LIMIT ` + strconv.Itoa(limit) + offset
// 		rows, err := db.Query(sqlStatement)
// 		defer db.Close()
// 		if err != nil {
// 			if err != sql.ErrNoRows {
// 				return nil, 0, 0, err
// 			}
// 		}
// 		defer rows.Close()

// 		if err != sql.ErrNoRows {
// 			for rows.Next() {
// 				err = rows.Scan(&pp.Promo_Id, &pp.Promo_Name, &pp.Promo_Code, &pp.Created_Date, &pp.Period, &pp.Quota, &pp.Usage, &pp.Status)

// 				if err != nil {
// 					return nil, 0, 0, err
// 				}

// 				if pp.Quota == "0" {
// 					pp.Quota = "Unlimited"
// 				}

// 				data = append(data, pp)

// 			}
// 		}

// 		count := db.QueryRow(`select count(*) from  promo_nextgen p inner join promo_detail_nextgen pd on p.id = pd.id`)
// 		_ = count.Scan(&total)
// 	}

// 	var nextPage int64
// 	if len(data) > 0 {
// 		nextPage = int64(params.PageNumber + 1)
// 	}

// 	return data, nextPage, total, nil
// }
