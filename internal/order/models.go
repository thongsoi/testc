package order

import (
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	ID     int    `htmx:"marketID" json:"marketID" db:"market_id"`
	EnName string `htmx:"enName" json:"enName" db:"en_name"`
}

type SubMarket struct {
	ID     int    `htmx:"submarketID" json:"submarketID" db:"submarket_id"`
	EnName string `htmx:"enName" json:"enName" db:"en_name"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	MarketID          int             `htmx:"marketID" json:"marketID" db:"market_id"`
	SubmarketID       int             `htmx:"submarketID" json:"submarketID" db:"submarket_id"`
	ProductID         string          `htmx:"productID" json:"productID" db:"product_id"`
	OrderTypeID       string          `htmx:"orderTypeID" json:"orderTypeID" db:"order_type_id"`
	ContractTypeID    int             `htmx:"contractTypeID" json:"contractTypeID" db:"contract_type_id"`
	QualityID         string          `htmx:"qualityID" json:"qualityID" db:"quality_id"`
	Price             decimal.Decimal `htmx:"price" json:"price" db:"price"`
	Quantity          decimal.Decimal `htmx:"quantity" json:"quantity" db:"quantity"`
	UomID             int             `htmx:"uomID" json:"uomID" db:"uom_id"`
	PackingID         int             `htmx:"packingID" json:"packingID" db:"packing_id"`
	PaymentTermID     int             `htmx:"paymentTermID" json:"paymentTermID" db:"payment_term_id"`
	DeliveryTermID    int             `htmx:"deliveryTermID" json:"deliveryTermID" db:"delivery_term_id"`
	ProvinceID        int             `htmx:"provinceID" json:"provinceID" db:"province_id"`
	DistrictID        int             `htmx:"districtID" json:"districtID" db:"district_id"`
	SubdistrictID     int             `htmx:"subdistrictID" json:"subdistrictID" db:"subdistrict_id"`
	FirstDeliveryDate *time.Time      `htmx:"firstDeliveryDate" json:"firstDeliveryDate" db:"first_delivery_date,omitempty"`
	LastDeliveryDate  *time.Time      `htmx:"lastDeliveryDate" json:"lastDeliveryDate" db:"last_delivery_date,omitempty"`
	Remark            *string         `htmx:"remark" json:"remark" db:"remark,omitempty"`
}
