'use client'
import { Metadata } from 'next'
import { ReactNode } from 'react'
import { WithSidebar } from './WithSidebar'
import { Fira_Code } from 'next/font/google'
import './globalStyle.css'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
 
const FiraCode = Fira_Code({
  weight: ['400', '500', '700'],
  display: 'swap',
  subsets: ['latin'],
  variable: '--font-fira-code'
})
type Props = {
  children: ReactNode,
}
export default function Layout({ children }: Props) {
  const client = new QueryClient()

  return (
    <html lang='ja' className={FiraCode.className}>
      <body>
        <QueryClientProvider client={client}>
          <WithSidebar>
            {children}
          </WithSidebar>
        </QueryClientProvider>
      </body>
    </html>
  )
}
 
export const metadata: Metadata = {
  title: 'speakit',
  description: 'speakit',
}
