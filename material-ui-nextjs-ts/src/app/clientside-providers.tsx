// app/providers.jsx
'use client'

import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import React from "react";
import {TransportProvider} from '@connectrpc/connect-query';
import {ConnectTransport} from "@/utils/connect";

const ClientsideProviders: React.FC<React.PropsWithChildren> = ({children}) => {
  const [queryClient] = React.useState(() => new QueryClient())

  return (
    <TransportProvider transport={ConnectTransport}>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </TransportProvider>
  )
}

export default ClientsideProviders
