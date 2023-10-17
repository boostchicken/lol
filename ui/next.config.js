/** @type {import('next').NextConfig} */
const { PHASE_DEVELOPMENT_SERVER } = require('next/constants')
 
module.exports = (phase, { defaultConfig }) => {
  if (phase === PHASE_DEVELOPMENT_SERVER) {
    return {
       async rewrites() {
            return { fallback: [{
                source: '/:path*',
                destination: `http://localhost:6969/:path*`,
            }]}
        },
        reactStrictMode: true,
        transpilePackages: ['@boostchicken/lol-api']
    }
  }
 
    return {
        output: 'export',
        reactStrictMode: true,
        transpilePackages: ['@boostchicken/lol-api']
    }
}