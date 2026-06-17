package translate

// TranslateDocument is generated as a reference stub.
// Executable command wiring lives under cmd/translate.go.
//
// Translates the input document from the source language to the target language.
// This synchronous operation supports text, HTML, or Word documents as the input
// document.
//
// TranslateDocument supports translations from English to any supported language,
// and from any supported language to English. Therefore, specify either the source
// language code or the target language code as “en” (English).
//
// If you set the Formality parameter, the request will fail if the target
// language does not support formality. For a list of target languages that support
// formality, see [Setting formality].
//
// [Setting formality]: https://docs.aws.amazon.com/translate/latest/dg/customizing-translations-formality.html
