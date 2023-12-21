/** @type {import('next').NextConfig} */
module.exports = (phase, { defaultConfig }) => {
 
    return {
        output: 'export',
        reactStrictMode: true
       // transpilePackages: ['@boostchicken/lol-api']
    }
}