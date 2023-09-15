// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *AuthInfoResponse) SetFields(src *AuthInfoResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "universal_rights":
			if len(subs) > 0 {
				var newDst, newSrc *Rights
				if (src == nil || src.UniversalRights == nil) && dst.UniversalRights == nil {
					continue
				}
				if src != nil {
					newSrc = src.UniversalRights
				}
				if dst.UniversalRights != nil {
					newDst = dst.UniversalRights
				} else {
					newDst = &Rights{}
					dst.UniversalRights = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UniversalRights = src.UniversalRights
				} else {
					dst.UniversalRights = nil
				}
			}
		case "is_admin":
			if len(subs) > 0 {
				return fmt.Errorf("'is_admin' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.IsAdmin = src.IsAdmin
			} else {
				var zero bool
				dst.IsAdmin = zero
			}

		case "access_method":
			if len(subs) == 0 && src == nil {
				dst.AccessMethod = nil
				continue
			} else if len(subs) == 0 {
				dst.AccessMethod = src.AccessMethod
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "api_key":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.AccessMethod.(*AuthInfoResponse_ApiKey)
					}
					if srcValid := srcTypeOk || src == nil || src.AccessMethod == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'api_key', while different oneof is set in source")
					}
					_, dstTypeOk := dst.AccessMethod.(*AuthInfoResponse_ApiKey)
					if dstValid := dstTypeOk || dst.AccessMethod == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'api_key', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *AuthInfoResponse_APIKeyAccess
						if srcTypeOk {
							newSrc = src.AccessMethod.(*AuthInfoResponse_ApiKey).ApiKey
						}
						if dstTypeOk {
							newDst = dst.AccessMethod.(*AuthInfoResponse_ApiKey).ApiKey
						} else if srcTypeOk {
							newDst = &AuthInfoResponse_APIKeyAccess{}
							dst.AccessMethod = &AuthInfoResponse_ApiKey{ApiKey: newDst}
						} else {
							dst.AccessMethod = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.AccessMethod = src.AccessMethod
						} else {
							dst.AccessMethod = nil
						}
					}
				case "oauth_access_token":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.AccessMethod.(*AuthInfoResponse_OauthAccessToken)
					}
					if srcValid := srcTypeOk || src == nil || src.AccessMethod == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'oauth_access_token', while different oneof is set in source")
					}
					_, dstTypeOk := dst.AccessMethod.(*AuthInfoResponse_OauthAccessToken)
					if dstValid := dstTypeOk || dst.AccessMethod == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'oauth_access_token', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OAuthAccessToken
						if srcTypeOk {
							newSrc = src.AccessMethod.(*AuthInfoResponse_OauthAccessToken).OauthAccessToken
						}
						if dstTypeOk {
							newDst = dst.AccessMethod.(*AuthInfoResponse_OauthAccessToken).OauthAccessToken
						} else if srcTypeOk {
							newDst = &OAuthAccessToken{}
							dst.AccessMethod = &AuthInfoResponse_OauthAccessToken{OauthAccessToken: newDst}
						} else {
							dst.AccessMethod = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.AccessMethod = src.AccessMethod
						} else {
							dst.AccessMethod = nil
						}
					}
				case "user_session":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.AccessMethod.(*AuthInfoResponse_UserSession)
					}
					if srcValid := srcTypeOk || src == nil || src.AccessMethod == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'user_session', while different oneof is set in source")
					}
					_, dstTypeOk := dst.AccessMethod.(*AuthInfoResponse_UserSession)
					if dstValid := dstTypeOk || dst.AccessMethod == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'user_session', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserSession
						if srcTypeOk {
							newSrc = src.AccessMethod.(*AuthInfoResponse_UserSession).UserSession
						}
						if dstTypeOk {
							newDst = dst.AccessMethod.(*AuthInfoResponse_UserSession).UserSession
						} else if srcTypeOk {
							newDst = &UserSession{}
							dst.AccessMethod = &AuthInfoResponse_UserSession{UserSession: newDst}
						} else {
							dst.AccessMethod = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.AccessMethod = src.AccessMethod
						} else {
							dst.AccessMethod = nil
						}
					}
				case "gateway_token":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.AccessMethod.(*AuthInfoResponse_GatewayToken_)
					}
					if srcValid := srcTypeOk || src == nil || src.AccessMethod == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'gateway_token', while different oneof is set in source")
					}
					_, dstTypeOk := dst.AccessMethod.(*AuthInfoResponse_GatewayToken_)
					if dstValid := dstTypeOk || dst.AccessMethod == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'gateway_token', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *AuthInfoResponse_GatewayToken
						if srcTypeOk {
							newSrc = src.AccessMethod.(*AuthInfoResponse_GatewayToken_).GatewayToken
						}
						if dstTypeOk {
							newDst = dst.AccessMethod.(*AuthInfoResponse_GatewayToken_).GatewayToken
						} else if srcTypeOk {
							newDst = &AuthInfoResponse_GatewayToken{}
							dst.AccessMethod = &AuthInfoResponse_GatewayToken_{GatewayToken: newDst}
						} else {
							dst.AccessMethod = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.AccessMethod = src.AccessMethod
						} else {
							dst.AccessMethod = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetIsConfigurationRequest) SetFields(src *GetIsConfigurationRequest, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message GetIsConfigurationRequest has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *IsConfiguration) SetFields(src *IsConfiguration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_registration":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRegistration
				if (src == nil || src.UserRegistration == nil) && dst.UserRegistration == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserRegistration
				}
				if dst.UserRegistration != nil {
					newDst = dst.UserRegistration
				} else {
					newDst = &IsConfiguration_UserRegistration{}
					dst.UserRegistration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserRegistration = src.UserRegistration
				} else {
					dst.UserRegistration = nil
				}
			}
		case "profile_picture":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_ProfilePicture
				if (src == nil || src.ProfilePicture == nil) && dst.ProfilePicture == nil {
					continue
				}
				if src != nil {
					newSrc = src.ProfilePicture
				}
				if dst.ProfilePicture != nil {
					newDst = dst.ProfilePicture
				} else {
					newDst = &IsConfiguration_ProfilePicture{}
					dst.ProfilePicture = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ProfilePicture = src.ProfilePicture
				} else {
					dst.ProfilePicture = nil
				}
			}
		case "end_device_picture":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_EndDevicePicture
				if (src == nil || src.EndDevicePicture == nil) && dst.EndDevicePicture == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDevicePicture
				}
				if dst.EndDevicePicture != nil {
					newDst = dst.EndDevicePicture
				} else {
					newDst = &IsConfiguration_EndDevicePicture{}
					dst.EndDevicePicture = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDevicePicture = src.EndDevicePicture
				} else {
					dst.EndDevicePicture = nil
				}
			}
		case "user_rights":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRights
				if (src == nil || src.UserRights == nil) && dst.UserRights == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserRights
				}
				if dst.UserRights != nil {
					newDst = dst.UserRights
				} else {
					newDst = &IsConfiguration_UserRights{}
					dst.UserRights = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserRights = src.UserRights
				} else {
					dst.UserRights = nil
				}
			}
		case "user_login":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserLogin
				if (src == nil || src.UserLogin == nil) && dst.UserLogin == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserLogin
				}
				if dst.UserLogin != nil {
					newDst = dst.UserLogin
				} else {
					newDst = &IsConfiguration_UserLogin{}
					dst.UserLogin = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserLogin = src.UserLogin
				} else {
					dst.UserLogin = nil
				}
			}
		case "admin_rights":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_AdminRights
				if (src == nil || src.AdminRights == nil) && dst.AdminRights == nil {
					continue
				}
				if src != nil {
					newSrc = src.AdminRights
				}
				if dst.AdminRights != nil {
					newDst = dst.AdminRights
				} else {
					newDst = &IsConfiguration_AdminRights{}
					dst.AdminRights = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AdminRights = src.AdminRights
				} else {
					dst.AdminRights = nil
				}
			}
		case "collaborator_rights":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_CollaboratorRights
				if (src == nil || src.CollaboratorRights == nil) && dst.CollaboratorRights == nil {
					continue
				}
				if src != nil {
					newSrc = src.CollaboratorRights
				}
				if dst.CollaboratorRights != nil {
					newDst = dst.CollaboratorRights
				} else {
					newDst = &IsConfiguration_CollaboratorRights{}
					dst.CollaboratorRights = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.CollaboratorRights = src.CollaboratorRights
				} else {
					dst.CollaboratorRights = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetIsConfigurationResponse) SetFields(src *GetIsConfigurationResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &IsConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AuthInfoResponse_APIKeyAccess) SetFields(src *AuthInfoResponse_APIKeyAccess, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "api_key":
			if len(subs) > 0 {
				var newDst, newSrc *APIKey
				if (src == nil || src.ApiKey == nil) && dst.ApiKey == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApiKey
				}
				if dst.ApiKey != nil {
					newDst = dst.ApiKey
				} else {
					newDst = &APIKey{}
					dst.ApiKey = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApiKey = src.ApiKey
				} else {
					dst.ApiKey = nil
				}
			}
		case "entity_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EntityIdentifiers
				if (src == nil || src.EntityIds == nil) && dst.EntityIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EntityIds
				}
				if dst.EntityIds != nil {
					newDst = dst.EntityIds
				} else {
					newDst = &EntityIdentifiers{}
					dst.EntityIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EntityIds = src.EntityIds
				} else {
					dst.EntityIds = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AuthInfoResponse_GatewayToken) SetFields(src *AuthInfoResponse_GatewayToken, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				var newDst, newSrc *GatewayIdentifiers
				if (src == nil || src.GatewayIds == nil) && dst.GatewayIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.GatewayIds
				}
				if dst.GatewayIds != nil {
					newDst = dst.GatewayIds
				} else {
					newDst = &GatewayIdentifiers{}
					dst.GatewayIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.GatewayIds = src.GatewayIds
				} else {
					dst.GatewayIds = nil
				}
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRegistration) SetFields(src *IsConfiguration_UserRegistration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "invitation":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRegistration_Invitation
				if (src == nil || src.Invitation == nil) && dst.Invitation == nil {
					continue
				}
				if src != nil {
					newSrc = src.Invitation
				}
				if dst.Invitation != nil {
					newDst = dst.Invitation
				} else {
					newDst = &IsConfiguration_UserRegistration_Invitation{}
					dst.Invitation = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Invitation = src.Invitation
				} else {
					dst.Invitation = nil
				}
			}
		case "contact_info_validation":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRegistration_ContactInfoValidation
				if (src == nil || src.ContactInfoValidation == nil) && dst.ContactInfoValidation == nil {
					continue
				}
				if src != nil {
					newSrc = src.ContactInfoValidation
				}
				if dst.ContactInfoValidation != nil {
					newDst = dst.ContactInfoValidation
				} else {
					newDst = &IsConfiguration_UserRegistration_ContactInfoValidation{}
					dst.ContactInfoValidation = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ContactInfoValidation = src.ContactInfoValidation
				} else {
					dst.ContactInfoValidation = nil
				}
			}
		case "admin_approval":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRegistration_AdminApproval
				if (src == nil || src.AdminApproval == nil) && dst.AdminApproval == nil {
					continue
				}
				if src != nil {
					newSrc = src.AdminApproval
				}
				if dst.AdminApproval != nil {
					newDst = dst.AdminApproval
				} else {
					newDst = &IsConfiguration_UserRegistration_AdminApproval{}
					dst.AdminApproval = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AdminApproval = src.AdminApproval
				} else {
					dst.AdminApproval = nil
				}
			}
		case "password_requirements":
			if len(subs) > 0 {
				var newDst, newSrc *IsConfiguration_UserRegistration_PasswordRequirements
				if (src == nil || src.PasswordRequirements == nil) && dst.PasswordRequirements == nil {
					continue
				}
				if src != nil {
					newSrc = src.PasswordRequirements
				}
				if dst.PasswordRequirements != nil {
					newDst = dst.PasswordRequirements
				} else {
					newDst = &IsConfiguration_UserRegistration_PasswordRequirements{}
					dst.PasswordRequirements = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.PasswordRequirements = src.PasswordRequirements
				} else {
					dst.PasswordRequirements = nil
				}
			}
		case "enabled":
			if len(subs) > 0 {
				return fmt.Errorf("'enabled' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Enabled = src.Enabled
			} else {
				var zero bool
				dst.Enabled = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_ProfilePicture) SetFields(src *IsConfiguration_ProfilePicture, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "disable_upload":
			if len(subs) > 0 {
				return fmt.Errorf("'disable_upload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DisableUpload = src.DisableUpload
			} else {
				dst.DisableUpload = nil
			}
		case "use_gravatar":
			if len(subs) > 0 {
				return fmt.Errorf("'use_gravatar' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UseGravatar = src.UseGravatar
			} else {
				dst.UseGravatar = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_EndDevicePicture) SetFields(src *IsConfiguration_EndDevicePicture, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "disable_upload":
			if len(subs) > 0 {
				return fmt.Errorf("'disable_upload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DisableUpload = src.DisableUpload
			} else {
				dst.DisableUpload = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRights) SetFields(src *IsConfiguration_UserRights, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "create_applications":
			if len(subs) > 0 {
				return fmt.Errorf("'create_applications' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreateApplications = src.CreateApplications
			} else {
				dst.CreateApplications = nil
			}
		case "create_clients":
			if len(subs) > 0 {
				return fmt.Errorf("'create_clients' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreateClients = src.CreateClients
			} else {
				dst.CreateClients = nil
			}
		case "create_gateways":
			if len(subs) > 0 {
				return fmt.Errorf("'create_gateways' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreateGateways = src.CreateGateways
			} else {
				dst.CreateGateways = nil
			}
		case "create_organizations":
			if len(subs) > 0 {
				return fmt.Errorf("'create_organizations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreateOrganizations = src.CreateOrganizations
			} else {
				dst.CreateOrganizations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserLogin) SetFields(src *IsConfiguration_UserLogin, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "disable_credentials_login":
			if len(subs) > 0 {
				return fmt.Errorf("'disable_credentials_login' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DisableCredentialsLogin = src.DisableCredentialsLogin
			} else {
				dst.DisableCredentialsLogin = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_AdminRights) SetFields(src *IsConfiguration_AdminRights, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "all":
			if len(subs) > 0 {
				return fmt.Errorf("'all' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.All = src.All
			} else {
				dst.All = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_CollaboratorRights) SetFields(src *IsConfiguration_CollaboratorRights, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "set_others_as_contacts":
			if len(subs) > 0 {
				return fmt.Errorf("'set_others_as_contacts' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SetOthersAsContacts = src.SetOthersAsContacts
			} else {
				dst.SetOthersAsContacts = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRegistration_Invitation) SetFields(src *IsConfiguration_UserRegistration_Invitation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}
		case "token_ttl":
			if len(subs) > 0 {
				return fmt.Errorf("'token_ttl' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TokenTtl = src.TokenTtl
			} else {
				dst.TokenTtl = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRegistration_ContactInfoValidation) SetFields(src *IsConfiguration_UserRegistration_ContactInfoValidation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}
		case "token_ttl":
			if len(subs) > 0 {
				return fmt.Errorf("'token_ttl' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TokenTtl = src.TokenTtl
			} else {
				dst.TokenTtl = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRegistration_AdminApproval) SetFields(src *IsConfiguration_UserRegistration_AdminApproval, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *IsConfiguration_UserRegistration_PasswordRequirements) SetFields(src *IsConfiguration_UserRegistration_PasswordRequirements, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "min_length":
			if len(subs) > 0 {
				return fmt.Errorf("'min_length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinLength = src.MinLength
			} else {
				dst.MinLength = nil
			}
		case "max_length":
			if len(subs) > 0 {
				return fmt.Errorf("'max_length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxLength = src.MaxLength
			} else {
				dst.MaxLength = nil
			}
		case "min_uppercase":
			if len(subs) > 0 {
				return fmt.Errorf("'min_uppercase' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinUppercase = src.MinUppercase
			} else {
				dst.MinUppercase = nil
			}
		case "min_digits":
			if len(subs) > 0 {
				return fmt.Errorf("'min_digits' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinDigits = src.MinDigits
			} else {
				dst.MinDigits = nil
			}
		case "min_special":
			if len(subs) > 0 {
				return fmt.Errorf("'min_special' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinSpecial = src.MinSpecial
			} else {
				dst.MinSpecial = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
