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
	m := map[string]int{}
	return m
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	units = Units()
	amount, exists := units[unit]
	if exists == false {
		return false
	} else {
		_, exist := bill[item]
		if exist == true {
			bill[item] += amount
		} else {
			bill[item] = amount
		}
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
	amount, exist := units[unit]
	if exists == false || exist == false {
		return false
	}
	amount2 := bill[item]
	if amount2-amount < 0 {
		return false
	} else if amount2-amount == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = amount2 - amount
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if exists == false {
		return 0, false
	}
	return quantity, exists
}
