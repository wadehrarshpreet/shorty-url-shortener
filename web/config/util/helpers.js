const path = require('path');

// Helper functions
function root(args) {
  args = Array.prototype.slice.call(arguments, 0);
  // eslint-disable-next-line
  return path.join.apply(path, [__dirname].concat('../../', ...args));
}

exports.root = root;
