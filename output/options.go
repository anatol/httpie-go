package output

type Options struct {
	PrintRequestHeader  bool
	PrintRequestBody    bool
	PrintResponseHeader bool
	PrintResponseBody   bool

	EnableFormat bool
	EnableColor  bool
}
