// app/providers.jsx
'use client'

import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import React from "react";
import {TransportProvider} from '@connectrpc/connect-query';
import {ConnectTransport} from "@/utils/connect";

import {LocalizationProvider} from "@mui/x-date-pickers";
import {AdapterDateFns} from '@mui/x-date-pickers/AdapterDateFns';

const ClientsideProviders: React.FC<React.PropsWithChildren> = ({children}) => {
  const [queryClient] = React.useState(() => new QueryClient())

  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <TransportProvider transport={ConnectTransport}>
        <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
      </TransportProvider>
    </LocalizationProvider>
  )
}

export default ClientsideProviders
