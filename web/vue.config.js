module.exports = {
  devServer: {
    proxy: {
      '^/auth': {
        target: 'http://localhost:10000',
        changeOrigin: true,
        logLevel: 'debug',
        pathRewrite: { '^/auth': '/' }
      },
      '^/user': {
        target: 'http://localhost:10001',
        changeOrigin: true,
        logLevel: 'debug',
        pathRewrite: { '^/user': '/' }
      }
    }
  }
}
