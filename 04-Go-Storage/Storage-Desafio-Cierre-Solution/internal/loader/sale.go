package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewLoaderSaleJSON returns a new pointer to a LoaderSaleJSON struct.
func NewLoaderSaleJSON(file *os.File) *LoaderSaleJSON {
	return &LoaderSaleJSON{file: file}
}

// LoaderSaleJSON is an struct that implements the LoaderSale interface.
type LoaderSaleJSON struct {
	// file is the file to handle read and write operations.
	file *os.File
}

// SaleJSON is the struct that represents the sale data in the json file.
type SaleJSON struct {
	Id         int     `json:"id"`
	Quantity   int     `json:"quantity"`
	ProductId  int     `json:"product_id"`
	InvoiceId  int     `json:"invoice_id"`
}

// Load loads the sale data from the json file.
func (l *LoaderSaleJSON) Load() (s []internal.Sale, err error) {
	// decode the json file
	var ss []SaleJSON
	err = json.NewDecoder(l.file).Decode(&ss)
	if err != nil {
		return
	}

	// serialize the sale data
	for _, v := range ss {
		s = append(s, internal.Sale{
			Id: v.Id,
			SaleAttributes: internal.SaleAttributes{
				Quantity:  v.Quantity,
				ProductId: v.ProductId,
				InvoiceId: v.InvoiceId,
			},
		})
	}

	return
}