package entities

import "errors"

type ProductDomain struct {
	ID   int
	Name string
}

func NewProduct(id int, name string) (pd *ProductDomain, err error) {
	if id == 0 {
		return nil, errors.New("error id tidak boleh kosong")
	}

	if name == "agung" {
		return nil, errors.New("nama tidak boleh agung")
	}

	return &ProductDomain{
		ID:   id,
		Name: name,
	}, nil
}

func (pd ProductDomain) GetName() string {
	return pd.Name
}

func (pd ProductDomain) GetId() int {
	return pd.ID
}
