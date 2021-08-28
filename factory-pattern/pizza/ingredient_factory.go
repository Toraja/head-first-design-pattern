package pizza

type IngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateClam() Topping
	CreateVeggie() Topping
}

type NYIngredientFactory struct{}

func NewNYIngredientFactory() NYIngredientFactory {
	return NYIngredientFactory{}
}

func (nf NYIngredientFactory) CreateDough() Dough {
	return ThinCrustDough{}
}

func (nf NYIngredientFactory) CreateSauce() Sauce {
	return MarinaraSauce{}
}

func (nf NYIngredientFactory) CreateCheese() Cheese {
	return ReggianoCheese{}
}

func (nf NYIngredientFactory) CreateClam() Topping {
	return FreshClam{}
}

func (nf NYIngredientFactory) CreateVeggie() Topping {
	return OrganicVeggie{}
}

type ChicagoIngredientFactory struct{}

func NewChicagoIngredientFactory() ChicagoIngredientFactory {
	return ChicagoIngredientFactory{}
}

func (cf ChicagoIngredientFactory) CreateDough() Dough {
	return ThickCrustDough{}
}

func (cf ChicagoIngredientFactory) CreateSauce() Sauce {
	return PlumTomatoSauce{}
}

func (cf ChicagoIngredientFactory) CreateCheese() Cheese {
	return MozzarellaCheese{}
}

func (cf ChicagoIngredientFactory) CreateClam() Topping {
	return FrozenClam{}
}

func (cf ChicagoIngredientFactory) CreateVeggie() Topping {
	return DecentVeggie{}
}
