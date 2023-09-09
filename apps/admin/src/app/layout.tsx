import { Metadata } from 'next'
import { QueryProvider } from './query'
import '../styles/global.css'
import { FiraCode } from '../styles/font'
import { ReactNode } from 'react'
import { WithSidebar } from './WithSidebar'

type Props = {
  children: ReactNode,
}
export default function AppLayout({ children }: Props) {
  return (
    <html lang='ja' className={FiraCode.className}>
      <body>
        <QueryProvider>
          <WithSidebar>
            {children}
          </WithSidebar>
        </QueryProvider>
      </body>
    </html>
  )
}
 
export const metadata: Metadata = {
  title: 'speakit',
  description: 'speakit',
}
