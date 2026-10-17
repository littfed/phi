// Package permission implements the tool-execution gate: policy modes,
// decisions, and the user approval flow.
//
// It knows nothing about individual tools. Callers hand it a Request built by
// the tool itself (see tooldef.WithPermission); an undeclared tool arrives as
// ActionUnknown, which asks rather than trusting the tool's name.
package permission
