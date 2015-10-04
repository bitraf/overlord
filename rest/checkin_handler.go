package rest

/*
import (
	"fmt"
	"net/http"
	"time"

	"github.com/bitraf/overlord/db"
	"github.com/labstack/echo"
)

type Checkin struct {
	Id        int       `json:"id"`
	Account   int       `json:"account"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

type Checkins struct {
	From   int       `json:"from"`
	Count  int       `json:"count"`
	Total  int       `json:"total"`
	Result []Checkin `json:"result"`
}

func toCheckin(row db.Checkin) Checkin {
	return Checkin{
		Id:        row.Id,
		Account:   row.Account.Id,
		Timestamp: row.Date,
		Type:      row.Type,
	}
}

func toCheckins(entry db.CheckinQuery) Checkins {
	ret := Checkins{}
	ret.From = entry.From
	ret.Count = entry.Count
	ret.Total = entry.Total
	ret.Result = make([]Checkin, 0)

	for _, value := range entry.Entries {
		ret.Result = append(ret.Result, toCheckin(value))
	}

	return ret
}

func (server *Server) getCheckin(c *echo.Context) *echo.HTTPError {
	id := intParam(c, "id", 0)
	row := db.Checkin{Id: id}
	has := server.db.FindCheckin(&row)

	if has {
		return c.JSON(http.StatusOK, toCheckin(row))
	}

	return server.error(c, http.StatusNotFound,
		fmt.Sprintf("Could not find checkin [%v]", id))
}

func (server *Server) getCheckins(c *echo.Context) *echo.HTTPError {
	query := db.CheckinQuery{}
	println(intParam(c, "from", 0))
	query.From = intParam(c, "from", 0)
	query.Count = intParam(c, "count", 10)

	server.db.FindCheckins(&query)

	return c.JSON(http.StatusOK, toCheckins(query))
}
*/
