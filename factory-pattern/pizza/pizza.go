package pizza

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type pizzaImpl struct {
	typ      string
	dough    Dough
	sauce    Sauce
	cheese   Cheese
	toppings []Topping
}

func (p pizzaImpl) Prepare() {
	fmt.Printf("Preparing %s pizza... done!\n", p.typ)
}

func (p pizzaImpl) Bake() {
	fmt.Println("Baking for 25 minutes at 350 degree.")
}

func (p pizzaImpl) Cut() {
	fmt.Println("Slicing into 6 pieces")
}

func (p pizzaImpl) Box() {
	fmt.Println("Placed in a box.")
}

func NewCheesePizza(factory IngredientFactory) Pizza {
	return pizzaImpl{
		typ:    "cheese",
		dough:  factory.CreateDough(),
		sauce:  factory.CreateSauce(),
		cheese: factory.CreateCheese(),
	}
}

func NewClamPizza(factory IngredientFactory) Pizza {
	return pizzaImpl{
		typ:      "clam",
		dough:    factory.CreateDough(),
		sauce:    factory.CreateSauce(),
		cheese:   factory.CreateCheese(),
		toppings: []Topping{factory.CreateClam()},
	}
}

func NewVeggiePizza(factory IngredientFactory) Pizza {
	return pizzaImpl{
		typ:      "veggie",
		dough:    factory.CreateDough(),
		sauce:    factory.CreateSauce(),
		cheese:   factory.CreateCheese(),
		toppings: []Topping{factory.CreateVeggie()},
	}
}
