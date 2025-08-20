import './App.css';
import { PrimeReactProvider } from 'primereact/api';
import { Button } from 'primereact/button';

import type { ComponentType } from 'react';
import Header from './component/header';
import Body from './component/body';

interface MyAppProps {
  Component: ComponentType<any>;
  pageProps: any;
}

export default function MyApp({ Component, pageProps }: MyAppProps) {
  return (
    <PrimeReactProvider>
      <Component {...pageProps} />
      <Header />
      <Body />
    </PrimeReactProvider>
  );
}