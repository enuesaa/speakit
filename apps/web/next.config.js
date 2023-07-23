const withPWA = require("next-pwa")({
  dest: "public",
  register: true,
  skipWaiting: true,
});

/** @type {import('next').NextConfig} */
const nextConfig = withPWA({
  reactStrictMode: true,
  output: 'export', // comment out this line when to use api routes.
  distDir: 'dist',
})

module.exports = nextConfig
