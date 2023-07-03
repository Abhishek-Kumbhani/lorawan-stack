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
import { connect, useSelector } from 'react-redux'
import { defineMessages } from 'react-intl'
import { Col, Row, Container } from 'react-grid-system'
import { Routes, Route } from 'react-router-dom'

import PageTitle from '@ttn-lw/components/page-title'
import Tabs from '@ttn-lw/components/tabs'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumbs from '@ttn-lw/components/breadcrumbs'

import withRequest from '@ttn-lw/lib/components/with-request'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import NotFoundRoute from '@ttn-lw/lib/components/not-found-route'

import TokensTable from '@account/containers/tokens-table'

import AuthorizationSettings from '@account/views/oauth-authorization-settings'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import { selectApplicationSiteName } from '@ttn-lw/lib/selectors/env'
import PropTypes from '@ttn-lw/lib/prop-types'

import { getAuthorizationsList } from '@account/store/actions/authorizations'
import { getClientsList } from '@account/store/actions/clients'

import {
  selectAuthorizationsError,
  selectAuthorizationsFetching,
} from '@account/store/selectors/authorizations'
import { selectUserId } from '@account/store/selectors/user'
import { selectClientById } from '@account/store/selectors/clients'

import style from './authorization.styl'

const m = defineMessages({
  authorizationSettings: 'Authorization settings',
  accessTokens: 'Active access tokens',
})

const AuthorizationOverview = props => {
  const {
    clientId,
    match: { path },
    siteName,
  } = props

  const client = useSelector(state => selectClientById(state, clientId))
  const clientName = client?.name || clientId

  useBreadcrumbs(
    'client-authorizations.single',
    <Breadcrumb path={`/client-authorizations/${clientId}`} content={clientId} />,
  )

  const basePath = `/client-authorizations/${clientId}`

  const tabs = [
    {
      title: m.authorizationSettings,
      name: 'overview',
      link: `${basePath}`,
    },
    { title: m.accessTokens, name: 'access-tokens', link: `${basePath}/access-tokens` },
  ]

  return (
    <>
      <IntlHelmet titleTemplate={`%s - ${clientId} - ${siteName}`} />
      <Breadcrumbs />
      <div className={style.titleSection}>
        <Container>
          <Row>
            <Col sm={12}>
              <PageTitle title={clientName} className={style.pageTitle} />
              <Tabs className={style.tabs} narrow tabs={tabs} />
            </Col>
          </Row>
        </Container>
      </div>
      <Routes>
        <Route exact path={path} component={AuthorizationSettings} />
        <Route exact path={`${path}/access-tokens`} component={TokensTable} />
        <NotFoundRoute />
      </Routes>
    </>
  )
}

AuthorizationOverview.propTypes = {
  clientId: PropTypes.string.isRequired,
  match: PropTypes.match.isRequired,
  siteName: PropTypes.string.isRequired,
}

export default connect(
  (state, props) => ({
    userId: selectUserId(state),
    clientId: props.match.params.clientId,
    fetching: selectAuthorizationsFetching(state),
    error: selectAuthorizationsError(state),
    siteName: selectApplicationSiteName(),
  }),
  dispatch => ({
    loadData: userId => {
      dispatch(getAuthorizationsList(userId))
      dispatch(attachPromise(getClientsList(undefined, ['name'])))
    },
  }),
)(withRequest(({ userId, loadData }) => loadData(userId))(AuthorizationOverview))
