package mturk

// CreateHIT is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The CreateHIT operation creates a new Human Intelligence Task (HIT). The new
// HIT is made available for Workers to find and accept on the Amazon Mechanical
// Turk website.
//
// This operation allows you to specify a new HIT by passing in values for the
// properties of the HIT, such as its title, reward amount and number of
// assignments. When you pass these values to CreateHIT , a new HIT is created for
// you, with a new HITTypeID . The HITTypeID can be used to create additional HITs
// in the future without needing to specify common parameters such as the title,
// description and reward amount each time.
//
// An alternative way to create HITs is to first generate a HITTypeID using the
// CreateHITType operation and then call the CreateHITWithHITType operation. This
// is the recommended best practice for Requesters who are creating large numbers
// of HITs.
//
// CreateHIT also supports several ways to provide question data: by providing a
// value for the Question parameter that fully specifies the contents of the HIT,
// or by providing a HitLayoutId and associated HitLayoutParameters .
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see [Amazon Mechanical Turk Pricing].
//
// [Amazon Mechanical Turk Pricing]: https://requester.mturk.com/pricing
