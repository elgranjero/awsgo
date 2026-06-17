package connectcases

// CreateRelatedItem is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// Creates a related item (comments, tasks, and contacts) and associates it with a
// case.
//
// There's a quota for the number of fields allowed in a Custom type related item.
// See [Amazon Connect Cases quotas].
//
// # Use cases
//
// Following are examples of related items that you may want to associate with a
// case:
//
// - Related contacts, such as calls, chats, emails tasks
//
// - Comments, for agent notes
//
// - SLAs, to capture target resolution goals
//
// - Cases, to capture related Amazon Connect Cases
//
// - Files, such as policy documentation or customer-provided attachments
//
// - Custom related items, which provide flexibility for you to define related
// items that such as bookings, orders, products, notices, and more
//
// Important things to know
//
// - If you are associating a contact to a case by passing in Contact for a type
// , you must have [DescribeContact]permission on the ARN of the contact that you provide in
// content.contact.contactArn .
//
// - A Related Item is a resource that is associated with a case. It may or may
// not have an external identifier linking it to an external resource (for example,
// a contactArn ). All Related Items have their own internal identifier, the
// relatedItemArn . Examples of related items include comments and contacts .
//
// - If you provide a value for performedBy.userArn you must also have [DescribeUser]
// permission on the ARN of the user that you provide.
//
// - The type field is reserved for internal use only.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
// [Amazon Connect Cases quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#cases-quotas
// [DescribeContact]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeContact.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
