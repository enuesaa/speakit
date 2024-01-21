import { Metadata } from 'next'
import { QueryProvider } from './query'
import { FiraCode } from './font'
import { ReactNode } from 'react'
import { WithSidebar } from './WithSidebar'
import './globalStyle.css'

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
