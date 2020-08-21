'use-strict';

const webpack = require('webpack');
const merge = require('webpack-merge');
const path = require('path');
const CircularDependencyPlugin = require('circular-dependency-plugin');

const helpers = require('./util/helpers');
const commonConfig = require('./webpack.common');

module.exports = merge(commonConfig, {
  devtool: 'eval-source-map',

  mode: 'development',

  entry: helpers.root('src/index.js'),
  output: {
    filename: 'assets/js/[name].js',
    chunkFilename: 'assets/js/[id].chunk.js'
  },
  optimization: {
    namedChunks: true,
    splitChunks: {
      chunks: 'all'
    }
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.optimize.OccurrenceOrderPlugin(),
    new CircularDependencyPlugin({
      exclude: /a\.js|node_modules/, // exclude node_modules
      failOnError: false // show a warning when there is a circular dependency
    })
  ],
  devServer: {
    contentBase: path.join(__dirname, '../dist'),
    historyApiFallback: true,
    port: 3002,
    disableHostCheck: true,
    host: '0.0.0.0',
    compress: true,
    hot: true,
    stats: 'errors-only' // none (or false), errors-only, minimal, normal (or true) and verbose
  }
});
