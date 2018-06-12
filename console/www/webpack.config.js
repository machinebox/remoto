const path = require('path');

const config = {
	plugins: [
	]
}

module.exports = {
  entry: {
    'app': './src/app.js'
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'static', 'js')
  }
};
