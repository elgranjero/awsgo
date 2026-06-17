package ec2

// ModifySpotFleetRequest is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the specified Spot Fleet request.
//
// You can only modify a Spot Fleet request of type maintain .
//
// While the Spot Fleet request is being modified, it is in the modifying state.
//
// To scale up your Spot Fleet, increase its target capacity. The Spot Fleet
// launches the additional Spot Instances according to the allocation strategy for
// the Spot Fleet request. If the allocation strategy is lowestPrice , the Spot
// Fleet launches instances using the Spot Instance pool with the lowest price. If
// the allocation strategy is diversified , the Spot Fleet distributes the
// instances across the Spot Instance pools. If the allocation strategy is
// capacityOptimized , Spot Fleet launches instances from Spot Instance pools with
// optimal capacity for the number of instances that are launching.
//
// To scale down your Spot Fleet, decrease its target capacity. First, the Spot
// Fleet cancels any open requests that exceed the new target capacity. You can
// request that the Spot Fleet terminate Spot Instances until the size of the fleet
// no longer exceeds the new target capacity. If the allocation strategy is
// lowestPrice , the Spot Fleet terminates the instances with the highest price per
// unit. If the allocation strategy is capacityOptimized , the Spot Fleet
// terminates the instances in the Spot Instance pools that have the least
// available Spot Instance capacity. If the allocation strategy is diversified ,
// the Spot Fleet terminates instances across the Spot Instance pools.
// Alternatively, you can request that the Spot Fleet keep the fleet at its current
// size, but not replace any Spot Instances that are interrupted or that you
// terminate manually.
//
// If you are finished with your Spot Fleet for now, but will use it again later,
// you can set the target capacity to 0.
