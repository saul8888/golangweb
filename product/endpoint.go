package product

type getProductsRequest struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type getAddProductRequest struct {
	//ID           primitive.ObjectID
	//ID           string
	ProductName string  `json:"product_name" form:"product_name" query:"product_name"`
	Description string  `json:"description" form:"description" query:"description"`
	TotalAmount int     `json:"total_amount" form:"total_amount" query:"total_amount"`
	TotalSold   int     `json:"total_sold" form:"total_sold" query:"total_sold"`
	Price       float32 `json:"price" form:"price" query:"price"`
	//CreatedAt   time.Time
	//UpdateAt    time.Time
}

type updateProductRequest struct {
	ID          string  `json:"id" form:"id" query:"id"`
	ProductName string  `json:"product_name" form:"product_name" query:"product_name"`
	Description string  `json:"description" form:"description" query:"description"`
	TotalAmount int     `json:"total_amount" form:"total_amount" query:"total_amount"`
	TotalSold   int     `json:"total_sold" form:"total_sold" query:"total_sold"`
	Price       float32 `json:"price" form:"price" query:"price"`
}
