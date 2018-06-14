const path = require('path');

const config = {
	plugins: [
	]
}

module.exports = {
  entry: {
    'app': './src/app.js',
    'editor': './src/editor.js'
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'static', 'js')
  }
};
