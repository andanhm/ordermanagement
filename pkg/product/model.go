package product

// Details holds the product details
type Details struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Slug              string        `json:"slug"`
	Price             int           `json:"price"`
	DiscountPrice     int           `json:"discount_price"`
	Category          string        `json:"category"`
	Color             string        `json:"color"`
	Sizes             []string      `json:"sizes"`
	Subcategory       string        `json:"subcategory"`
	Sale              bool          `json:"sale"`
	Article           string        `json:"article"`
	Quantity          int           `json:"quantity"`
	Img               string        `json:"img"`
	Options           []interface{} `json:"options"`
	FulfilledByDuka   bool          `json:"fulfilled_by_duka"`
	ShippedFromAbroad bool          `json:"shipped_from_abroad"`
	DukaApproved      bool          `json:"duka_approved"`
	Vendor            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"vendor"`
	Ratings struct {
		StarRatings float64 `json:"star_ratings"`
		Votes       int     `json:"votes"`
	} `json:"ratings"`
}
