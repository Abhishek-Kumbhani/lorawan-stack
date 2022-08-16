// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ListFrequencyPlansRequest) SetFields(src *ListFrequencyPlansRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "base_frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'base_frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseFrequency = src.BaseFrequency
			} else {
				var zero uint32
				dst.BaseFrequency = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *FrequencyPlanDescription) SetFields(src *FrequencyPlanDescription, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}
		case "base_id":
			if len(subs) > 0 {
				return fmt.Errorf("'base_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseId = src.BaseId
			} else {
				var zero string
				dst.BaseId = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "base_frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'base_frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseFrequency = src.BaseFrequency
			} else {
				var zero uint32
				dst.BaseFrequency = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListFrequencyPlansResponse) SetFields(src *ListFrequencyPlansResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "frequency_plans":
			if len(subs) > 0 {
				return fmt.Errorf("'frequency_plans' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FrequencyPlans = src.FrequencyPlans
			} else {
				dst.FrequencyPlans = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetPhyVersionsRequest) SetFields(src *GetPhyVersionsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "band_id":
			if len(subs) > 0 {
				return fmt.Errorf("'band_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BandId = src.BandId
			} else {
				var zero string
				dst.BandId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetPhyVersionsResponse) SetFields(src *GetPhyVersionsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "version_info":
			if len(subs) > 0 {
				return fmt.Errorf("'version_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.VersionInfo = src.VersionInfo
			} else {
				dst.VersionInfo = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListBandsRequest) SetFields(src *ListBandsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "band_id":
			if len(subs) > 0 {
				return fmt.Errorf("'band_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BandId = src.BandId
			} else {
				var zero string
				dst.BandId = zero
			}
		case "phy_version":
			if len(subs) > 0 {
				return fmt.Errorf("'phy_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PhyVersion = src.PhyVersion
			} else {
				var zero PHYVersion
				dst.PhyVersion = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription) SetFields(src *BandDescription, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}
		case "beacon":
			if len(subs) > 0 {
				var newDst, newSrc *BandDescription_Beacon
				if (src == nil || src.Beacon == nil) && dst.Beacon == nil {
					continue
				}
				if src != nil {
					newSrc = src.Beacon
				}
				if dst.Beacon != nil {
					newDst = dst.Beacon
				} else {
					newDst = &BandDescription_Beacon{}
					dst.Beacon = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Beacon = src.Beacon
				} else {
					dst.Beacon = nil
				}
			}
		case "ping_slot_frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'ping_slot_frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PingSlotFrequency = src.PingSlotFrequency
			} else {
				dst.PingSlotFrequency = nil
			}
		case "max_uplink_channels":
			if len(subs) > 0 {
				return fmt.Errorf("'max_uplink_channels' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxUplinkChannels = src.MaxUplinkChannels
			} else {
				var zero uint32
				dst.MaxUplinkChannels = zero
			}
		case "uplink_channels":
			if len(subs) > 0 {
				return fmt.Errorf("'uplink_channels' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UplinkChannels = src.UplinkChannels
			} else {
				dst.UplinkChannels = nil
			}
		case "max_downlink_channels":
			if len(subs) > 0 {
				return fmt.Errorf("'max_downlink_channels' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxDownlinkChannels = src.MaxDownlinkChannels
			} else {
				var zero uint32
				dst.MaxDownlinkChannels = zero
			}
		case "downlink_channels":
			if len(subs) > 0 {
				return fmt.Errorf("'downlink_channels' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DownlinkChannels = src.DownlinkChannels
			} else {
				dst.DownlinkChannels = nil
			}
		case "sub_bands":
			if len(subs) > 0 {
				return fmt.Errorf("'sub_bands' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SubBands = src.SubBands
			} else {
				dst.SubBands = nil
			}
		case "data_rates":
			if len(subs) > 0 {
				return fmt.Errorf("'data_rates' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DataRates = src.DataRates
			} else {
				dst.DataRates = nil
			}
		case "freq_multiplier":
			if len(subs) > 0 {
				return fmt.Errorf("'freq_multiplier' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FreqMultiplier = src.FreqMultiplier
			} else {
				var zero uint64
				dst.FreqMultiplier = zero
			}
		case "implements_cf_list":
			if len(subs) > 0 {
				return fmt.Errorf("'implements_cf_list' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ImplementsCfList = src.ImplementsCfList
			} else {
				var zero bool
				dst.ImplementsCfList = zero
			}
		case "cf_list_type":
			if len(subs) > 0 {
				return fmt.Errorf("'cf_list_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CfListType = src.CfListType
			} else {
				var zero CFListType
				dst.CfListType = zero
			}
		case "receive_delay_1":
			if len(subs) > 0 {
				return fmt.Errorf("'receive_delay_1' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ReceiveDelay_1 = src.ReceiveDelay_1
			} else {
				dst.ReceiveDelay_1 = nil
			}
		case "receive_delay_2":
			if len(subs) > 0 {
				return fmt.Errorf("'receive_delay_2' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ReceiveDelay_2 = src.ReceiveDelay_2
			} else {
				dst.ReceiveDelay_2 = nil
			}
		case "join_accept_delay_1":
			if len(subs) > 0 {
				return fmt.Errorf("'join_accept_delay_1' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinAcceptDelay_1 = src.JoinAcceptDelay_1
			} else {
				dst.JoinAcceptDelay_1 = nil
			}
		case "join_accept_delay_2":
			if len(subs) > 0 {
				return fmt.Errorf("'join_accept_delay_2' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinAcceptDelay_2 = src.JoinAcceptDelay_2
			} else {
				dst.JoinAcceptDelay_2 = nil
			}
		case "max_fcnt_gap":
			if len(subs) > 0 {
				return fmt.Errorf("'max_fcnt_gap' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxFcntGap = src.MaxFcntGap
			} else {
				var zero uint64
				dst.MaxFcntGap = zero
			}
		case "supports_dynamic_adr":
			if len(subs) > 0 {
				return fmt.Errorf("'supports_dynamic_adr' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SupportsDynamicAdr = src.SupportsDynamicAdr
			} else {
				var zero bool
				dst.SupportsDynamicAdr = zero
			}
		case "adr_ack_limit":
			if len(subs) > 0 {
				return fmt.Errorf("'adr_ack_limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AdrAckLimit = src.AdrAckLimit
			} else {
				var zero ADRAckLimitExponent
				dst.AdrAckLimit = zero
			}
		case "min_retransmit_timeout":
			if len(subs) > 0 {
				return fmt.Errorf("'min_retransmit_timeout' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinRetransmitTimeout = src.MinRetransmitTimeout
			} else {
				dst.MinRetransmitTimeout = nil
			}
		case "max_retransmit_timeout":
			if len(subs) > 0 {
				return fmt.Errorf("'max_retransmit_timeout' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxRetransmitTimeout = src.MaxRetransmitTimeout
			} else {
				dst.MaxRetransmitTimeout = nil
			}
		case "tx_offset":
			if len(subs) > 0 {
				return fmt.Errorf("'tx_offset' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TxOffset = src.TxOffset
			} else {
				dst.TxOffset = nil
			}
		case "max_adr_data_rate_index":
			if len(subs) > 0 {
				return fmt.Errorf("'max_adr_data_rate_index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxAdrDataRateIndex = src.MaxAdrDataRateIndex
			} else {
				var zero DataRateIndex
				dst.MaxAdrDataRateIndex = zero
			}
		case "tx_param_setup_req_support":
			if len(subs) > 0 {
				return fmt.Errorf("'tx_param_setup_req_support' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TxParamSetupReqSupport = src.TxParamSetupReqSupport
			} else {
				var zero bool
				dst.TxParamSetupReqSupport = zero
			}
		case "default_max_eirp":
			if len(subs) > 0 {
				return fmt.Errorf("'default_max_eirp' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DefaultMaxEirp = src.DefaultMaxEirp
			} else {
				var zero float32
				dst.DefaultMaxEirp = zero
			}
		case "lora_coding_rate":
			if len(subs) > 0 {
				return fmt.Errorf("'lora_coding_rate' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LoraCodingRate = src.LoraCodingRate
			} else {
				var zero string
				dst.LoraCodingRate = zero
			}
		case "default_rx2_parameters":
			if len(subs) > 0 {
				var newDst, newSrc *BandDescription_Rx2Parameters
				if (src == nil || src.DefaultRx2Parameters == nil) && dst.DefaultRx2Parameters == nil {
					continue
				}
				if src != nil {
					newSrc = src.DefaultRx2Parameters
				}
				if dst.DefaultRx2Parameters != nil {
					newDst = dst.DefaultRx2Parameters
				} else {
					newDst = &BandDescription_Rx2Parameters{}
					dst.DefaultRx2Parameters = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DefaultRx2Parameters = src.DefaultRx2Parameters
				} else {
					dst.DefaultRx2Parameters = nil
				}
			}
		case "boot_dwell_time":
			if len(subs) > 0 {
				var newDst, newSrc *BandDescription_DwellTime
				if (src == nil || src.BootDwellTime == nil) && dst.BootDwellTime == nil {
					continue
				}
				if src != nil {
					newSrc = src.BootDwellTime
				}
				if dst.BootDwellTime != nil {
					newDst = dst.BootDwellTime
				} else {
					newDst = &BandDescription_DwellTime{}
					dst.BootDwellTime = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.BootDwellTime = src.BootDwellTime
				} else {
					dst.BootDwellTime = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListBandsResponse) SetFields(src *ListBandsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "descriptions":
			if len(subs) > 0 {
				return fmt.Errorf("'descriptions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Descriptions = src.Descriptions
			} else {
				dst.Descriptions = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetPhyVersionsResponse_VersionInfo) SetFields(src *GetPhyVersionsResponse_VersionInfo, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "band_id":
			if len(subs) > 0 {
				return fmt.Errorf("'band_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BandId = src.BandId
			} else {
				var zero string
				dst.BandId = zero
			}
		case "phy_versions":
			if len(subs) > 0 {
				return fmt.Errorf("'phy_versions' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PhyVersions = src.PhyVersions
			} else {
				dst.PhyVersions = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_Beacon) SetFields(src *BandDescription_Beacon, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "data_rate_index":
			if len(subs) > 0 {
				return fmt.Errorf("'data_rate_index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DataRateIndex = src.DataRateIndex
			} else {
				var zero DataRateIndex
				dst.DataRateIndex = zero
			}
		case "coding_rate":
			if len(subs) > 0 {
				return fmt.Errorf("'coding_rate' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CodingRate = src.CodingRate
			} else {
				var zero string
				dst.CodingRate = zero
			}
		case "inverted_polarity":
			if len(subs) > 0 {
				return fmt.Errorf("'inverted_polarity' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InvertedPolarity = src.InvertedPolarity
			} else {
				var zero bool
				dst.InvertedPolarity = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_Channel) SetFields(src *BandDescription_Channel, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Frequency = src.Frequency
			} else {
				var zero uint64
				dst.Frequency = zero
			}
		case "min_data_rate":
			if len(subs) > 0 {
				return fmt.Errorf("'min_data_rate' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinDataRate = src.MinDataRate
			} else {
				var zero DataRateIndex
				dst.MinDataRate = zero
			}
		case "max_data_rate":
			if len(subs) > 0 {
				return fmt.Errorf("'max_data_rate' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxDataRate = src.MaxDataRate
			} else {
				var zero DataRateIndex
				dst.MaxDataRate = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_SubBandParameters) SetFields(src *BandDescription_SubBandParameters, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "min_frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'min_frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinFrequency = src.MinFrequency
			} else {
				var zero uint64
				dst.MinFrequency = zero
			}
		case "max_frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'max_frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxFrequency = src.MaxFrequency
			} else {
				var zero uint64
				dst.MaxFrequency = zero
			}
		case "duty_cycle":
			if len(subs) > 0 {
				return fmt.Errorf("'duty_cycle' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DutyCycle = src.DutyCycle
			} else {
				var zero float32
				dst.DutyCycle = zero
			}
		case "max_eirp":
			if len(subs) > 0 {
				return fmt.Errorf("'max_eirp' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxEirp = src.MaxEirp
			} else {
				var zero float32
				dst.MaxEirp = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_BandDataRate) SetFields(src *BandDescription_BandDataRate, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "rate":
			if len(subs) > 0 {
				var newDst, newSrc *DataRate
				if (src == nil || src.Rate == nil) && dst.Rate == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rate
				}
				if dst.Rate != nil {
					newDst = dst.Rate
				} else {
					newDst = &DataRate{}
					dst.Rate = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rate = src.Rate
				} else {
					dst.Rate = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_Rx2Parameters) SetFields(src *BandDescription_Rx2Parameters, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "data_rate_index":
			if len(subs) > 0 {
				return fmt.Errorf("'data_rate_index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DataRateIndex = src.DataRateIndex
			} else {
				var zero DataRateIndex
				dst.DataRateIndex = zero
			}
		case "frequency":
			if len(subs) > 0 {
				return fmt.Errorf("'frequency' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Frequency = src.Frequency
			} else {
				var zero uint64
				dst.Frequency = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BandDescription_DwellTime) SetFields(src *BandDescription_DwellTime, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplinks":
			if len(subs) > 0 {
				return fmt.Errorf("'uplinks' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Uplinks = src.Uplinks
			} else {
				dst.Uplinks = nil
			}
		case "downlinks":
			if len(subs) > 0 {
				return fmt.Errorf("'downlinks' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Downlinks = src.Downlinks
			} else {
				dst.Downlinks = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListBandsResponse_VersionedBandDescription) SetFields(src *ListBandsResponse_VersionedBandDescription, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "band":
			if len(subs) > 0 {
				return fmt.Errorf("'band' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Band = src.Band
			} else {
				dst.Band = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
