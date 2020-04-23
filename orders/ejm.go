package orders

import (
	"net/http"

	"github.com/labstack/echo"
)

// User
type User struct {
	/*
		ProductName string  `json:"product_name" form:"product_name" query:"product_name"`
		Description string  `json:"description" form:"description" query:"description"`
		TotalAmount int     `json:"total_amount" form:"total_amount" query:"total_amount"`
		TotalSold   int     `json:"total_sold" form:"total_sold" query:"total_sold"`
		Price       float32 `json:"price" form:"price" query:"price"`
	*/
	Name string `json:"name" form:"name" query:"name"`
	//Description string  `json:"description" form:"description" query:"description"`
	//TotalAmount int     `json:"total_amount" form:"total_amount" query:"total_amount"`
	//TotalSold   int     `json:"total_sold" form:"total_sold" query:"total_sold"`
	//Price       float32 `json:"price" form:"price" query:"price"`
}

func Ro(r *echo.Group) {
	// Routes
	r.GET("/get1", GetProdu)

}

// Handler

func GetProdu(c echo.Context) (err error) {
	u := new(User)
	//var u User
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

// Handler
/*
func GetProdu(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}
*/

/*
	// If the file doesn't exist, create it, or append to the file

	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}


	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
*/
