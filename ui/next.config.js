/** @type {import('next').NextConfig} */
const { PHASE_DEVELOPMENT_SERVER } = require('next/constants')
 const withWorkbox = require("next-with-workbox");

module.exports = withWorkbox({
  workbox: {
    dest: "public",
    swDest: "sw.js",
    swSrc: "worker.js",
    force: true,
  },
});
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
        transpilePackages: ['@boostchicken/lol-api'],
        modularizeImports: {
            '@boostchicken/lol-api': {
                transform: '@boostchicken/lol-api/hooks/{{ member }}'
            }
        },
    }
  }
 
    return {
        output: 'export',
        reactStrictMode: true,
        transpilePackages: ['@boostchicken/lol-api'],
        modularizeImports: {
            '@boostchicken/lol-api': {
                transform: '@boostchicken/lol-api/hooks/{{ member }}'
            }
        },
    }
}