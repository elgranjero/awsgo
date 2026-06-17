package ec2

// ModifyFleet is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the specified EC2 Fleet.
//
// You can only modify an EC2 Fleet request of type maintain .
//
// While the EC2 Fleet is being modified, it is in the modifying state.
//
// To scale up your EC2 Fleet, increase its target capacity. The EC2 Fleet
// launches the additional Spot Instances according to the allocation strategy for
// the EC2 Fleet request. If the allocation strategy is lowest-price , the EC2
// Fleet launches instances using the Spot Instance pool with the lowest price. If
// the allocation strategy is diversified , the EC2 Fleet distributes the instances
// across the Spot Instance pools. If the allocation strategy is capacity-optimized
// , EC2 Fleet launches instances from Spot Instance pools with optimal capacity
// for the number of instances that are launching.
//
// To scale down your EC2 Fleet, decrease its target capacity. First, the EC2
// Fleet cancels any open requests that exceed the new target capacity. You can
// request that the EC2 Fleet terminate Spot Instances until the size of the fleet
// no longer exceeds the new target capacity. If the allocation strategy is
// lowest-price , the EC2 Fleet terminates the instances with the highest price per
// unit. If the allocation strategy is capacity-optimized , the EC2 Fleet
// terminates the instances in the Spot Instance pools that have the least
// available Spot Instance capacity. If the allocation strategy is diversified ,
// the EC2 Fleet terminates instances across the Spot Instance pools.
// Alternatively, you can request that the EC2 Fleet keep the fleet at its current
// size, but not replace any Spot Instances that are interrupted or that you
// terminate manually.
//
// If you are finished with your EC2 Fleet for now, but will use it again later,
// you can set the target capacity to 0.
