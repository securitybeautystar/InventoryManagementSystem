package handlers_products

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/database"
	"github.com/LeonLow97/inventory-management-system-golang-react-postgresql/utils"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {

	// CheckUserGroup: IMS User and Operations
	if !CheckProductsUserGroup(w, req) {return}

	// Retrieve products from SQL database
	var data []handlers.Product
	// To handle nullable columns in a database table
	var productName, productDescription, productSku, productColour, productCategory, productBrand sql.NullString
	var productCost sql.NullFloat64

	rows, err := database.GetProducts()
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in GetProducts: ", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&productName, &productDescription, &productSku, &productColour, &productCategory, &productBrand, &productCost)
		if err != nil {
			utils.InternalServerError(w, "Internal Server Error in Scanning GetProducts: ", err)
			return
		}

		response := handlers.Product{
			ProductName: productName.String,
			ProductDescription: productDescription.String,
			ProductSku: productSku.String,
			ProductColour: productColour.String,
			ProductCategory: productCategory.String,
			ProductBrand: productBrand.String,
			ProductCost: float32(productCost.Float64),
		}

		data = append(data, response)
	}
	jsonStatus := struct {
		Code int `json:"code"`
		Response []handlers.Product `json:"response"`
	}{
		Response: data,
		Code: http.StatusOK,
	}

	bs, err := json.Marshal(jsonStatus);
	if err != nil {
		utils.InternalServerError(w, "Internal Server Error in Marshal JSON body in GetProducts: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(bs));
}

