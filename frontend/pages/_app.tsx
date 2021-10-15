import "@/css/styles.css";
import Head from "next/head";
import { AppProps } from "next/app";

import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";

export const queryClient = new QueryClient();
export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <link rel="icon" href="/favicon.ico" />

        <title>PLACEHOLDER</title>

        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta charSet="utf-8" />
        <meta
          name="description"
          content="An automated approach to small team project management"
        />
        <meta property="og:title" content="Deliberate" />
        <meta property="og:type" content="website" />
        <meta
          property="og:description"
          content="An automated approach to small team project management"
        />
        <meta
          property="og:image"
          content="https://via.placeholder.com/450x200.png"
        />
        <meta property="og:url" content="" />
        <meta name="twitter:card" content="summary_large_image" />
      </Head>

      <QueryClientProvider client={queryClient}>
        <Component {...pageProps} />
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </>
  );
}
