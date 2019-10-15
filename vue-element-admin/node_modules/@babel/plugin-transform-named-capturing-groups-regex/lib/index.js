"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = _default;

var _regexpuCore = _interopRequireDefault(require("regexpu-core"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _default({
  types: t
}, options) {
  const {
    runtime = true
  } = options;

  if (typeof runtime !== "boolean") {
    throw new Error("The 'runtime' option must be boolean");
  }

  return {
    name: "transform-named-capturing-groups-regex",
    visitor: {
      RegExpLiteral(path) {
        const node = path.node;

        if (!/\(\?<(?![=!])/.test(node.pattern)) {
          return;
        }

        const namedCapturingGroups = {};
        const result = (0, _regexpuCore.default)(node.pattern, node.flags, {
          namedGroup: true,
          lookbehind: true,

          onNamedGroup(name, index) {
            namedCapturingGroups[name] = index;
          }

        });

        if (Object.keys(namedCapturingGroups).length > 0) {
          node.pattern = result;

          if (runtime && !isRegExpTest(path)) {
            path.replaceWith(t.callExpression(this.addHelper("wrapRegExp"), [node, t.valueToNode(namedCapturingGroups)]));
          }
        }
      }

    }
  };
}

function isRegExpTest(path) {
  return path.parentPath.isMemberExpression({
    object: path.node,
    computed: false
  }) && path.parentPath.get("property").isIdentifier({
    name: "test"
  });
}