import './globals.css'
import type { Metadata} from 'next'
import { Ubuntu_Mono } from 'next/font/google'

const font = Ubuntu_Mono({ weight: '400', preload: true, subsets: ['latin'] })

export const metadata: Metadata = {
  manifest: "manifest.json",
  title: "BoostLoL",

  description: "Admin interface for BoostLoL"
}
export const viewport = {
  width: 'device-width',
  initialScale: 1,
  maximumScale: 1,
  themeColor: 'dark',
}
export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html data-bs-theme="dark" lang="en">
      <body className={font.className}>{children}</body>
    </html>
  )
}
