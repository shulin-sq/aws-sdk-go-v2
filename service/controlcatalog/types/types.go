// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A summary of the domain that a common control or an objective belongs to.
type AssociatedDomainSummary struct {

	// The Amazon Resource Name (ARN) of the related domain.
	Arn *string

	// The name of the related domain.
	Name *string

	noSmithyDocumentSerde
}

// A summary of the objective that a common control supports.
type AssociatedObjectiveSummary struct {

	// The Amazon Resource Name (ARN) of the related objective.
	Arn *string

	// The name of the related objective.
	Name *string

	noSmithyDocumentSerde
}

// An optional filter that narrows the results to a specific objective.
type CommonControlFilter struct {

	// The objective that's used as filter criteria. You can use this parameter to
	// specify one objective ARN at a time. Passing multiple ARNs in the
	// CommonControlFilter isn’t currently supported.
	Objectives []ObjectiveResourceFilter

	noSmithyDocumentSerde
}

// A summary of metadata for a common control.
type CommonControlSummary struct {

	// The Amazon Resource Name (ARN) that identifies the common control.
	//
	// This member is required.
	Arn *string

	// The time when the common control was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The description of the common control.
	//
	// This member is required.
	Description *string

	// The domain that the common control belongs to.
	//
	// This member is required.
	Domain *AssociatedDomainSummary

	// The time when the common control was most recently updated.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The name of the common control.
	//
	// This member is required.
	Name *string

	// The objective that the common control belongs to.
	//
	// This member is required.
	Objective *AssociatedObjectiveSummary

	noSmithyDocumentSerde
}

// The domain resource that's being used as a filter.
type DomainResourceFilter struct {

	// The Amazon Resource Name (ARN) of the domain.
	Arn *string

	noSmithyDocumentSerde
}

// A summary of metadata for a domain.
type DomainSummary struct {

	// The Amazon Resource Name (ARN) that identifies the domain.
	//
	// This member is required.
	Arn *string

	// The time when the domain was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The description of the domain.
	//
	// This member is required.
	Description *string

	// The time when the domain was most recently updated.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The name of the domain.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// An optional filter that narrows the list of objectives to a specific domain.
type ObjectiveFilter struct {

	// The domain that's used as filter criteria. You can use this parameter to
	// specify one domain ARN at a time. Passing multiple ARNs in the ObjectiveFilter
	// isn’t currently supported.
	Domains []DomainResourceFilter

	noSmithyDocumentSerde
}

// The objective resource that's being used as a filter.
type ObjectiveResourceFilter struct {

	// The Amazon Resource Name (ARN) of the objective.
	Arn *string

	noSmithyDocumentSerde
}

// A summary of metadata for an objective.
type ObjectiveSummary struct {

	// The Amazon Resource Name (ARN) that identifies the objective.
	//
	// This member is required.
	Arn *string

	// The time when the objective was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The description of the objective.
	//
	// This member is required.
	Description *string

	// The domain that the objective belongs to.
	//
	// This member is required.
	Domain *AssociatedDomainSummary

	// The time when the objective was most recently updated.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The name of the objective.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
