package waf

// UpdateGeoMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes GeoMatchConstraint objects in an GeoMatchSet . For each GeoMatchConstraint
// object, you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change an GeoMatchConstraint object, you delete the existing object and add a
// new one.
//
// - The Type . The only valid value for Type is Country .
//
// - The Value , which is a two character code for the country to add to the
// GeoMatchConstraint object. Valid codes are listed in GeoMatchConstraint$Value.
//
// To create and configure an GeoMatchSet , perform the following steps:
//
// - Submit a CreateGeoMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateGeoMatchSetrequest.
//
// - Submit an UpdateGeoMatchSet request to specify the country that you want AWS
// WAF to watch for.
//
// When you update an GeoMatchSet , you specify the country that you want to add
// and/or the country that you want to delete. If you want to change a country, you
// delete the existing country and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
