package customer

type getCustomersRequest struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type getAddCustomerRequest struct {
	//ID           primitive.ObjectID
	//ID           string
	Name        string `json:"name" form:"name" query:"name"`
	Email       string `json:"email" form:"email" query:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number"`
	Password    string `json:"password" form:"password" query:"password"`
	//CreatedAt   time.Time
	//UpdateAt    time.Time
}

type updateCustomerRequest struct {
	ID          string `json:"id" form:"id" query:"id"`
	Name        string `json:"name" form:"name" query:"name"`
	Email       string `json:"email" form:"email" query:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number"`
	Password    string `json:"password" form:"password" query:"password"`
}

//Embeddings  []float64       `json:"embeddings"`
//Addresses   []model.Address `json:"addresses"`
//Tags        []string        `json:"tags"`
