package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Box model
type Box struct {
	IsFragile bool      `json:"isFragile" example:"true"`
	CreatedAt time.Time `json:"createdAt" example:"2021-01-01T20:52:41.483612611Z"`
	Price     float32   `json:"price" example:"1.23"`
}

// Service layer
type Service struct {
	res Response
}

// BoxRequest request for new box
type BoxRequest struct {
	IsFragile bool    `json:"isFragile" example:"true"`
	Price     float32 `json:"price" example:"1.23"`
}

// BigBoxRequest for new bix box
type BigBoxRequest struct {
	Boxes []Box  `json:"boxes"`
	What  string `json:"what"`
	If    struct {
		There []string `json:"there"`
		Are   struct {
			Some  bool   `json:"some"`
			Scary string `json:"scary"`
			Stuff struct {
				Inside bool `json:"inside"`
			} `json:"stuff"`
		} `json:"are"`
	} `json:"if"`
}

// BoxesResponse response of boxes
type BoxesResponse struct {
	Boxes []Box `json:"boxes"`
}

// New create new box
// @Summary Create new box
// @Schemes
// @Description creates new box
// @Tags box
// @Accept json
// @Produce json
// @Param boxRequest body service.BoxRequest true "New Box Request"
// @Success 201 {object} service.Response
// @Failure 500 {object} service.Response
// @Router /boxes [post]
func (s Service) New(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("demo") != true {
		c.AbortWithStatusJSON(http.StatusUnauthorized, s.res.error("unauthorized"))
		return
	}

	var req BoxRequest
	err := c.BindJSON(&req)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, s.res.internal())
		return
	}

	var w Box
	w.IsFragile = req.IsFragile
	w.Price = req.Price
	w.CreatedAt = time.Now().UTC().Local()

	// create new box
	c.JSON(http.StatusCreated, s.res.success())
}

// List returns all boxes
// @Summary List all boxes
// @Schemes
// @Description Returns a list of boxes
// @Tags box
// @Accept json
// @Produce json
// @Success 200 {object} service.BoxesResponse
// @Failure 500 {object} service.Response
// @Router /boxes [get]
func (s Service) List(c *gin.Context) {

	boxes := []Box{
		{
			IsFragile: true,
			Price:     5.77,
			CreatedAt: time.Now().Local().UTC(),
		},
		{
			IsFragile: false,
			Price:     1.23,
			CreatedAt: time.Now().Local().UTC(),
		},
	}

	c.JSON(http.StatusOK, boxes)

}

// NewBigBox create Big Box
// @Summary Create new Big box
// @Schemes
// @Description creates new Big box
// @Tags box
// @Accept json
// @Produce json
// @Param boxRequest body service.BigBoxRequest true "New Big Box Request"
// @Success 201 {object} service.Response
// @Failure 500 {object} service.Response
// @Router /boxes/big [post]
func (s Service) NewBigBox(c *gin.Context) {
	c.JSON(http.StatusOK, s.res.success())
}
