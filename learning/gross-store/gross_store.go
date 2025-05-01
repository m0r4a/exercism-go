package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	b := make(map[string]int)

	return b
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] += u

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// check if the item is on the bill
	a, exists := bill[item]

	if !exists {
		return false
	}

	// check if the unit is on the units map
	u, exists := units[unit]

	if !exists {
		return false
	}

	// check if the new quantity would be less than 0
	if a-u < 0 {
		return false
	}

	// Remove the item if the amount of the items is 0
	if (a - u) == 0 {
		delete(bill, item)

		return true
	}

	bill[item] -= u

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	a, b := bill[item]

	return a, b
}
