package model

type M_page_param struct {
	PageNumber int
}

type M_Job_List struct {
	Category        string `json:"category" form:"category" query:"category"`
	BrandId         int    `json:"brand_id" form:"brand_id" query:"brand_id"`
	Brand           string `json:"brand" form:"brand" query:"brand"`
	Product_Id      int    `json:"product_id" form:"product_id" query:"product_id"`
	Product_Code    string `json:"product_code" form:"product_code" query:"product_code"`
	Product_Name    string `json:"product_name" form:"product_name" query:"product_name"`
	Daruma_Sku      string `json:"daruma_sku" form:"daruma_sku" query:"daruma_sku"`
	Mpn             string `json:"mpn" form:"mpn" query:"mpn"`
	Price           int64  `json:"price" form:"price" query:"price"`
	Discount_Price  int64  `json:"discount_price" form:"discount_price" query:"discount_price"`
	Price_Post_Tax  int64  `json:"price_post_tax" form:"price_post_tax" query:"price_post_tax"`
	Supplier_Price  int64  `json:"supplier_price" form:"supplier_price" query:"supplier_price"`
	Supplier_Name   string `json:"supplier_name" form:"supplier_name" query:"supplier_name"`
	Stock_Type      string `json:"stock_type" form:"stock_type" query:"stock_type"`
	Inventory_Stock int    `json:"inventory_stock" form:"inventory_stock" query:"inventory_stock"`
	Supplier_Stock  int    `json:"supplier_stock" form:"supplier_stock" query:"supplier_stock"`
	Indent_Stock    int    `json:"indent_stock" form:"indent_stock" query:"indent_stock"`
}
