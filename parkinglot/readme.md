Design a parking lot system that supports parking and unparking vehicles in a clean, extensible, and maintainable way.

The system should efficiently manage parking slots across multiple floors, handle different vehicle types, calculate parking charges, and integrate with a payment system.

Objectives
Allow vehicles to enter and exit the parking lot
Allocate appropriate parking slots based on vehicle type
Support multiple floors
Track parking duration and calculate charges
Process payment before exit
Ensure system is scalable and easy to extend

Supported Vehicle Types
Bike
Car
Truck

Each vehicle type may have different parking requirements.

Parking Rules
Each slot is designed for a specific vehicle type
A vehicle can be parked in:
its own type slot, or
a larger slot (e.g., Car in Truck slot)
A slot can hold only one vehicle at a time
A vehicle receives a ticket upon entry
The ticket is required to exit
_________________________________________________

Vehicle Type -> Bike Car Truck
Vehicle ->  Reg number, Type

Vehicle enters and takes a TICKET -> ticket has a ticket number (uuid) 
Ticket can be hourly or day pass. so hourly prices, daily prices.. how define the prices ?? Vehicle defines the prices..

Vehicle locks a SLOT. slot has ID and Type.... how do we know slot is locked.. Slot has a state.. VACANT/OCCUPIED

Vehicle exits.. gives ticket.. we check the "enter" timestamp and "Exit" timestamp. multiple by vehicle. calculate prices
 
PaymentManger -> calculate per vehicle & time 

