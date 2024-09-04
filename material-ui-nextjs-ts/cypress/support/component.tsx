// ***********************************************************
// This example support/component.ts is processed and
// loaded automatically before your test files.
//
// This is a great place to put global configuration and
// behavior that modifies Cypress.
//
// You can change the location of this file or turn off
// automatically serving support files with the
// 'supportFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/configuration
// ***********************************************************

// Import commands.js using ES2015 syntax:
import './commands'

// Alternatively you can use CommonJS syntax:
// require('./commands')

import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import {TransportProvider} from '@connectrpc/connect-query';
import { mount } from 'cypress/react18'
import {Toaster} from "react-hot-toast";
import * as React from "react";
import {ConnectTransport} from "../../src/utils/connect";
import {createConnectTransport} from "@connectrpc/connect-web";
import {createRegistry} from "@bufbuild/protobuf";
// Augment the Cypress namespace to include type definitions for
// your custom command.
// Alternatively, can be defined in cypress/support/component.d.ts
// with a <reference path="./component" /> at the top of your spec.
declare global {
  namespace Cypress {
    interface Chainable {
      mount: typeof mount
    }
  }
}

const testTransport = createConnectTransport({
  baseUrl: 'http://mock-cypress-test.com',
  jsonOptions: {
    typeRegistry: createRegistry(
      //Account,
    )
  },
});

const customMount: typeof mount = (component, options = {}) => {
  const { ...mountOptions } = options
  const queryClient = new QueryClient()

  const wrapped = <TransportProvider transport={testTransport}>
    <QueryClientProvider client={queryClient}>
      <Toaster />
      {component}
    </QueryClientProvider>
  </TransportProvider>

  return mount(wrapped, mountOptions)
}

Cypress.Commands.add('mount', customMount)

// Example use:
// cy.mount(<MyComponent />)
