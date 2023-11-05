import './globals.css'
import type { Metadata} from 'next'
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  manifest: "manifest.json",
  title: "BoostLoL",
  viewport: {
    width: "device-width",
    initialScale: 1,
    viewportFit: "cover",
    colorScheme: 'dark'
  },

  description: "Admin interface for BoostLoL"
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html data-bs-theme="dark" lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
