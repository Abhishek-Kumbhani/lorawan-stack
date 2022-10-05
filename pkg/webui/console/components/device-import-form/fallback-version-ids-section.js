// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { Col, Row } from 'react-grid-system'
import { useSelector } from 'react-redux'

import { useFormContext } from '@ttn-lw/components/form'
import ErrorNotification from '@ttn-lw/components/error-notification'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import OtherHint from '@console/containers/device-profile-section/hints/other-hint'
import VersionIdsSection from '@console/containers/device-profile-section'

import { hasSelectedDeviceRepositoryOther } from '@console/lib/device-utils'

import { listBrands } from '@console/store/actions/device-repository'

import { selectSelectedApplicationId } from '@console/store/selectors/applications'

const FallbackVersionIdsSection = () => {
  const appId = useSelector(selectSelectedApplicationId)
  const { values } = useFormContext()
  const { version_ids } = values
  const version = version_ids
  const hasSelectedOther = hasSelectedDeviceRepositoryOther(version)
  const showOtherHint = hasSelectedOther

  return (
    <Row>
      <Col>
        <RequireRequest
          requestAction={listBrands(appId, {}, ['name', 'lora_alliance_vendor_id'])}
          errorRenderFunction={ErrorNotification}
          spinnerProps={{ center: false, inline: true }}
        >
          <VersionIdsSection />
          {showOtherHint && <OtherHint manualGuideDocsPath="/devices/adding-devices/" />}
        </RequireRequest>
      </Col>
    </Row>
  )
}

export default FallbackVersionIdsSection
