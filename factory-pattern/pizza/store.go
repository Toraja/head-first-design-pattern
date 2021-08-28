package pizza

import "fmt"

// While in the book, there are concrete store classes for each region, here there
// is only one concrete store to simplify the code.

type Store interface {
	OrderPizza(pizzaType string) Pizza
	CreatePizza(pizzaType string) Pizza
}

type storeImpl struct {
	ingredientFactory IngredientFactory
}

func NewStore(ingredientFactory IngredientFactory) Store {
	return storeImpl{
		ingredientFactory: ingredientFactory,
	}
}

func (s storeImpl) OrderPizza(pizzaType string) Pizza {
	pizza := s.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (s storeImpl) CreatePizza(pizzaType string) Pizza {
	switch pizzaType {
	case "cheese":
		return NewCheesePizza(s.ingredientFactory)
	case "clam":
		return NewClamPizza(s.ingredientFactory)
	case "veggie":
		return NewVeggiePizza(s.ingredientFactory)
	default:
		panic(fmt.Sprintf("Unknows pizza type: %s", pizzaType))
	}
}
