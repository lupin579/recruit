package model

type Commodity struct {
	CmdtID      int     `json:"cmdtID" db:"cmdt_id"`
	CmdtName    string  `json:"cmdtName" db:"cmdt_name"`
	Price       int     `json:"price" db:"price"`
	Quantity    int     `json:"quantity" db:"quantity"`
	OnSale      int8    `json:"onSale" db:"on_sale"`
	OwnerID     int     `json:"ownerID" db:"owner_id"`
	Description *string `json:"description" db:"description"`
}
