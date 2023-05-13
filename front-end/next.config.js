/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  experimental: {
    largePageDataBytes: 128 * 100000,
  },
  output: 'standalone',
}

module.exports = nextConfig
