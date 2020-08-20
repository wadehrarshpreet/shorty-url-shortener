'use-strict';

const isProd = process.env.NODE_ENV === 'production';
const merge = require('lodash/merge');

const commonConfig = {};
const config = isProd ? require('./constant.prod') : require('./constant.dev');

module.exports = merge(commonConfig, config);
