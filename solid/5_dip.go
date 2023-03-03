package solid

// DIP - Dependency Inversion Principle
// is a software design principle that states that high-level modules
// should not depend on low-level modules, but both should depend on abstractions.
// This principle aims to reduce coupling between components and increase
// flexibility and modifiability of software systems.

type EmailNotifier interface {
	SendOrderConfirmationEmail(customerName string, orderID string) error
}

type CustomerOrderProcessor struct {
	emailNotifier EmailNotifier
}

func (c *CustomerOrderProcessor) ProcessOrder(customerName string, orderID string) error {
	// Process the order here

	// Notify the customer via email
	err := c.emailNotifier.SendOrderConfirmationEmail(customerName, orderID)
	if err != nil {
		return err
	}

	return nil
}
