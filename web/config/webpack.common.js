'use-strict';

const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const get = require('lodash/get');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const WebpackChunkHash = require('webpack-chunk-hash');

const stringifyValue = require('./util/stringifyValue');
const CONST = require('./constant');
const helpers = require('./util/helpers');

const { NODE_ENV } = process.env;
const isProd = NODE_ENV === 'production';

const plugins = [];
if (process.env.BUNDLE_ANALYSE) {
  plugins.push(new BundleAnalyzerPlugin());
}
// the path(s) that should be cleaned
const pathsToClean = isProd ? ['dist/*'] : [];

// the clean options to use
const cleanOptions = {
  root: helpers.root(),
  verbose: true,
  dry: false
};
module.exports = {
  entry: {
    app: helpers.root('src/index.js')
  },
  output: {
    path: helpers.root('dist'),
    publicPath: '/assets/'
  },
  resolve: {
    extensions: ['.js', '.json', '.css', '.scss', '.html'],
    modules: ['src', 'node_modules']
  },

  module: {
    rules: [
      // JS files
      {
        test: /\.(js|jsx)$/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              cacheDirectory: true
            }
          }
        ],
        exclude: /node_modules/
      },
      // SCSS files
      {
        test: /\.(sa|sc|c)ss$/,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          {
            loader: 'postcss-loader'
          },
          'sass-loader'
        ]
      },
      {
        test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'file-loader?mimetype=image/svg+xml'
      },
      {
        test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'file-loader?mimetype=application/font-woff'
      },
      {
        test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'file-loader?mimetype=application/font-woff'
      },
      {
        test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'file-loader?mimetype=application/octet-stream'
      },
      {
        test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'file-loader'
      },
      {
        test: /\.(png|jpg|gif)$/,
        use: [
          {
            loader: 'file-loader',
            options: {}
          }
        ]
      }
    ]
  },

  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(NODE_ENV)
      },
      ...stringifyValue(get(CONST, ['webpack', 'clientConstants']))
    }),

    new webpack.ProvidePlugin({
      // make fetch available
      fetch: 'exports-loader?self.fetch!whatwg-fetch'
    }),

    new webpack.ContextReplacementPlugin(
      /\.\/locale$/,
      'empty-module',
      false,
      /js$/
    ),

    new HtmlWebpackPlugin({
      template: helpers.root('src/public/index.html'),
      inject: 'body',
      hash: true,
      filename: '../index.html',
      title: 'Shorty'
    }),
    new HtmlWebpackPlugin({
      template: helpers.root('src/public/index.dev.html'),
      inject: 'body',
      hash: true
    }),
    new MiniCssExtractPlugin({
      filename: isProd ? 'css/[name].[contenthash].css' : 'css/[name].css',
      chunkFilename: !isProd
        ? '[name]/[name].css'
        : '[name]/[name].[contenthash].css'
    }),

    new CopyWebpackPlugin([
      {
        from: helpers.root('src/public'),
        to: helpers.root('dist')
      }
    ]),
    new CleanWebpackPlugin(pathsToClean, cleanOptions),
    new WebpackChunkHash({ algorithm: 'md5' }),
    ...plugins
  ],

  target: 'web'
};
