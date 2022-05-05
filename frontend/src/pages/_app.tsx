import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { Provider } from 'react-redux'
import Layout from '../components/Layout'
import Header from '../components/Header'
import Menu from '../components/Menu'
import { store } from '../store'

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <Header />
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </Provider>
  )
}