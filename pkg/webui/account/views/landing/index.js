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
import { Switch, Route } from 'react-router-dom'
import classnames from 'classnames'
import { useSelector } from 'react-redux'

import authRoutes from '@account/constants/auth-routes'

import sidebarStyle from '@ttn-lw/components/navigation/side/side.styl'

import Footer from '@ttn-lw/containers/footer'
import LogBackInModal from '@ttn-lw/containers/log-back-in-modal'

import { FullViewErrorInner } from '@ttn-lw/lib/components/full-view-error'

import Header from '@account/containers/header'

import { selectIsLoggedIn } from '@ttn-lw/lib/store/selectors/status'

import style from './landing.styl'

const GenericNotFound = () => <FullViewErrorInner error={{ statusCode: 404 }} />

const Landing = () => {
  const isLoggedIn = useSelector(selectIsLoggedIn)

  return (
    <div className={style.container}>
      {!isLoggedIn && <LogBackInModal />}
      <Header />
      <main className={style.main}>
        <div className={style.stage} id="stage">
          <div id="sidebar" className={sidebarStyle.container} />
          <div className={classnames('breadcrumbs', style.desktopBreadcrumbs)} />
          <Switch>
            {authRoutes.map(route => (
              <Route {...route} key={route.path} />
            ))}
            <Route component={GenericNotFound} />
          </Switch>
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default Landing
