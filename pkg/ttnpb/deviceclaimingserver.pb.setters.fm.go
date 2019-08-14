// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"

	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
)

func (dst *ClaimEndDeviceRequest) SetFields(src *ClaimEndDeviceRequest, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "target_application_ids":
			if len(subs) > 0 {
				newDst := &dst.TargetApplicationIDs
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.TargetApplicationIDs
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.TargetApplicationIDs = src.TargetApplicationIDs
				} else {
					var zero ApplicationIdentifiers
					dst.TargetApplicationIDs = zero
				}
			}
		case "target_device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetDeviceID = src.TargetDeviceID
			} else {
				var zero string
				dst.TargetDeviceID = zero
			}
		case "invalidate_authentication_code":
			if len(subs) > 0 {
				return fmt.Errorf("'invalidate_authentication_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InvalidateAuthenticationCode = src.InvalidateAuthenticationCode
			} else {
				var zero bool
				dst.InvalidateAuthenticationCode = zero
			}

		case "source_device":
			if len(subs) == 0 && src == nil {
				dst.SourceDevice = nil
				continue
			} else if len(subs) == 0 {
				dst.SourceDevice = src.SourceDevice
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "authenticated_identifiers":
					if _, ok := dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_); !ok {
						dst.SourceDevice = &ClaimEndDeviceRequest_AuthenticatedIdentifiers_{}
					}
					if len(oneofSubs) > 0 {
						newDst := dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers
						if newDst == nil {
							newDst = &ClaimEndDeviceRequest_AuthenticatedIdentifiers{}
							dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers = newDst
						}
						var newSrc *ClaimEndDeviceRequest_AuthenticatedIdentifiers
						if src != nil {
							newSrc = src.GetAuthenticatedIdentifiers()
						}
						if err := newDst.SetFields(newSrc, subs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers = src.GetAuthenticatedIdentifiers()
						} else {
							dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers = nil
						}
					}
				case "qr_code":
					if _, ok := dst.SourceDevice.(*ClaimEndDeviceRequest_QRCode); !ok {
						dst.SourceDevice = &ClaimEndDeviceRequest_QRCode{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'qr_code' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.SourceDevice.(*ClaimEndDeviceRequest_QRCode).QRCode = src.GetQRCode()
					} else {
						dst.SourceDevice.(*ClaimEndDeviceRequest_QRCode).QRCode = nil
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

func (dst *ClaimEndDeviceRequest_AuthenticatedIdentifiers) SetFields(src *ClaimEndDeviceRequest_AuthenticatedIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.JoinEUI = zero
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEUI = src.DevEUI
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.EUI64
				dst.DevEUI = zero
			}
		case "authentication_code":
			if len(subs) > 0 {
				return fmt.Errorf("'authentication_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AuthenticationCode = src.AuthenticationCode
			} else {
				dst.AuthenticationCode = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
