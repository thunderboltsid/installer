package v1

// AWSResourceReference is a reference to a specific AWS resource by ID, ARN, or filters.
// Only one of ID, ARN or Filters may be specified. Specifying more than one will result in
// a validation error.
// +union
type AWSResourceReference struct {
	// Type determines how the reference will fetch the AWS resource.
	// +unionDiscriminator
	// +kubebuilder:validation:Enum:="id";"arn";"filters"
	// +kubebuilder:validation:Required
	Type AWSResourceReferenceType `json:"type"`
	// ID of resource
	// +optional
	ID *string `json:"id,omitempty"`
	// ARN of resource
	// +optional
	ARN *string `json:"arn,omitempty"`
	// Filters is a set of filters used to identify a resource
	// +optional
	Filters *[]AWSResourceFilter `json:"filters,omitempty"`
}

// AWSResourceReferenceType is an enumeration of different resource reference types.
type AWSResourceReferenceType string

const (
	// AWSIDReferenceType is a resource reference based on the object ID.
	AWSIDReferenceType AWSResourceReferenceType = "id"

	// AWSARNReferenceType is a resource reference based on the object ARN.
	AWSARNReferenceType AWSResourceReferenceType = "arn"

	// AWSFiltersReferenceType is a resource reference based on filters.
	AWSFiltersReferenceType AWSResourceReferenceType = "filters"
)

// AWSResourceFilter is a filter used to identify an AWS resource
type AWSResourceFilter struct {
	// Name of the filter. Filter names are case-sensitive.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Values includes one or more filter values. Filter values are case-sensitive.
	// +optional
	Values []string `json:"values,omitempty"`
}
