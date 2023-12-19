/** @type {import('next').NextConfig} */
module.exports = (phase, { defaultConfig }) => {
 
    return {
        eslint: {
            ignoreDuringBuilds: true,
          },
          typescript: {
            // !! WARN !!
            // Dangerously allow production builds to successfully complete even if
            // your project has type errors.
            // !! WARN !!
            ignoreBuildErrors: true,
          },
        output: 'export',
        reactStrictMode: true,
        transpilePackages: ['@boostchicken/lol-api']
    }
}