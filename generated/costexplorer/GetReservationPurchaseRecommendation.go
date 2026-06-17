package costexplorer

// GetReservationPurchaseRecommendation is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Gets recommendations for reservation purchases. These recommendations might
// help you to reduce your costs. Reservations provide a discounted hourly rate (up
// to 75%) compared to On-Demand pricing.
//
// Amazon Web Services generates your recommendations by identifying your
// On-Demand usage during a specific time period and collecting your usage into
// categories that are eligible for a reservation. After Amazon Web Services has
// these categories, it simulates every combination of reservations in each
// category of usage to identify the best number of each type of Reserved Instance
// (RI) to purchase to maximize your estimated savings.
//
// For example, Amazon Web Services automatically aggregates your Amazon EC2
// Linux, shared tenancy, and c4 family usage in the US West (Oregon) Region and
// recommends that you buy size-flexible regional reservations to apply to the c4
// family usage. Amazon Web Services recommends the smallest size instance in an
// instance family. This makes it easier to purchase a size-flexible Reserved
// Instance (RI). Amazon Web Services also shows the equal number of normalized
// units. This way, you can purchase any instance size that you want. For this
// example, your RI recommendation is for c4.large because that is the smallest
// size instance in the c4 instance family.
