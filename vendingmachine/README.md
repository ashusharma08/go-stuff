The Problem Statement
"Design and implement a thread-safe Vending Machine that handles multiple products, tracks inventory, and processes payments using different denominations."

Core Requirements
Inventory Management: The machine should hold multiple items (e.g., Coke, Pepsi, Chips) with different prices and quantities.

Payment Handling: Users can insert coins or notes (e.g., 1, 5, 10, 25 cents).

State Management: The machine must behave differently based on its current state:

Idle: Waiting for money.

Accepting Money: Updating the current balance.

Dispensing: Releasing the product and calculating change.

Transaction Integrity: If a product is out of stock or the user cancels, the money must be returned.


Product{
    Name, 
    Brand,
    Weight,
    Quantity,
    Slot?,
    ...
}

Vending machine{
    products []*Product
}

// Add a product to inventory
// Remove items from inventory
// Checkout  - diff payment methods.
// User?? i guess not needed...