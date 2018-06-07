const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const config = {
  plugins: []
}

module.exports = {
  entry: {
    'app': './src/index.js'
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'dist')
  },
  module: {
    rules: [
        {
            test: /\.js$/,
            loader: 'babel-loader',
            query: {
                presets: ['es2015']
            }
        }
    ]
  }
};
