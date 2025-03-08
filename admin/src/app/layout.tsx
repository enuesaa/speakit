'use client'
import { ReactNode } from 'react'
import { Fira_Code } from 'next/font/google'
import './globalStyle.css'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { Sidebar } from './Sidebar'

const FiraCode = Fira_Code({
  weight: ['400', '500', '700'],
  display: 'swap',
  subsets: ['latin'],
  variable: '--font-fira-code',
})
type Props = {
  children: ReactNode
}
export default function Layout({ children }: Props) {
  const client = new QueryClient()

  return (
    <html lang='ja' className={FiraCode.className}>
      <body>
        <QueryClientProvider client={client}>
          <main style={{ display: 'flex' }}>
            <Sidebar />
            <section>{children}</section>
          </main>
        </QueryClientProvider>
      </body>
    </html>
  )
}
