// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

const (

	// ErrCodeAcceleratorNotDisabledException for service response error code
	// "AcceleratorNotDisabledException".
	//
	// The accelerator that you specified could not be disabled.
	ErrCodeAcceleratorNotDisabledException = "AcceleratorNotDisabledException"

	// ErrCodeAcceleratorNotFoundException for service response error code
	// "AcceleratorNotFoundException".
	//
	// The accelerator that you specified doesn't exist.
	ErrCodeAcceleratorNotFoundException = "AcceleratorNotFoundException"

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have access permission.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAssociatedEndpointGroupFoundException for service response error code
	// "AssociatedEndpointGroupFoundException".
	//
	// The listener that you specified has an endpoint group associated with it.
	// You must remove all dependent resources from a listener before you can delete
	// it.
	ErrCodeAssociatedEndpointGroupFoundException = "AssociatedEndpointGroupFoundException"

	// ErrCodeAssociatedListenerFoundException for service response error code
	// "AssociatedListenerFoundException".
	//
	// The accelerator that you specified has a listener associated with it. You
	// must remove all dependent resources from an accelerator before you can delete
	// it.
	ErrCodeAssociatedListenerFoundException = "AssociatedListenerFoundException"

	// ErrCodeByoipCidrNotFoundException for service response error code
	// "ByoipCidrNotFoundException".
	//
	// The CIDR that you specified was not found or is incorrect.
	ErrCodeByoipCidrNotFoundException = "ByoipCidrNotFoundException"

	// ErrCodeEndpointGroupAlreadyExistsException for service response error code
	// "EndpointGroupAlreadyExistsException".
	//
	// The endpoint group that you specified already exists.
	ErrCodeEndpointGroupAlreadyExistsException = "EndpointGroupAlreadyExistsException"

	// ErrCodeEndpointGroupNotFoundException for service response error code
	// "EndpointGroupNotFoundException".
	//
	// The endpoint group that you specified doesn't exist.
	ErrCodeEndpointGroupNotFoundException = "EndpointGroupNotFoundException"

	// ErrCodeIncorrectCidrStateException for service response error code
	// "IncorrectCidrStateException".
	//
	// The CIDR that you specified is not valid for this action. For example, the
	// state of the CIDR might be incorrect for this action.
	ErrCodeIncorrectCidrStateException = "IncorrectCidrStateException"

	// ErrCodeInternalServiceErrorException for service response error code
	// "InternalServiceErrorException".
	//
	// There was an internal error for AWS Global Accelerator.
	ErrCodeInternalServiceErrorException = "InternalServiceErrorException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// An argument that you specified is invalid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// There isn't another item to return.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidPortRangeException for service response error code
	// "InvalidPortRangeException".
	//
	// The port numbers that you specified are not valid numbers or are not unique
	// for this accelerator.
	ErrCodeInvalidPortRangeException = "InvalidPortRangeException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Processing your request would cause you to exceed an AWS Global Accelerator
	// limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeListenerNotFoundException for service response error code
	// "ListenerNotFoundException".
	//
	// The listener that you specified doesn't exist.
	ErrCodeListenerNotFoundException = "ListenerNotFoundException"
)
