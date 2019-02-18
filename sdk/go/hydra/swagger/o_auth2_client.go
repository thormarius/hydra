/*
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"time"
)

type OAuth2Client struct {

	// AllowedCORSOrigins are one or more URLs (scheme://host[:port]) which are allowed to make CORS requests to the /oauth/token endpoint. If this array is empty, the sever's CORS origin configuration (`CORS_ALLOWED_ORIGINS`) will be used instead. If this array is set, the allowed origins are appended to the server's CORS origin configuration. Be aware that environment variable `CORS_ENABLED` MUST be set to `true` for this to work.
	AllowedCorsOrigins []string `json:"allowed_cors_origins,omitempty"`

	// Audience is a whitelist defining the audiences this client is allowed to request tokens for. An audience limits the applicability of an OAuth 2.0 Access Token to, for example, certain API endpoints. The value is a list of URLs. URLs MUST NOT contain whitespaces.
	Audience []string `json:"audience,omitempty"`

	// ClientID  is the id for this client.
	ClientId string `json:"client_id,omitempty"`

	// Name is the human-readable string name of the client to be presented to the end-user during authorization.
	ClientName string `json:"client_name,omitempty"`

	// Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
	ClientSecret string `json:"client_secret,omitempty"`

	// SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire. The time is represented as the number of seconds from 1970-01-01T00:00:00Z as measured in UTC until the date/time of expiration.
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at,omitempty"`

	// ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
	ClientUri string `json:"client_uri,omitempty"`

	// Contacts is a array of strings representing ways to contact people responsible for this client, typically email addresses.
	Contacts []string `json:"contacts,omitempty"`

	// CreatedAt returns the timestamp of the client's creation.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// GrantTypes is an array of grant types the client is allowed to use.
	GrantTypes []string `json:"grant_types,omitempty"`

	Jwks JsonWebKeySet `json:"jwks,omitempty"`

	// URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
	JwksUri string `json:"jwks_uri,omitempty"`

	// LogoURI is an URL string that references a logo for the client.
	LogoUri string `json:"logo_uri,omitempty"`

	// Owner is a string identifying the owner of the OAuth 2.0 Client.
	Owner string `json:"owner,omitempty"`

	// PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
	PolicyUri string `json:"policy_uri,omitempty"`

	// RedirectURIs is an array of allowed redirect urls for the client, for example http://mydomain/oauth/callback .
	RedirectUris []string `json:"redirect_uris,omitempty"`

	// JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.
	RequestObjectSigningAlg string `json:"request_object_signing_alg,omitempty"`

	// Array of request_uri values that are pre-registered by the RP for use at the OP. Servers MAY cache the contents of the files referenced by these URIs and not retrieve them at the time they are used in a request. OPs can require that request_uri values used be pre-registered with the require_request_uri_registration discovery parameter.
	RequestUris []string `json:"request_uris,omitempty"`

	// ResponseTypes is an array of the OAuth 2.0 response type strings that the client can use at the authorization endpoint.
	ResponseTypes []string `json:"response_types,omitempty"`

	// Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
	Scope string `json:"scope,omitempty"`

	// URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.
	SectorIdentifierUri string `json:"sector_identifier_uri,omitempty"`

	// SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
	SubjectType string `json:"subject_type,omitempty"`

	// Requested Client Authentication method for the Token Endpoint. The options are client_secret_post, client_secret_basic, private_key_jwt, and none.
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"`

	// TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
	TosUri string `json:"tos_uri,omitempty"`

	// UpdatedAt returns the timestamp of the last update.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.
	UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg,omitempty"`
}
