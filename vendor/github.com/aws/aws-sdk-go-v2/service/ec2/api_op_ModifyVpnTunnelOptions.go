// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the options for a VPN tunnel in an Amazon Web Services Site-to-Site VPN
// connection. You can modify multiple options for a tunnel in a single request,
// but you can only modify one tunnel at a time. For more information, see
// Site-to-Site VPN tunnel options for your Site-to-Site VPN connection
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html) in the Amazon
// Web Services Site-to-Site VPN User Guide.
func (c *Client) ModifyVpnTunnelOptions(ctx context.Context, params *ModifyVpnTunnelOptionsInput, optFns ...func(*Options)) (*ModifyVpnTunnelOptionsOutput, error) {
	if params == nil {
		params = &ModifyVpnTunnelOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpnTunnelOptions", params, optFns, c.addOperationModifyVpnTunnelOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpnTunnelOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnTunnelOptionsInput struct {

	// The tunnel options to modify.
	//
	// This member is required.
	TunnelOptions *types.ModifyVpnTunnelOptionsSpecification

	// The ID of the Amazon Web Services Site-to-Site VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	// The external IP address of the VPN tunnel.
	//
	// This member is required.
	VpnTunnelOutsideIpAddress *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyVpnTunnelOptionsOutput struct {

	// Describes a VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpnTunnelOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnTunnelOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnTunnelOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyVpnTunnelOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnTunnelOptions(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opModifyVpnTunnelOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpnTunnelOptions",
	}
}
