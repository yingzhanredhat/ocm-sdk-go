// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies a Capacity Reservation Fleet. When you modify the total target
// capacity of a Capacity Reservation Fleet, the Fleet automatically creates new
// Capacity Reservations, or modifies or cancels existing Capacity Reservations in
// the Fleet to meet the new total target capacity. When you modify the end date
// for the Fleet, the end dates for all of the individual Capacity Reservations in
// the Fleet are updated accordingly.
func (c *Client) ModifyCapacityReservationFleet(ctx context.Context, params *ModifyCapacityReservationFleetInput, optFns ...func(*Options)) (*ModifyCapacityReservationFleetOutput, error) {
	if params == nil {
		params = &ModifyCapacityReservationFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCapacityReservationFleet", params, optFns, c.addOperationModifyCapacityReservationFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCapacityReservationFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCapacityReservationFleetInput struct {

	// The ID of the Capacity Reservation Fleet to modify.
	//
	// This member is required.
	CapacityReservationFleetId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The date and time at which the Capacity Reservation Fleet expires. When the
	// Capacity Reservation Fleet expires, its state changes to expired and all of the
	// Capacity Reservations in the Fleet expire. The Capacity Reservation Fleet
	// expires within an hour after the specified time. For example, if you specify
	// 5/31/2019 , 13:30:55 , the Capacity Reservation Fleet is guaranteed to expire
	// between 13:30:55 and 14:30:55 on 5/31/2019 . You can't specify EndDate and
	// RemoveEndDate in the same request.
	EndDate *time.Time

	// Indicates whether to remove the end date from the Capacity Reservation Fleet.
	// If you remove the end date, the Capacity Reservation Fleet does not expire and
	// it remains active until you explicitly cancel it using the
	// CancelCapacityReservationFleet action. You can't specify RemoveEndDate and
	// EndDate in the same request.
	RemoveEndDate *bool

	// The total number of capacity units to be reserved by the Capacity Reservation
	// Fleet. This value, together with the instance type weights that you assign to
	// each instance type used by the Fleet determine the number of instances for which
	// the Fleet reserves capacity. Both values are based on units that make sense for
	// your workload. For more information, see Total target capacity (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#target-capacity)
	// in the Amazon EC2 User Guide.
	TotalTargetCapacity *int32

	noSmithyDocumentSerde
}

type ModifyCapacityReservationFleetOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCapacityReservationFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyCapacityReservationFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyCapacityReservationFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyCapacityReservationFleet"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyCapacityReservationFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCapacityReservationFleet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyCapacityReservationFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyCapacityReservationFleet",
	}
}
