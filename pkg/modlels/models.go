package models

// a data type which we'll use to send data to the templates from the handlers
type TemplateData struct {
	// maps of strings to stuff - very multipurpose
	// map[string]interface{} is a general purpose one
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	FlashMsg   string
	WarningMsg string
	Error      string
}
