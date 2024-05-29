// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the password policy settings for the Amazon Web Services account. This
// operation does not support partial updates. No parameters are required, but if
// you do not specify a parameter, that parameter's value reverts to its default
// value. See the Request Parameters section for each parameter's default value.
// Also note that some parameters do not allow the default parameter to be
// explicitly set. Instead, to invoke the default value, do not include that
// parameter when you invoke the operation. For more information about using a
// password policy, see Managing an IAM password policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html)
// in the IAM User Guide.
func (c *Client) UpdateAccountPasswordPolicy(ctx context.Context, params *UpdateAccountPasswordPolicyInput, optFns ...func(*Options)) (*UpdateAccountPasswordPolicyOutput, error) {
	if params == nil {
		params = &UpdateAccountPasswordPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountPasswordPolicy", params, optFns, c.addOperationUpdateAccountPasswordPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountPasswordPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountPasswordPolicyInput struct {

	// Allows all IAM users in your account to use the Amazon Web Services Management
	// Console to change their own passwords. For more information, see Permitting IAM
	// users to change their own passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_enable-user-change.html)
	// in the IAM User Guide. If you do not specify a value for this parameter, then
	// the operation uses the default value of false . The result is that IAM users in
	// the account do not automatically have permissions to change their own password.
	AllowUsersToChangePassword bool

	// Prevents IAM users who are accessing the account via the Amazon Web Services
	// Management Console from setting a new console password after their password has
	// expired. The IAM user cannot access the console until an administrator resets
	// the password. If you do not specify a value for this parameter, then the
	// operation uses the default value of false . The result is that IAM users can
	// change their passwords after they expire and continue to sign in as the user. In
	// the Amazon Web Services Management Console, the custom password policy option
	// Allow users to change their own password gives IAM users permissions to
	// iam:ChangePassword for only their user and to the iam:GetAccountPasswordPolicy
	// action. This option does not attach a permissions policy to each user, rather
	// the permissions are applied at the account-level for all users by IAM. IAM users
	// with iam:ChangePassword permission and active access keys can reset their own
	// expired console password using the CLI or API.
	HardExpiry *bool

	// The number of days that an IAM user password is valid. If you do not specify a
	// value for this parameter, then the operation uses the default value of 0 . The
	// result is that IAM user passwords never expire.
	MaxPasswordAge *int32

	// The minimum number of characters allowed in an IAM user password. If you do not
	// specify a value for this parameter, then the operation uses the default value of
	// 6 .
	MinimumPasswordLength *int32

	// Specifies the number of previous passwords that IAM users are prevented from
	// reusing. If you do not specify a value for this parameter, then the operation
	// uses the default value of 0 . The result is that IAM users are not prevented
	// from reusing previous passwords.
	PasswordReusePrevention *int32

	// Specifies whether IAM user passwords must contain at least one lowercase
	// character from the ISO basic Latin alphabet (a to z). If you do not specify a
	// value for this parameter, then the operation uses the default value of false .
	// The result is that passwords do not require at least one lowercase character.
	RequireLowercaseCharacters bool

	// Specifies whether IAM user passwords must contain at least one numeric
	// character (0 to 9). If you do not specify a value for this parameter, then the
	// operation uses the default value of false . The result is that passwords do not
	// require at least one numeric character.
	RequireNumbers bool

	// Specifies whether IAM user passwords must contain at least one of the following
	// non-alphanumeric characters: ! @ # $ % ^ & * ( ) _ + - = [ ] { } | ' If you do
	// not specify a value for this parameter, then the operation uses the default
	// value of false . The result is that passwords do not require at least one symbol
	// character.
	RequireSymbols bool

	// Specifies whether IAM user passwords must contain at least one uppercase
	// character from the ISO basic Latin alphabet (A to Z). If you do not specify a
	// value for this parameter, then the operation uses the default value of false .
	// The result is that passwords do not require at least one uppercase character.
	RequireUppercaseCharacters bool

	noSmithyDocumentSerde
}

type UpdateAccountPasswordPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccountPasswordPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAccountPasswordPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAccountPasswordPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateAccountPasswordPolicyResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateAccountPasswordPolicy",
	}
}

type opUpdateAccountPasswordPolicyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateAccountPasswordPolicyResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateAccountPasswordPolicyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "iam"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "iam"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("iam")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateAccountPasswordPolicyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateAccountPasswordPolicyResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
