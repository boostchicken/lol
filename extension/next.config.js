/** @type {import('next').NextConfig} */

module.exports = (phase, { defaultConfig }) => {

    return {
        output: 'export',
        transpilePackages: ['@boostchicken/lol-api']
    }
}