// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Specifies when to create snapshots of EBS volumes.
type CreateRule struct {
	_ struct{} `type:"structure"`

	// The interval between snapshots. The supported values are 2, 3, 4, 6, 8, 12,
	// and 24.
	//
	// Interval is a required field
	Interval *int64 `min:"1" type:"integer" required:"true"`

	// The interval unit.
	//
	// IntervalUnit is a required field
	IntervalUnit IntervalUnitValues `type:"string" required:"true" enum:"true"`

	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time.
	Times []string `type:"list"`
}

// String returns the string representation
func (s CreateRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRule"}

	if s.Interval == nil {
		invalidParams.Add(aws.NewErrParamRequired("Interval"))
	}
	if s.Interval != nil && *s.Interval < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 1))
	}
	if len(s.IntervalUnit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IntervalUnit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.Interval != nil {
		v := *s.Interval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Interval", protocol.Int64Value(v), metadata)
	}
	if len(s.IntervalUnit) > 0 {
		v := s.IntervalUnit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IntervalUnit", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Times != nil {
		v := s.Times

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Times", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Specifies the retention rule for cross-Region snapshot copies.
type CrossRegionCopyRetainRule struct {
	_ struct{} `type:"structure"`

	// The amount of time to retain each snapshot. The maximum is 100 years. This
	// is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int64 `min:"1" type:"integer"`

	// The unit of time for time-based retention.
	IntervalUnit RetentionIntervalUnitValues `type:"string" enum:"true"`
}

// String returns the string representation
func (s CrossRegionCopyRetainRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CrossRegionCopyRetainRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CrossRegionCopyRetainRule"}
	if s.Interval != nil && *s.Interval < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CrossRegionCopyRetainRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.Interval != nil {
		v := *s.Interval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Interval", protocol.Int64Value(v), metadata)
	}
	if len(s.IntervalUnit) > 0 {
		v := s.IntervalUnit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IntervalUnit", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Specifies a rule for cross-Region snapshot copies.
type CrossRegionCopyRule struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the AWS KMS customer master key (CMK) to
	// use for EBS encryption. If this parameter is not specified, your AWS managed
	// CMK for EBS is used.
	CmkArn *string `type:"string"`

	// Copy all user-defined tags from the source snapshot to the copied snapshot.
	CopyTags *bool `type:"boolean"`

	// To encrypt a copy of an unencrypted snapshot if encryption by default is
	// not enabled, enable encryption using this parameter. Copies of encrypted
	// snapshots are encrypted, even if this parameter is false or if encryption
	// by default is not enabled.
	//
	// Encrypted is a required field
	Encrypted *bool `type:"boolean" required:"true"`

	// The retention rule.
	RetainRule *CrossRegionCopyRetainRule `type:"structure"`

	// The target Region.
	//
	// TargetRegion is a required field
	TargetRegion *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CrossRegionCopyRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CrossRegionCopyRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CrossRegionCopyRule"}

	if s.Encrypted == nil {
		invalidParams.Add(aws.NewErrParamRequired("Encrypted"))
	}

	if s.TargetRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetRegion"))
	}
	if s.RetainRule != nil {
		if err := s.RetainRule.Validate(); err != nil {
			invalidParams.AddNested("RetainRule", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CrossRegionCopyRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.CmkArn != nil {
		v := *s.CmkArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CmkArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CopyTags != nil {
		v := *s.CopyTags

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CopyTags", protocol.BoolValue(v), metadata)
	}
	if s.Encrypted != nil {
		v := *s.Encrypted

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Encrypted", protocol.BoolValue(v), metadata)
	}
	if s.RetainRule != nil {
		v := s.RetainRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RetainRule", v, metadata)
	}
	if s.TargetRegion != nil {
		v := *s.TargetRegion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetRegion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Specifies a rule for enabling fast snapshot restore. You can enable fast
// snapshot restore based on either a count or a time interval.
type FastRestoreRule struct {
	_ struct{} `type:"structure"`

	// The Availability Zones in which to enable fast snapshot restore.
	//
	// AvailabilityZones is a required field
	AvailabilityZones []string `min:"1" type:"list" required:"true"`

	// The number of snapshots to be enabled with fast snapshot restore.
	Count *int64 `min:"1" type:"integer"`

	// The amount of time to enable fast snapshot restore. The maximum is 100 years.
	// This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int64 `min:"1" type:"integer"`

	// The unit of time for enabling fast snapshot restore.
	IntervalUnit RetentionIntervalUnitValues `type:"string" enum:"true"`
}

// String returns the string representation
func (s FastRestoreRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FastRestoreRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FastRestoreRule"}

	if s.AvailabilityZones == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZones"))
	}
	if s.AvailabilityZones != nil && len(s.AvailabilityZones) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZones", 1))
	}
	if s.Count != nil && *s.Count < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Count", 1))
	}
	if s.Interval != nil && *s.Interval < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FastRestoreRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.AvailabilityZones != nil {
		v := s.AvailabilityZones

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AvailabilityZones", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Count", protocol.Int64Value(v), metadata)
	}
	if s.Interval != nil {
		v := *s.Interval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Interval", protocol.Int64Value(v), metadata)
	}
	if len(s.IntervalUnit) > 0 {
		v := s.IntervalUnit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IntervalUnit", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Detailed information about a lifecycle policy.
type LifecyclePolicy struct {
	_ struct{} `type:"structure"`

	// The local date and time when the lifecycle policy was created.
	DateCreated *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The local date and time when the lifecycle policy was last modified.
	DateModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The description of the lifecycle policy.
	Description *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string `type:"string"`

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string `type:"string"`

	// The configuration of the lifecycle policy
	PolicyDetails *PolicyDetails `type:"structure"`

	// The identifier of the lifecycle policy.
	PolicyId *string `type:"string"`

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues `type:"string" enum:"true"`

	// The description of the status.
	StatusMessage *string `type:"string"`

	// The tags.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s LifecyclePolicy) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LifecyclePolicy) MarshalFields(e protocol.FieldEncoder) error {
	if s.DateCreated != nil {
		v := *s.DateCreated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DateCreated",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.DateModified != nil {
		v := *s.DateModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DateModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExecutionRoleArn != nil {
		v := *s.ExecutionRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExecutionRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyArn != nil {
		v := *s.PolicyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyDetails != nil {
		v := s.PolicyDetails

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PolicyDetails", v, metadata)
	}
	if s.PolicyId != nil {
		v := *s.PolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Summary information about a lifecycle policy.
type LifecyclePolicySummary struct {
	_ struct{} `type:"structure"`

	// The description of the lifecycle policy.
	Description *string `type:"string"`

	// The identifier of the lifecycle policy.
	PolicyId *string `type:"string"`

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues `type:"string" enum:"true"`

	// The tags.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s LifecyclePolicySummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LifecyclePolicySummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyId != nil {
		v := *s.PolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Specifies optional parameters to add to a policy. The set of valid parameters
// depends on the combination of policy type and resource type.
type Parameters struct {
	_ struct{} `type:"structure"`

	// [EBS Snapshot Management – Instance policies only] Indicates whether to
	// exclude the root volume from snapshots created using CreateSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSnapshots.html).
	// The default is false.
	ExcludeBootVolume *bool `type:"boolean"`
}

// String returns the string representation
func (s Parameters) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Parameters) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExcludeBootVolume != nil {
		v := *s.ExcludeBootVolume

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExcludeBootVolume", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Specifies the configuration of a lifecycle policy.
type PolicyDetails struct {
	_ struct{} `type:"structure"`

	// A set of optional parameters for the policy.
	Parameters *Parameters `type:"structure"`

	// The valid target resource types and actions a policy can manage. The default
	// is EBS_SNAPSHOT_MANAGEMENT.
	PolicyType PolicyTypeValues `type:"string" enum:"true"`

	// The resource type.
	ResourceTypes []ResourceTypeValues `min:"1" type:"list"`

	// The schedule of policy-defined actions.
	Schedules []Schedule `min:"1" type:"list"`

	// The single tag that identifies targeted resources for this policy.
	TargetTags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s PolicyDetails) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PolicyDetails) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PolicyDetails"}
	if s.ResourceTypes != nil && len(s.ResourceTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceTypes", 1))
	}
	if s.Schedules != nil && len(s.Schedules) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Schedules", 1))
	}
	if s.TargetTags != nil && len(s.TargetTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetTags", 1))
	}
	if s.Schedules != nil {
		for i, v := range s.Schedules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Schedules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TargetTags != nil {
		for i, v := range s.TargetTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PolicyDetails) MarshalFields(e protocol.FieldEncoder) error {
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Parameters", v, metadata)
	}
	if len(s.PolicyType) > 0 {
		v := s.PolicyType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PolicyType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ResourceTypes != nil {
		v := s.ResourceTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ResourceTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Schedules != nil {
		v := s.Schedules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Schedules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TargetTags != nil {
		v := s.TargetTags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TargetTags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Specifies the retention rule for a lifecycle policy. You can retain snapshots
// based on either a count or a time interval.
type RetainRule struct {
	_ struct{} `type:"structure"`

	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	Count *int64 `min:"1" type:"integer"`

	// The amount of time to retain each snapshot. The maximum is 100 years. This
	// is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int64 `min:"1" type:"integer"`

	// The unit of time for time-based retention.
	IntervalUnit RetentionIntervalUnitValues `type:"string" enum:"true"`
}

// String returns the string representation
func (s RetainRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetainRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetainRule"}
	if s.Count != nil && *s.Count < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Count", 1))
	}
	if s.Interval != nil && *s.Interval < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Interval", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RetainRule) MarshalFields(e protocol.FieldEncoder) error {
	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Count", protocol.Int64Value(v), metadata)
	}
	if s.Interval != nil {
		v := *s.Interval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Interval", protocol.Int64Value(v), metadata)
	}
	if len(s.IntervalUnit) > 0 {
		v := s.IntervalUnit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IntervalUnit", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Specifies a backup schedule.
type Schedule struct {
	_ struct{} `type:"structure"`

	// Copy all user-defined tags on a source volume to snapshots of the volume
	// created by this policy.
	CopyTags *bool `type:"boolean"`

	// The creation rule.
	CreateRule *CreateRule `type:"structure"`

	// The rule for cross-Region snapshot copies.
	CrossRegionCopyRules []CrossRegionCopyRule `type:"list"`

	// The rule for enabling fast snapshot restore.
	FastRestoreRule *FastRestoreRule `type:"structure"`

	// The name of the schedule.
	Name *string `type:"string"`

	// The retention rule.
	RetainRule *RetainRule `type:"structure"`

	// The tags to apply to policy-created resources. These user-defined tags are
	// in addition to the AWS-added lifecycle tags.
	TagsToAdd []Tag `type:"list"`

	// A collection of key/value pairs with values determined dynamically when the
	// policy is executed. Keys may be any valid Amazon EC2 tag key. Values must
	// be in one of the two following formats: $(instance-id) or $(timestamp). Variable
	// tags are only valid for EBS Snapshot Management – Instance policies.
	VariableTags []Tag `type:"list"`
}

// String returns the string representation
func (s Schedule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Schedule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Schedule"}
	if s.CreateRule != nil {
		if err := s.CreateRule.Validate(); err != nil {
			invalidParams.AddNested("CreateRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.CrossRegionCopyRules != nil {
		for i, v := range s.CrossRegionCopyRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CrossRegionCopyRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.FastRestoreRule != nil {
		if err := s.FastRestoreRule.Validate(); err != nil {
			invalidParams.AddNested("FastRestoreRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.RetainRule != nil {
		if err := s.RetainRule.Validate(); err != nil {
			invalidParams.AddNested("RetainRule", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagsToAdd != nil {
		for i, v := range s.TagsToAdd {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagsToAdd", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VariableTags != nil {
		for i, v := range s.VariableTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "VariableTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Schedule) MarshalFields(e protocol.FieldEncoder) error {
	if s.CopyTags != nil {
		v := *s.CopyTags

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CopyTags", protocol.BoolValue(v), metadata)
	}
	if s.CreateRule != nil {
		v := s.CreateRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CreateRule", v, metadata)
	}
	if s.CrossRegionCopyRules != nil {
		v := s.CrossRegionCopyRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CrossRegionCopyRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FastRestoreRule != nil {
		v := s.FastRestoreRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "FastRestoreRule", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RetainRule != nil {
		v := s.RetainRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RetainRule", v, metadata)
	}
	if s.TagsToAdd != nil {
		v := s.TagsToAdd

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TagsToAdd", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VariableTags != nil {
		v := s.VariableTags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "VariableTags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Specifies a tag for a resource.
type Tag struct {
	_ struct{} `type:"structure"`

	// The tag key.
	//
	// Key is a required field
	Key *string `type:"string" required:"true"`

	// The tag value.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Tag) MarshalFields(e protocol.FieldEncoder) error {
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
