package ft

import (
	"testing"

	"books/head-first-design-pattern/factory-pattern/pizza"

	"github.com/stretchr/testify/assert"
)

/**
 * There are 2 style of pizza store
 * - NewYork
 * - Chicago
 *
 * They create 3 types of pizza
 * - Cheese
 * - Clam
 * - Veggies
 *
 * Pizza is composed of the below ingredients
 * - dough
 * - sauce
 * - cheese
 * - ... other toppings depending on pizzas
 *
 * Each pizza needs to be created with local ingredients
 * - NewYork
 *   - Dough: Thin crust
 *   - Sauce: Marinara
 *   - Cheese: Reggiano
 *   - Clam: Fresh
 *   - Veggies: Organic
 * - Chicago
 *   - Dough: Thick crust
 *   - Sauce: Plum tomato
 *   - Cheese: Mozzarella
 *   - Clam: Frozen
 *   - Veggies: Decent
 */

func TestNYStoreCreatesNYStyleCheesePizza(t *testing.T) {
	nyStore := pizza.NewStore(pizza.NewNYIngredientFactory())
	assert.Equal(t, pizza.NewCheesePizza(pizza.NewNYIngredientFactory()), nyStore.CreatePizza("cheese"))
}

func TestNYStoreCreatesNYClamPizza(t *testing.T) {
	nyStore := pizza.NewStore(pizza.NewNYIngredientFactory())
	assert.Equal(t, pizza.NewClamPizza(pizza.NewNYIngredientFactory()), nyStore.CreatePizza("clam"))
}

func TestNYStoreCreatesNYVeggiePizza(t *testing.T) {
	nyStore := pizza.NewStore(pizza.NewNYIngredientFactory())
	assert.Equal(t, pizza.NewVeggiePizza(pizza.NewNYIngredientFactory()), nyStore.CreatePizza("veggie"))
}

func TestChicagoStoreCreatesChicagoStyleCheesePizza(t *testing.T) {
	chicagoStore := pizza.NewStore(pizza.NewChicagoIngredientFactory())
	assert.Equal(t, pizza.NewCheesePizza(pizza.NewChicagoIngredientFactory()), chicagoStore.CreatePizza("cheese"))
}

func TestChicagoStoreCreatesChicagoStyleClamPizza(t *testing.T) {
	chicagoStore := pizza.NewStore(pizza.NewChicagoIngredientFactory())
	assert.Equal(t, pizza.NewClamPizza(pizza.NewChicagoIngredientFactory()), chicagoStore.CreatePizza("clam"))
}

func TestChicagoStoreCreatesNYVeggiePizza(t *testing.T) {
	chicagoStore := pizza.NewStore(pizza.NewChicagoIngredientFactory())
	assert.Equal(t, pizza.NewVeggiePizza(pizza.NewChicagoIngredientFactory()), chicagoStore.CreatePizza("veggie"))
}

func TestNYIngredientFactoryProduceNYIngredients(t *testing.T) {
	nyFactory := pizza.NewNYIngredientFactory()
	assert.Equal(t, pizza.ThinCrustDough{}, nyFactory.CreateDough())
	assert.Equal(t, pizza.MarinaraSauce{}, nyFactory.CreateSauce())
	assert.Equal(t, pizza.ReggianoCheese{}, nyFactory.CreateCheese())
	assert.Equal(t, pizza.FreshClam{}, nyFactory.CreateClam())
	assert.Equal(t, pizza.OrganicVeggie{}, nyFactory.CreateVeggie())
}

func TestChicagoIngredientFactoryProduceChicagoIngredients(t *testing.T) {
	chicagoFactory := pizza.NewChicagoIngredientFactory()
	assert.Equal(t, pizza.ThickCrustDough{}, chicagoFactory.CreateDough())
	assert.Equal(t, pizza.PlumTomatoSauce{}, chicagoFactory.CreateSauce())
	assert.Equal(t, pizza.MozzarellaCheese{}, chicagoFactory.CreateCheese())
	assert.Equal(t, pizza.FrozenClam{}, chicagoFactory.CreateClam())
	assert.Equal(t, pizza.DecentVeggie{}, chicagoFactory.CreateVeggie())
}

func TestOrderPizzaAtNYStore(t *testing.T) {
	nyStore := pizza.NewStore(pizza.NewNYIngredientFactory())
	nyStore.OrderPizza("cheese")
	// assert using your eyes...
}
