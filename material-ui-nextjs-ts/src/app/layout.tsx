import * as React from 'react';
import {AppRouterCacheProvider} from '@mui/material-nextjs/v14-appRouter';
import {ThemeProvider} from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import theme from '@/theme';
import {Toaster} from "react-hot-toast";
import ClientsideProviders from "@/app/clientside-providers";

export default function RootLayout(props: { children: React.ReactNode }) {
  return (
    <html lang="en">
    <body>
    <AppRouterCacheProvider options={{enableCssLayer: true}}>
      <ThemeProvider theme={theme}>

          <ClientsideProviders>
            {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
            <CssBaseline/>
            <Toaster />
            {props.children}
          </ClientsideProviders>
      </ThemeProvider>
    </AppRouterCacheProvider>
    </body>
    </html>
  );
}
