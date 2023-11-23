package franchise

type IncompleteFranchise struct {
	Data FranchiseDTO
}

func NewIncompleteFranchise(franchiseDTO FranchiseDTO) IncompleteFranchise {
	return IncompleteFranchise{
		Data: franchiseDTO,
	}
}
