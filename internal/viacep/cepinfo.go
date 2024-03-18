package viacep

type CepInfo struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"Logradouro"`
	Complemento string `json:"Complemento"`
	Bairro      string `json:"Bairro"`
	Localidade  string `json:"Localidade"`
	Uf          string `json:"Uf"`
	Ibge        string `json:"Ibge"`
	Gia         string `json:"Gia"`
	Ddd         string `json:"Ddd"`
	Siafi       string `json:"Siafi"`
}
