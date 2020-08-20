'use-strict';

module.exports = (hash = {}) => {
  const returnObj = {};
  Object.keys(hash).forEach((key) => {
    returnObj[key] = JSON.stringify(hash[key]);
  });
  return returnObj;
};
